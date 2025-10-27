package rest

import (
	"ecommerce/config"
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
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	config *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		config:         config,
		productHandler: productHandler,
		userHandler:    userHandler,
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

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	fmt.Println("Starting", server.config.ServiceName, "version", server.config.Version, "on port", server.config.Port)

	port := strconv.Itoa(server.config.Port)
	err := http.ListenAndServe(":"+port, wrappedMux)
	if err != nil {
		fmt.Println("Server error", err)
		os.Exit(1)
	}
}
