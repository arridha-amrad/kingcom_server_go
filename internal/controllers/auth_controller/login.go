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
	body, errBind := request.GetBody[dto.Login](c, ctrl.validate)
	if errBind != nil {
		response.ResValidationErr(c, ctrl.logger, errBind)
		return
	}

	// find user
	user, err := ctrl.userSvc.FindByEmailOrUsername(body.Identity)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	if user == nil {
		response.ResErr(c, ctrl.logger, http.StatusNotFound, err, "")
		return
	}
	if !user.IsVerified {
		response.ResErr(c, ctrl.logger, http.StatusConflict, errors.New("unverified"), "please verify your account first")
		return
	}

	// compare password
	if err := ctrl.authSvc.VerifyPassword(user.Password, body.Password); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusUnauthorized, err, "invalid credentials")
		return
	}

	// generate access and refresh token
	tokens, err := ctrl.authSvc.CreateAuthTokens(
		c.Request.Context(),
		user.ID.String(),
		user.JwtVersion,
	)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	// set response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		constants.COOKIE_REFRESH_TOKEN,
		tokens.RefreshToken,
		constants.REFRESH_TOKEN_MAX_AGE, "/", "", os.Getenv("GO_ENV") == "production", true,
	)
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": tokens.AccessToken,
	})
}
