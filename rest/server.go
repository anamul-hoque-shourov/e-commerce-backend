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
	userHandler    *user.Handler
	productHandler *product.Handler
	cartHandler    *cart.Handler
}

func NewServer(
	config *config.Config,
	userHandler *user.Handler,
	productHandler *product.Handler,
	cartHandler *cart.Handler,
) *Server {
	return &Server{
		config:         config,
		userHandler:    userHandler,
		productHandler: productHandler,
		cartHandler:    cartHandler,
	}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.userHandler.RegisterRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)
	server.cartHandler.RegisterRoutes(mux, manager)

	fmt.Println("Starting", server.config.ServiceName, "version", server.config.Version, "on port", server.config.Port)

	port := strconv.Itoa(server.config.Port)
	err := http.ListenAndServe(":"+port, wrappedMux)
	if err != nil {
		fmt.Println("Server error", err)
		os.Exit(1)
	}
}
