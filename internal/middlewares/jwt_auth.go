package middlewares

import (
	"errors"
	"kingcom_api/internal/constants"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/response"
	authservice "kingcom_api/internal/services/authService"
	cacheservice "kingcom_api/internal/services/cache_service"
	"strings"

	"github.com/gin-gonic/gin"
)

type JwtAuthMiddleware struct {
	logger   *lib.Logger
	jwtSvc   authservice.AuthService
	cacheSvc cacheservice.CacheService
}

func NewJwtAuthMiddleware(
	logger *lib.Logger,
	jwtSvc authservice.AuthService,
	cacheSvc cacheservice.CacheService,
) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		logger:   logger,
		jwtSvc:   jwtSvc,
		cacheSvc: cacheSvc,
	}
}

func (m *JwtAuthMiddleware) Setup() {}

func (m *JwtAuthMiddleware) Handler(c *gin.Context) {
	res := response.New(c, m.logger)

	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		err := errors.New("authorization header is empty")
		res.ResErrUnauthorized(err)
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authorization, bearerPrefix) {
		err := errors.New("invalid bearer format")
		res.ResErrUnauthorized(err)
		return
	}

	tokenStr := strings.TrimSpace(strings.TrimPrefix(authorization, bearerPrefix))

	payload, err := m.jwtSvc.VerifyJwt(tokenStr)
	if err != nil {
		res.ResErrUnauthorized(err)
		return
	}

	if _, err := m.cacheSvc.FindAccessToken(c.Request.Context(), payload.Jti); err != nil {
		res.ResInternalServerErr(err)
		return
	}

	c.Set(constants.ACCESS_TOKEN_PAYLOAD, payload)

	c.Next()
}
