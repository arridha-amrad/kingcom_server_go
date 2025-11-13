package cartcontroller

import (
	"errors"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *CartController) DeleteCart(c *gin.Context) {

	res := response.New(c, ctrl.logger)

	strCartId := c.Param("cartId")

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

	cartId, err := uuid.Parse(strCartId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	cart, err := ctrl.cartService.FindOne(cartId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if cart == nil {
		err := errors.New("cart not found")
		res.ResErr(http.StatusNotFound, err, err.Error())
		return
	}

	if cart.UserID != userId {
		err := errors.New("you are not own this cart")
		res.ResErr(http.StatusForbidden, err, "")
		return
	}

	if err := ctrl.cartService.Delete(cartId); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	res.ResOk(gin.H{"message": "one cart item deleted"})
}
