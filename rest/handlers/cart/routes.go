package cart

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /cart",
		manager.With(
			http.HandlerFunc(h.AddToCart),
		),
	)
}
