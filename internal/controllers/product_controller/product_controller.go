package productcontroller

import (
	"kingcom_api/internal/lib"
	"kingcom_api/internal/services"
)

type ProductController struct {
	validator      *lib.Validator
	logger         *lib.Logger
	userService    services.UserService
	productService services.ProductService
}

func New(
	validator *lib.Validator,
	logger *lib.Logger,
	userService services.UserService,
	productService services.ProductService,
) *ProductController {
	return &ProductController{
		validator:      validator,
		logger:         logger,
		userService:    userService,
		productService: productService,
	}
}
