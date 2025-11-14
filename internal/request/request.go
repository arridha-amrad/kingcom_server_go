package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"kingcom_api/internal/constants"
	"kingcom_api/internal/lib"
	authservice "kingcom_api/internal/services/authService"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func GetBody[T any](c *gin.Context, validate *lib.Validator) (*T, map[string]string) {
	var input T
	msgErrors := make(map[string]string)
	if err := c.ShouldBindJSON(&input); err != nil {
		msgErrors["bind"] = err.Error()
		return nil, msgErrors
	}
	if err := validate.Struct(input); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				message := lib.Messages[e.Tag()]
				if strings.Contains(message, "%s") {
					msgErrors[strings.ToLower(e.Field())] = fmt.Sprintf(message, e.Param())
				} else {
					msgErrors[strings.ToLower(e.Field())] = message
				}
			}
		}
		return nil, msgErrors
	}
	return &input, nil
}

func PrintBody(logger *lib.Logger, input any) {
	jsonBody, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		logger.Error("failed to marshal body", zap.Error(err))
		return
	}
	logger.Info("CreateOrder body", zap.String("body", string(jsonBody)))
}

func ExtractAccessTokenPayload(c *gin.Context) (*authservice.JWTPayload, error) {
	raw, exists := c.Get(constants.ACCESS_TOKEN_PAYLOAD)
	if !exists {
		return nil, errors.New("access token payload not found")
	}
	payload, ok := raw.(*authservice.JWTPayload)
	if !ok {
		return nil, errors.New("access token payload has invalid type")
	}
	return payload, nil
}

func GetRefreshToken(c *gin.Context) (string, error) {
	rawRefreshToken, err := c.Cookie(constants.COOKIE_REFRESH_TOKEN)
	if err != nil {
		return "", errors.New("refresh token is missing")
	}
	return rawRefreshToken, nil
}
