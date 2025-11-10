package authcontroller

import (
	"fmt"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	svc "kingcom_api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *AuthController) ResendVerification(c *gin.Context) {
	body, errValidation := request.GetBody[dto.ResendVerification](c, ctrl.validate)
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
		response.ResErr(c, ctrl.logger, http.StatusNotFound, err, "user not found")
		return
	}
	if user.IsVerified {
		response.ResErr(c, ctrl.logger, http.StatusConflict, err, "your email has been verified")
		return
	}

	// create verification data
	data, token, err := ctrl.authSvc.CreateVerificationToken(ctx, user.ID.String())
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
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
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": fmt.Sprintf("An email has been sent to %s. Please follow the instruction to verify your account.", user.Email)},
	)
}
