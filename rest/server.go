package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/cart"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	config         *config.Config
	userHandler    *user.UserHandler
	productHandler *product.ProductHandler
	cartHandler    *cart.CartHandler
}

func NewServer(
	config *config.Config,
	userHandler *user.UserHandler,
	productHandler *product.ProductHandler,
	cartHandler *cart.CartHandler,
) *Server {
	return &Server{
		config:         config,
		userHandler:    userHandler,
		productHandler: productHandler,
		cartHandler:    cartHandler,
	}
}

func (server *Server) Start() {
	middlewareManager := middlewares.NewMiddlewareManager()
	middlewareManager.GlobalManager(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := middlewareManager.WrapMux(mux)

	server.userHandler.RegisterRoutes(mux, middlewareManager)
	server.productHandler.RegisterRoutes(mux, middlewareManager)
	server.cartHandler.RegisterRoutes(mux, middlewareManager)

	fmt.Println("Starting", server.config.ServiceName, "version", server.config.Version, "on port", server.config.Port)

	port := strconv.Itoa(server.config.Port)
	err := http.ListenAndServe(":"+port, wrappedMux)
	if err != nil {
		fmt.Println("Server error", err)
		os.Exit(1)
	}
}
