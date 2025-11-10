package cart

import (
	"ecommerce/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	service     Service
}

func NewHandler(
	middlewares *middlewares.Middlewares,
	service Service,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		service:     service,
	}
}
