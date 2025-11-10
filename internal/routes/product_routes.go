package routes

import (
	productcontroller "kingcom_api/internal/controllers/product_controller"
	"kingcom_api/internal/lib"
	"kingcom_api/internal/middlewares"
)

type ProductRoutes struct {
	logger  *lib.Logger
	handler *lib.RequestHandler
	ctrl    *productcontroller.ProductController
	jwtAuth *middlewares.JwtAuthMiddleware
}

func NewProductRoutes(
	logger *lib.Logger,
	handler *lib.RequestHandler,
	ctrl *productcontroller.ProductController,
	jwtAuth *middlewares.JwtAuthMiddleware,
) *ProductRoutes {
	return &ProductRoutes{
		logger:  logger,
		handler: handler,
		ctrl:    ctrl,
		jwtAuth: jwtAuth,
	}
}

func (r *ProductRoutes) Setup() {
	rtr := r.handler.Gin.Group("/api/products")
	{
		rtr.POST("", r.jwtAuth.Handler, r.ctrl.CreateProduct)
		rtr.GET("", r.ctrl.FetchAllProducts)
		rtr.GET(":slug", r.ctrl.FetchOneWithDetails)
	}
}
