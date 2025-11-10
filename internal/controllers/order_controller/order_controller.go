package ordercontroller

import (
	"kingcom_api/internal/lib"
	"kingcom_api/internal/services"
)

type OrderController struct {
	validator      *lib.Validator
	logger         *lib.Logger
	orderService   services.OrderService
	productService services.ProductService
}

func New(
	validator *lib.Validator,
	logger *lib.Logger,
	orderService services.OrderService,
	productService services.ProductService,
) *OrderController {
	return &OrderController{
		validator:      validator,
		logger:         logger,
		orderService:   orderService,
		productService: productService,
	}
}
