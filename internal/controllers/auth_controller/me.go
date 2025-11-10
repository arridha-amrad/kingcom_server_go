package authcontroller

import (
	"errors"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *AuthController) Me(c *gin.Context) {

	res := response.New(c, ctrl.logger)

	payload, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		res.ResErrUnauthorized(err)
		return
	}

	userId, err := uuid.Parse(payload.UserId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	user, err := ctrl.userSvc.FindById(userId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if user == nil {
		err := errors.New("user not found")
		res.ResErr(http.StatusNotFound, err, err.Error())
		return
	}

	res.ResOk(gin.H{
		"user": user,
	})

}
