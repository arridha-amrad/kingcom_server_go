package controllers

import (
	"errors"
	"fmt"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"kingcom_api/internal/services"
	"kingcom_api/internal/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransController struct {
	env          *lib.Env
	logger       *lib.Logger
	validator    *lib.Validator
	orderService services.OrderService
}

func NewMidtransController(env *lib.Env, logger *lib.Logger, validator *lib.Validator, orderService services.OrderService) *MidtransController {
	return &MidtransController{
		env:          env,
		logger:       logger,
		validator:    validator,
		orderService: orderService,
	}
}

var s = snap.Client{}

func (ctrl *MidtransController) CreateTransactionToken(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	body, errV := request.GetBody[dto.CreateTransactionTokenParam](c, ctrl.validator)
	if errV != nil {
		res.ResErrValidation(errV)
		return
	}

	uOrderId, err := uuid.Parse(body.OrderId)
	if err != nil {
		res.ResErr(http.StatusInternalServerError, err, "")
		return
	}

	order, err := ctrl.orderService.FindById(uOrderId)
	if err != nil {
		res.ResErr(http.StatusInternalServerError, err, "")
		return
	}

	if order == nil {
		res.ResErr(http.StatusNotFound, nil, "order not found")
		return
	}

	s.New(ctrl.env.Midtrans.ServerKey, midtrans.Sandbox)

	orderId := fmt.Sprintf("%s_%s", body.OrderId, fmt.Sprint(time.Now().Unix()))

	items := make([]midtrans.ItemDetails, 0, len(order.OrderItems)+1)

	for _, item := range order.OrderItems {
		items = append(items, midtrans.ItemDetails{
			Name:  item.Product.Name,
			Qty:   int32(item.Quantity),
			Price: int64(item.FinalPriceAtOrder) / int64(item.Quantity),
		})
	}

	items = append(items, midtrans.ItemDetails{
		Name:  fmt.Sprintf("%s (Courier)", order.Shipping.Name),
		Qty:   1,
		Price: int64(order.Shipping.Cost),
	})

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: order.Total,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: order.User.Name,
			Email: order.User.Email,
			ShipAddr: &midtrans.CustomerAddress{
				Address: order.Shipping.Address,
			},
		},
		Items: &items,
	}

	snapResp, _ := s.CreateTransaction(req)

	res.ResOk(gin.H{"token": snapResp.Token, "redirectUrl": snapResp.RedirectURL})
}

func (ctrl *MidtransController) HandleNotification(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	body, errV := request.GetBody[dto.MidtransNotificationParams](c, ctrl.validator)
	if errV != nil {
		res.ResErrValidation(errV)
		return
	}

	ctrl.logger.Info(fmt.Sprintf("notification body: %v", body))

	shouldProcessNotification := body.StatusCode == "200" &&
		(body.TransactionStatus == "settlement" || body.TransactionStatus == "capture")

	if !shouldProcessNotification {
		err := errors.New("not processing notification")
		res.ResErr(http.StatusUnauthorized, err, "")
		return
	}

	isValid := utils.VerifySHA512Signature(body.SignatureKey, body.OrderID, body.StatusCode, body.GrossAmount, ctrl.env.Midtrans.ServerKey)
	if !isValid {
		err := errors.New("invalid signature")
		res.ResErr(http.StatusUnauthorized, err, err.Error())
		return
	}

	strOrderId := strings.Split(body.OrderID, "_")[0]

	orderId, err := uuid.Parse(strOrderId)
	if err != nil {
		res.ResErr(http.StatusInternalServerError, err, "")
		return
	}

	order, err := ctrl.orderService.FindById(orderId)
	if err != nil {
		res.ResErr(http.StatusInternalServerError, err, "")
		return
	}

	if order == nil {
		res.ResErr(http.StatusNotFound, nil, "order not found")
		return
	}

	order.Status = "paid"
	order.PaymentMethod = body.PaymentType
	if err := ctrl.orderService.Update(order); err != nil {
		res.ResErr(http.StatusInternalServerError, err, "")
		return
	}

	res.ResOk(gin.H{"message": "success"})

}
