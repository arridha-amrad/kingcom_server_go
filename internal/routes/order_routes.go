package routes

import (
	orderController "kingcom_api/internal/controllers/order_controller"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/middlewares"
)

type OrderRoutes struct {
	handler *lib.RequestHandler
	ctrl    *orderController.OrderController
	jwtAuth *middlewares.JwtAuthMiddleware
}

func NewOrderRoutes(
	handler *lib.RequestHandler,
	ctrl *orderController.OrderController,
	jwtAuth *middlewares.JwtAuthMiddleware,
) *OrderRoutes {
	return &OrderRoutes{
		handler: handler,
		ctrl:    ctrl,
		jwtAuth: jwtAuth,
	}
}

func (o *OrderRoutes) Setup() {
	rtr := o.handler.Gin.Group("/api/order")
	{
		rtr.POST("", o.jwtAuth.Handler, o.ctrl.PlaceOrder)
		rtr.GET("", o.jwtAuth.Handler, o.ctrl.FetchUserOrders)
	}
}
