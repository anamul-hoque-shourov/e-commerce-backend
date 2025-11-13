package user

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (handler *UserHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.MiddlewareManager) {
	mux.Handle("POST /users",
		manager.CustomManager(
			http.HandlerFunc(handler.CreateUser),
		),
	)
	mux.Handle("POST /users/login",
		manager.CustomManager(
			http.HandlerFunc(handler.Login),
		),
	)
}
