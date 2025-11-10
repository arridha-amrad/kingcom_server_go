package authcontroller

import (
	"kingcom_api/internal/constants"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *AuthController) Logout(c *gin.Context) {
	refToken, err := request.GetRefreshToken(c)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusUnauthorized, err, "")
		return
	}

	payload, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusUnauthorized, err, "")
		return
	}

	// invalidate access and refresh token
	if err := ctrl.authSvc.DeleteAuthTokens(c.Request.Context(), refToken, payload.Jti); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(constants.COOKIE_REFRESH_TOKEN, "", -1, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "Logout"})
}
