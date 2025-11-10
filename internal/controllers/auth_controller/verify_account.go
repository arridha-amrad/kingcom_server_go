package authcontroller

import (
	"errors"
	"kingcom_api/internal/constants"
	"kingcom_api/internal/dto"
	"kingcom_api/internal/request"
	"kingcom_api/internal/response"
	"kingcom_api/internal/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *AuthController) VerifyNewAccount(c *gin.Context) {
	body, errValidation := request.GetBody[dto.VerifyNewAccount](c, ctrl.validate)
	if errValidation != nil {
		response.ResValidationErr(c, ctrl.logger, errValidation)
		return
	}

	hashedToken := utils.HashWithSHA256(body.Token)
	ctx := c.Request.Context()

	// find verification data in redis
	payload, err := ctrl.cacheSvc.FindVerificationToken(ctx, hashedToken)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusNotFound, err, "")
		return
	}

	// compare the code
	if payload.Code != body.Code {
		response.ResErr(c, ctrl.logger, http.StatusUnauthorized, errors.New("code not match"), "")
		return
	}
	userId, err := uuid.Parse(payload.UserId)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	ctrl.logger.Info(userId)

	// find the user
	user, err := ctrl.userSvc.FindById(userId)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}
	if user == nil {
		response.ResErr(c, ctrl.logger, http.StatusNotFound, errors.New("user not found"), "user not found")
		return
	}
	if user.IsVerified {
		response.ResErr(c, ctrl.logger, http.StatusConflict, errors.New("already verified"), "account has been verified")
		return
	}

	// mark the user as verified
	if err := ctrl.userSvc.VerifyUser(userId); err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	// allow user to login
	// generate access and refresh token
	tokens, err := ctrl.authSvc.CreateAuthTokens(
		ctx,
		user.ID.String(),
		user.JwtVersion,
	)
	if err != nil {
		response.ResErr(c, ctrl.logger, http.StatusInternalServerError, err, "")
		return
	}

	// delete verification data from redis
	if err := ctrl.cacheSvc.DeleteVerificationToken(ctx, hashedToken); err != nil {
		ctrl.logger.Error(err)
	}

	// set response
	c.SetCookie(
		constants.COOKIE_REFRESH_TOKEN,
		tokens.RefreshToken,
		constants.REFRESH_TOKEN_MAX_AGE, "/", "", os.Getenv("GO_ENV") == "production", true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": tokens.AccessToken,
	})
}
