package authcontroller

import (
	"errors"
	"fmt"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	svc "kingcom_api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *AuthController) ResendVerification(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	body, errValidation := request.GetBody[dto.ResendVerification](c, ctrl.validate)
	if errValidation != nil {
		res.ResErrValidation(errValidation)
		return
	}

	ctx := c.Request.Context()

	// find the user
	user, err := ctrl.userSvc.FindByEmail(body.Email)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if user == nil {
		err := errors.New("user not found")
		res.ResErr(http.StatusNotFound, err, err.Error())
		return
	}
	if user.IsVerified {
		err := errors.New("please verify your account first")
		res.ResErr(http.StatusUnauthorized, err, err.Error())
		return
	}

	// create verification data
	data, token, err := ctrl.authSvc.CreateVerificationToken(ctx, user.ID.String())
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// send the email
	if err := ctrl.mailSvc.SendVerificationEmail(ctx, svc.VerificationParams{
		Code: data.Code,
		SendEmailParams: svc.SendEmailParams{
			Name:  user.Name,
			Email: user.Email,
		},
	},
	); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	res.ResOk(gin.H{
		"token":   token,
		"message": fmt.Sprintf("An email has been sent to %s. Please follow the instruction to verify your account.", user.Email)},
	)
}
