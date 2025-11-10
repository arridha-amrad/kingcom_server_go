package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDatabase),
	fx.Provide(NewEnv),
	fx.Provide(NewRequestHandler),
	fx.Provide(GetLogger),
	fx.Provide(NewRedis),
	fx.Provide(NewValidator),
)
