package ordercontroller

import (
	"kingcom_api/internal/controllers/handler"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/models"
	"kingcom_api/internal/request"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *OrderController) PlaceOrder(c *gin.Context) {

	hh := handler.NewHandlerHelper(c, ctrl.logger)

	body, errV := request.GetBody[dto.CreateOrder](c, ctrl.validator)
	if errV != nil {
		hh.ResErrValidation(errV)
		return
	}

	tp, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		hh.ResErrUnauthorized(err)
		return
	}

	userId, err := uuid.Parse(tp.UserId)
	if err != nil {
		hh.ResInternalServerErr(err)
		return
	}

	// collect the cartId of each items
	cartIds := make([]uuid.UUID, 0, len(body.Items))

	// check product availability
	for _, item := range body.Items {
		cartIds = append(cartIds, item.CartID)
		prod, err := ctrl.productService.FindById(item.ProductID)
		if err != nil {
			hh.ResInternalServerErr(err)
			return
		}
		if prod == nil {
			hh.ResErr(http.StatusNotFound, err, "product not found")
			return
		}
		if prod.Stock <= 0 {
			hh.ResErr(http.StatusConflict, err, "product out of stock")
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
		orderItems = append(orderItems, models.OrderItem{Quantity: item.Quantity, ProductID: item.ProductID})
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
		hh.ResInternalServerErr(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Your order is made successfully",
	})
}
