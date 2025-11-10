package cartcontroller

import (
	"kingcom_api/internal/controllers/handler"
	"kingcom_api/internal/request"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *CartController) FetchCart(c *gin.Context) {
	hh := handler.NewHandlerHelper(c, ctrl.logger)
	tp, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		hh.ResErrUnauthorized(err)
		return
	}

	uid, err := uuid.Parse(tp.UserId)
	if err != nil {
		hh.ResInternalServerErr(err)
		return
	}

	carts, err := ctrl.cartService.FindWithProduct(uid)
	if err != nil {
		hh.ResInternalServerErr(err)
		return
	}

	hh.ResOk(gin.H{"carts": carts})

}
