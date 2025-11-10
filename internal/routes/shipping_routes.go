package routes

import (
	shippingcontroller "kingcom_api/internal/controllers/shipping_controller"
	"kingcom_api/internal/lib"
)

type ShippingRoutes struct {
	handler *lib.RequestHandler
	ctrl    *shippingcontroller.ShippingController
}

func NewShippingRoutes(
	handler *lib.RequestHandler,
	ctrl *shippingcontroller.ShippingController,
) *ShippingRoutes {
	return &ShippingRoutes{
		handler: handler,
		ctrl:    ctrl,
	}
}

func (r *ShippingRoutes) Setup() {
	rtr := r.handler.Gin.Group("/api/shipping")
	{
		rtr.GET("/provinces", r.ctrl.FetchProvinces)
		rtr.GET("/cities/:provinceID", r.ctrl.FetchCities)
		rtr.GET("/districts/:cityID", r.ctrl.GetDistricts)
		rtr.POST("/cost", r.ctrl.CalcCost)
	}
}
