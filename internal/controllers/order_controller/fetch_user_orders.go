package ordercontroller

import (
	"kingcom_api/internal/controllers/handler"
	"kingcom_api/internal/request"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (o *OrderController) FetchUserOrders(c *gin.Context) {
	hh := handler.NewHandlerHelper(c, o.logger)
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

	orders, err := o.orderService.FindByUserId(userId)
	if err != nil {
		hh.ResInternalServerErr(err)
		return
	}

	hh.ResOk(orders)
}
