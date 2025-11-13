package product

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (handler *ProductHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.MiddlewareManager) {
	mux.Handle("GET /products",
		manager.CustomManager(
			http.HandlerFunc(handler.GetProducts),
		),
	)
	mux.Handle("POST /products",
		manager.CustomManager(
			http.HandlerFunc(handler.CreateProduct),
			handler.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle("GET /products/{id}",
		manager.CustomManager(
			http.HandlerFunc(handler.GetProduct),
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.CustomManager(
			http.HandlerFunc(handler.UpdateProduct),
			handler.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle("DELETE /products/{id}",
		manager.CustomManager(
			http.HandlerFunc(handler.DeleteProduct),
			handler.middlewares.AuthenticateJWT,
		),
	)
}
