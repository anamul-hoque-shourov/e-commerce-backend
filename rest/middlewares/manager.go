package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type MiddlewareManager struct {
	globalMiddlewares []Middleware
}

func NewMiddlewareManager() *MiddlewareManager {
	return &MiddlewareManager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (manager *MiddlewareManager) GlobalManager(middlewares ...Middleware) {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
}

func (manager *MiddlewareManager) CustomManager(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func (manager *MiddlewareManager) WrapMux(mux http.Handler) http.Handler {
	for _, middleware := range manager.globalMiddlewares {
		mux = middleware(mux)
	}
	return mux
}
