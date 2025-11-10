package routes

import (
	cartcontroller "kingcom_api/internal/controllers/cart_controller"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/middlewares"
)

type CartRoutes struct {
	handler *lib.RequestHandler
	ctrl    *cartcontroller.CartController
	jwtAuth *middlewares.JwtAuthMiddleware
}

func NewCartRoutes(
	handler *lib.RequestHandler,
	ctrl *cartcontroller.CartController,
	jwtAuth *middlewares.JwtAuthMiddleware,
) *CartRoutes {
	return &CartRoutes{
		handler: handler,
		ctrl:    ctrl,
		jwtAuth: jwtAuth,
	}
}

func (r *CartRoutes) Setup() {
	rtr := r.handler.Gin.Group("/api/cart")
	{
		rtr.GET("", r.jwtAuth.Handler, r.ctrl.FetchCart)
		rtr.POST("/add", r.jwtAuth.Handler, r.ctrl.AddToCart)
	}
}
