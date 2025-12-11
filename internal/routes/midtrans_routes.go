package routes

import (
	"kingcom_api/internal/controllers"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/middlewares"
)

type MidtransRoutes struct {
	handler *lib.RequestHandler
	ctrl    *controllers.MidtransController
	jwtAuth *middlewares.JwtAuthMiddleware
}

func NewMidtransRoutes(
	handler *lib.RequestHandler,
	ctrl *controllers.MidtransController,
	jwtAuth *middlewares.JwtAuthMiddleware,
) *MidtransRoutes {
	return &MidtransRoutes{
		handler: handler,
		ctrl:    ctrl,
		jwtAuth: jwtAuth,
	}
}

func (r *MidtransRoutes) Setup() {
	rtr := r.handler.Gin.Group("/api/midtrans")
	{
		rtr.POST("/notification", r.ctrl.HandleNotification)
		rtr.POST("/token", r.ctrl.CreateTransactionToken)
	}
}
