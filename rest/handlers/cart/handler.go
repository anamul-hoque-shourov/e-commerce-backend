package cart

import (
	"ecommerce/rest/middlewares"
)

type CartHandler struct {
	middlewares *middlewares.Middlewares
	cartService CartService
}

func NewCartHandler(
	middlewares *middlewares.Middlewares,
	cartService CartService,
) *CartHandler {
	return &CartHandler{
		middlewares: middlewares,
		cartService: cartService,
	}
}
