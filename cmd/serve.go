package cmd

import (
	"ecommerce/handlers"
	"ecommerce/routers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	globalRouter := routers.GlobalRouter(mux)

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductById))

	fmt.Println("Server is running on port 8080...")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
