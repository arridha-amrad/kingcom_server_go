package authcontroller

import (
	"errors"
	"kingcom_api/internal/constants"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (ctrl *AuthController) Login(c *gin.Context) {

	res := response.New(c, ctrl.logger)

	body, errValidation := request.GetBody[dto.Login](c, ctrl.validate)
	if errValidation != nil {
		res.ResErrValidation(errValidation)
		return
	}

	// find user
	user, err := ctrl.userSvc.FindByEmailOrUsername(body.Identity)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if user == nil {
		res.ResErr(http.StatusNotFound, err, "user not found")
		return
	}
	if !user.IsVerified {
		err := errors.New("please verify your account first")
		res.ResErr(http.StatusUnauthorized, err, err.Error())
		return
	}

	// compare password
	if err := ctrl.authSvc.VerifyPassword(user.Password, body.Password); err != nil {
		res.ResErr(http.StatusUnauthorized, err, "invalid credentials")
		return
	}

	// generate access and refresh token
	tokens, err := ctrl.authSvc.CreateAuthTokens(
		c.Request.Context(),
		user.ID.String(),
		user.JwtVersion,
		string(user.Role),
	)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// set response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		constants.COOKIE_REFRESH_TOKEN,
		tokens.RefreshToken,
		constants.REFRESH_TOKEN_MAX_AGE, "/", "", os.Getenv("GO_ENV") == "production", true,
	)

	res.ResOk(gin.H{
		"user":  user,
		"token": tokens.AccessToken,
	})

}
