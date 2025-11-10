package modules

import (
	"context"
	"fmt"
	"kingcom_api/internal/controllers"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/middlewares"
	"kingcom_api/internal/repositories"
	"kingcom_api/internal/routes"
	"kingcom_api/internal/services"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var CommonModule = fx.Options(
	controllers.Module,
	services.Module,
	repositories.Module,
	lib.Module,
	middlewares.Module,
	routes.Module,
	fx.WithLogger(func() fxevent.Logger {
		return lib.GetLogger().GetFxLogger()
	}),
	fx.Invoke(registerHooks),
)

func registerHooks(
	lifecycle fx.Lifecycle,
	h *lib.RequestHandler,
	env *lib.Env,
	logger *lib.Logger,
	routes *routes.Routes,
	middleware *middlewares.Middlewares,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				middleware.Setup()
				routes.Setup()
				var port = ":" + env.ServerPort
				go h.Gin.Run(port)
				logger.Info(fmt.Sprintf("== Starting application in %s ==", env.AppUrl))
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Stopping application")
				return nil
			},
		},
	)
}
