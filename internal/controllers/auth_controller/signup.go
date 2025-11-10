package authcontroller

import (
	"errors"
	"fmt"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/models"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"kingcom_api/internal/services"
	"kingcom_api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *AuthController) SignUp(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	body, errValidation := request.GetBody[dto.SignUp](c, ctrl.validate)
	if errValidation != nil {
		res.ResErrValidation(errValidation)
		return
	}

	ctx := c.Request.Context()

	// Username must be unique
	user, err := ctrl.userSvc.FindByUsername(body.Username)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if user != nil {
		err := errors.New("username has been registered")
		res.ResErr(http.StatusConflict, err, err.Error())
		return
	}

	// Email must be unique
	user, err = ctrl.userSvc.FindByEmail(body.Email)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if user != nil {
		err := errors.New("email has been registered")
		res.ResErr(http.StatusConflict, err, err.Error())
		return
	}

	// hash password
	pwd, err := ctrl.authSvc.HashPassword(body.Password)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// insert user data
	jwtVersion, err := utils.GenerateRandomBytes(8)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	var newUser = models.User{
		Username:   body.Username,
		Name:       body.Name,
		Email:      body.Email,
		Password:   pwd,
		Provider:   models.ProviderCredentials,
		IsVerified: false,
		JwtVersion: jwtVersion,
		Role:       models.RoleUser,
	}
	if err := ctrl.userSvc.Create(&newUser); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// create verification data
	data, token, err := ctrl.authSvc.CreateVerificationToken(ctx, newUser.ID.String())
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// send email verification
	if err := ctrl.mailSvc.SendVerificationEmail(ctx, services.VerificationParams{
		SendEmailParams: services.SendEmailParams{
			Name:  newUser.Name,
			Email: newUser.Email,
		},
		Code: data.Code,
	},
	); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	res.ResOk(gin.H{
		"token":   token,
		"message": fmt.Sprintf("An email has been sent to %s. Please follow the instruction to verify your account.", newUser.Email),
	})
}
