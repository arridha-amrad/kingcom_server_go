package repositories

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewCacheRepository),
	fx.Provide(NewProductRepository),
	fx.Provide(NewCartRepository),
	fx.Provide(NewOrderRepository),
)
