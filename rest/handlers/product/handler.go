package product

import (
	"ecommerce/rest/middlewares"
)

type ProductHandler struct {
	middlewares    *middlewares.Middlewares
	productService ProductService
}

func NewProductHandler(
	middlewares *middlewares.Middlewares,
	productService ProductService,
) *ProductHandler {
	return &ProductHandler{
		middlewares:    middlewares,
		productService: productService,
	}
}
