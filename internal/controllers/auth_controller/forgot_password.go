package authcontroller

import (
	"errors"
	"fmt"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"kingcom_api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *AuthController) ForgotPassword(c *gin.Context) {
	body, errValidation := request.GetBody[dto.ForgotPassword](c, ctrl.validate)
	if errValidation != nil {
		response.ResValidationErr(c, ctrl.logger, errValidation)
		return
	}

	ctx := c.Request.Context()

	// find the user
	user, err := ctrl.userSvc.FindByEmail(body.Email)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	if user == nil {
		err := errors.New("user not found")
		response.ResErr(c, ctrl.logger, http.StatusNotFound, err, err.Error())
		return
	}
	if !user.IsVerified {
		response.ResErr(c, ctrl.logger, http.StatusBadRequest, errors.New("unverified user"), "please verify your account first")
		return
	}

	// generate token for pwd reset
	token, err := ctrl.authSvc.CreatePwdResetToken(ctx, user.ID.String())
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	// send the link via email
	if err := ctrl.mailSvc.SendPasswordResetEmail(ctx,
		services.PasswordResetParams{
			Token: token,
			SendEmailParams: services.SendEmailParams{
				Name:  user.Name,
				Email: user.Email,
			},
		}); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	// set response
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("An email has been sent to %s. Please follow the instruction to reset your password", body.Email),
	})
}
