package services

import (
	authservice "kingcom_api/internal/services/authService"
	cacheservice "kingcom_api/internal/services/cache_service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(authservice.New),
	fx.Provide(cacheservice.New),
	fx.Provide(NewMailService),
	fx.Provide(NewUserService),
	fx.Provide(NewProductService),
	fx.Provide(NewCartService),
	fx.Provide(NewOrderService),
)
