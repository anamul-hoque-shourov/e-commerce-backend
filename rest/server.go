package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(config config.Config) {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	fmt.Println("Starting", config.ServiceName, "version", config.Version, "on port", config.Port)

	port := strconv.Itoa(config.Port)
	err := http.ListenAndServe(":"+port, wrappedMux)
	if err != nil {
		fmt.Println("Server error", err)
		os.Exit(1)
	}
}
