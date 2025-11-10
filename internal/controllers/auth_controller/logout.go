package authcontroller

import (
	"kingcom_api/internal/constants"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *AuthController) Logout(c *gin.Context) {

	res := response.New(c, ctrl.logger)

	refToken, err := request.GetRefreshToken(c)
	if err != nil {
		res.ResErr(http.StatusUnauthorized, err, "")
		return
	}

	payload, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		res.ResErr(http.StatusUnauthorized, err, "")
		return
	}

	// invalidate access and refresh token
	if err := ctrl.authSvc.DeleteAuthTokens(c.Request.Context(), refToken, payload.Jti); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(constants.COOKIE_REFRESH_TOKEN, "", -1, "/", "", false, false)

	res.ResOk(gin.H{
		"message": "Logout",
	})
}
