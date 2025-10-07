package product

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)
	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			middlewares.AuthenticateJWT,
		),
	)
	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			middlewares.AuthenticateJWT,
		),
	)
	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			middlewares.AuthenticateJWT,
		),
	)
}
