package cartcontroller

import (
	"kingcom_api/internal/dto"
	"kingcom_api/internal/models"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *CartController) AddToCart(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	body, errV := request.GetBody[dto.AddToCart](c, ctrl.validator)
	if errV != nil {
		res.ResErrValidation(errV)
		return
	}

	tp, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	userId, err := uuid.Parse(tp.UserId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	cart := models.Cart{
		Quantity:  body.Quantity,
		UserID:    userId,
		ProductID: body.ProductID,
	}

	if err := ctrl.cartService.Add(&cart); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	res.ResCreated("Added to cart")
}
