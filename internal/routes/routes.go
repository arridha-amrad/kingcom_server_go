package routes

import (
	"go.uber.org/fx"
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	authRoutes *AuthRoutes,
	prodRoutes *ProductRoutes,
	shippingRoutes *ShippingRoutes,
	cartRoutes *CartRoutes,
	orderRoutes *OrderRoutes,
	midtransRoutes *MidtransRoutes,
) *Routes {
	return &Routes{
		authRoutes,
		prodRoutes,
		shippingRoutes,
		cartRoutes,
		orderRoutes,
		midtransRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

var Module = fx.Options(
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewProductRoutes),
	fx.Provide(NewShippingRoutes),
	fx.Provide(NewCartRoutes),
	fx.Provide(NewOrderRoutes),
	fx.Provide(NewMidtransRoutes),
	fx.Provide(NewRoutes),
)
