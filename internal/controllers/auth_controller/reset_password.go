package authcontroller

import (
	"errors"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"kingcom_api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *AuthController) ResetPassword(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	body, errValidation := request.GetBody[dto.ResetPassword](c, ctrl.validate)
	if errValidation != nil {
		res.ResErrValidation(errValidation)
		return
	}

	hashedToken := utils.HashWithSHA256(body.Token)
	ctx := c.Request.Context()

	// find pwdReset data in redis
	payload, err := ctrl.cacheSvc.FindPasswordResetToken(ctx, hashedToken)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// find the user
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

	// hash password
	newPassword, err := ctrl.authSvc.HashPassword(body.Password)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// update user password and its jwtVersion
	newJwtVersion, err := utils.GenerateRandomBytes(4)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if err := ctrl.userSvc.UpdatePassword(userId, newPassword, newJwtVersion); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// delete pwdReset data from redis
	if err := ctrl.cacheSvc.DeletePasswordResetToken(ctx, hashedToken); err != nil {
		ctrl.logger.Error(err)
	}

	res.ResOk(gin.H{"message": "Reset password is successful"})

}
