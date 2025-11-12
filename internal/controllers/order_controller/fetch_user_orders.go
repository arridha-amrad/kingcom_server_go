package ordercontroller

import (
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (o *OrderController) FetchUserOrders(c *gin.Context) {
	res := response.New(c, o.logger)
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

	orders, err := o.orderService.FindByUserId(userId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	res.ResOk(gin.H{"orders": orders})
}
