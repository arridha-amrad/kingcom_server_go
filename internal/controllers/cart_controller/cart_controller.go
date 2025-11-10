package cartcontroller

import (
	"kingcom_api/internal/lib"
	"kingcom_api/internal/services"
)

type CartController struct {
	cartService services.CartService
	validator   *lib.Validator
	logger      *lib.Logger
}

func New(
	cartService services.CartService,
	validator *lib.Validator,
	logger *lib.Logger,
) *CartController {
	return &CartController{
		cartService: cartService,
		validator:   validator,
		logger:      logger,
	}
}
