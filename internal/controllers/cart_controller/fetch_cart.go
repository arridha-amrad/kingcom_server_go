package cartcontroller

import (
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *CartController) FetchCart(c *gin.Context) {
	res := response.New(c, ctrl.logger)
	tp, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		res.ResErrUnauthorized(err)
		return
	}

	uid, err := uuid.Parse(tp.UserId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	carts, err := ctrl.cartService.FindWithProduct(uid)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	res.ResOk(gin.H{"carts": carts})

}
