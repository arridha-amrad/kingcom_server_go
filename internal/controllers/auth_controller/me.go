package authcontroller

import (
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *AuthController) Me(c *gin.Context) {
	payload, err := request.ExtractAccessTokenPayload(c)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusBadRequest, err, "")
		return
	}

	userId, err := uuid.Parse(payload.UserId)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	user, err := ctrl.userSvc.FindById(userId)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusNotFound, err, "")
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
