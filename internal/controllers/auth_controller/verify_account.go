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
	res := response.New(c, ctrl.logger)

	body, errValidation := request.GetBody[dto.VerifyNewAccount](c, ctrl.validate)
	if errValidation != nil {
		res.ResErrValidation(errValidation)
		return
	}

	hashedToken := utils.HashWithSHA256(body.Token)
	ctx := c.Request.Context()

	// find verification data in redis
	payload, err := ctrl.cacheSvc.FindVerificationToken(ctx, hashedToken)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	// compare the code
	if payload.Code != body.Code {
		res.ResErrUnauthorized(errors.New("invalid code"))
		return
	}
	userId, err := uuid.Parse(payload.UserId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}

	ctrl.logger.Info(userId)

	// find the user
	user, err := ctrl.userSvc.FindById(userId)
	if err != nil {
		res.ResInternalServerErr(err)
		return
	}
	if user == nil {
		res.ResInternalServerErr(err)
		return
	}
	if user.IsVerified {
		err := errors.New("unverified account")
		res.ResErr(http.StatusConflict, err, err.Error())
		return
	}

	// mark the user as verified
	if err := ctrl.userSvc.VerifyUser(userId); err != nil {
		res.ResInternalServerErr(err)
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
		res.ResInternalServerErr(err)
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

	res.ResOk(gin.H{
		"user":  user,
		"token": tokens.AccessToken,
	})
}
