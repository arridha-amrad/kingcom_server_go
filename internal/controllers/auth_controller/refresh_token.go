package authcontroller

import (
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
	refToken, err := request.GetRefreshToken(c)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusUnauthorized, err, "")
		return
	}

	hashedToken := utils.HashWithSHA256(refToken)
	ctx := c.Request.Context()

	// find refresh token
	payload, err := ctrl.cacheSvc.FindRefreshToken(ctx, hashedToken)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusNotFound, err, "")
		return
	}

	// find the user
	userId, err := uuid.Parse(payload.UserId)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	user, err := ctrl.userSvc.FindById(userId)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	if user == nil {
		response.ResErr(c, ctrl.logger, http.StatusNotFound, err, "user not found")
		return
	}

	// delete old auth tokens
	if err := ctrl.cacheSvc.DeleteAccessToken(ctx, payload.Jti); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	if err := ctrl.cacheSvc.DeleteRefreshToken(ctx, refToken); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	// generate new access and refresh token
	tokens, err := ctrl.authSvc.CreateAuthTokens(
		ctx,
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
		"token": tokens.AccessToken,
	})
}
