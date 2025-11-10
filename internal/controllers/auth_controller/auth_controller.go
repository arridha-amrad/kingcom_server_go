package authcontroller

import (
	"kingcom_api/internal/lib"
	"kingcom_api/internal/services"
	authservice "kingcom_api/internal/services/authService"
	cacheservice "kingcom_api/internal/services/cache_service"
)

type AuthController struct {
	logger   *lib.Logger
	validate *lib.Validator
	authSvc  authservice.AuthService
	mailSvc  services.MailService
	cacheSvc cacheservice.CacheService
	userSvc  services.UserService
}

func New(
	logger *lib.Logger,
	validate *lib.Validator,
	authSvc authservice.AuthService,
	mailSvc services.MailService,
	userSvc services.UserService,
	cacheSvc cacheservice.CacheService,
) *AuthController {
	return &AuthController{
		logger:   logger,
		validate: validate,
		authSvc:  authSvc,
		mailSvc:  mailSvc,
		userSvc:  userSvc,
		cacheSvc: cacheSvc,
	}
}
