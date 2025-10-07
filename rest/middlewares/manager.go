package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (manager *Manager) Use(middlewares ...Middleware) {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
}

func (manager *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	next := handler
	for _, middleware := range middlewares {
		next = middleware(next)
	}
	return next
}

func (manager *Manager) WrapMux(mux http.Handler) http.Handler {
	next := mux
	for _, middleware := range manager.globalMiddlewares {
		next = middleware(next)
	}
	return next
}
