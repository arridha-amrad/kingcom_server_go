package controllers

import (
	authcontroller "kingcom_api/internal/controllers/auth_controller"
	cartcontroller "kingcom_api/internal/controllers/cart_controller"
	orderController "kingcom_api/internal/controllers/order_controller"
	productcontroller "kingcom_api/internal/controllers/product_controller"
	shippingcontroller "kingcom_api/internal/controllers/shipping_controller"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(authcontroller.New),
	fx.Provide(productcontroller.New),
	fx.Provide(shippingcontroller.New),
	fx.Provide(cartcontroller.New),
	fx.Provide(orderController.New),
)
