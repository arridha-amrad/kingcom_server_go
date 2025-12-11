package ordercontroller

import (
	"kingcom_api/internal/dto"
	"kingcom_api/internal/models"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *OrderController) PlaceOrder(c *gin.Context) {

	res := response.New(c, ctrl.logger)

	body, errV := request.GetBody[dto.CreateOrder](c, ctrl.validator)
	if errV != nil {
		res.ResErrValidation(errV)
		return
	}

	request.PrintBody(ctrl.logger, body)

	tp, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		res.ResErrUnauthorized(err)
		return
	}

	userId, err := uuid.Parse(tp.UserId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// collect the cartId of each items
	cartIds := make([]uuid.UUID, 0, len(body.Items))

	// check product availability
	for _, item := range body.Items {
		cartIds = append(cartIds, item.CartID)
		prod, err := ctrl.productService.FindById(item.ProductID)
		request.PrintBody(ctrl.logger, prod)
		if err != nil {
			res.ResInternalServerErr(err)
			return
		}
		if prod == nil {
			res.ResErr(http.StatusNotFound, err, "product not found")
			return
		}
		if prod.Stock <= 0 {
			res.ResErr(http.StatusConflict, err, "product out of stock")
			return
		}
	}

	// compose the cart model
	shipping := models.OrderShipping{
		Name:        body.Shipping.Name,
		Code:        body.Shipping.Code,
		Service:     body.Shipping.Code,
		Description: body.Shipping.Code,
		Cost:        body.Shipping.Cost,
		Etd:         body.Shipping.Etd,
		Address:     body.Shipping.Address,
	}

	orderItems := make([]models.OrderItem, 0, len(body.Items))
	for _, item := range body.Items {
		orderItems = append(orderItems, models.OrderItem{
			Quantity:          item.Quantity,
			ProductID:         item.ProductID,
			PriceAtOrder:      item.PriceAtOrder,
			DiscountAtOrder:   item.DiscountAtOrder,
			FinalPriceAtOrder: item.FinalPriceAtOrder,
		})
	}

	order := models.Order{
		Status:     models.OrderStatusPending,
		Total:      body.Total,
		CreatedAt:  time.Now(),
		Shipping:   shipping,
		UserID:     userId,
		OrderItems: orderItems,
	}

	if err := ctrl.orderService.PlaceOrder(&order, cartIds); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Your order is made successfully",
	})
}
