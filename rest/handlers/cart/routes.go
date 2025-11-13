package cart

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (handler *CartHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.MiddlewareManager) {
	mux.Handle("POST /cart",
		manager.CustomManager(
			http.HandlerFunc(handler.AddToCart),
		),
	)
}
