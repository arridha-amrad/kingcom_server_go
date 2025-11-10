package cartcontroller

import (
	"kingcom_api/internal/controllers/handler"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/models"
	"kingcom_api/internal/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *CartController) AddToCart(c *gin.Context) {
	hh := handler.NewHandlerHelper(c, ctrl.logger)

	body, errV := request.GetBody[dto.AddToCart](c, ctrl.validator)
	if errV != nil {
		hh.ResErrValidation(errV)
		return
	}

	tp, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		hh.ResInternalServerErr(err)
		return
	}

	userId, err := uuid.Parse(tp.UserId)
	if err != nil {
		hh.ResInternalServerErr(err)
		return
	}

	cart := models.Cart{
		Quantity:  body.Quantity,
		UserID:    userId,
		ProductID: body.ProductID,
	}

	if err := ctrl.cartService.Add(&cart); err != nil {
		hh.ResInternalServerErr(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Added to cart"})
}
