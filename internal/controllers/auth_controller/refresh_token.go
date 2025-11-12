package authcontroller

import (
	"errors"
	"kingcom_api/internal/constants"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"kingcom_api/internal/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	res := response.New(c, ctrl.logger)

	refToken, err := request.GetRefreshToken(c)
	if err != nil {
		res.ResErrUnauthorized(err)
		return
	}

	hashedToken := utils.HashWithSHA256(refToken)
	ctx := c.Request.Context()

	// find refresh token
	payload, err := ctrl.cacheSvc.FindRefreshToken(ctx, hashedToken)
	if err != nil {
		res.ResErrUnauthorized(err)
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

	// delete old auth tokens
	if err := ctrl.cacheSvc.DeleteAccessToken(ctx, payload.Jti); err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if err := ctrl.cacheSvc.DeleteRefreshToken(ctx, refToken); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// generate new access and refresh token
	tokens, err := ctrl.authSvc.CreateAuthTokens(
		ctx,
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
		"token": tokens.AccessToken,
	})
}
