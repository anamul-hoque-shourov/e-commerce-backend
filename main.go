package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

var productList []Product

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Apple iPhone 14",
		Description: "The latest iPhone with advanced features.",
		Price:       999.99,
		ImageUrl:    "https://example.com/iphone14.jpg",
	}
	product2 := Product{
		ID:          2,
		Title:       "Samsung Galaxy S22",
		Description: "A powerful smartphone with a stunning display.",
		Price:       899.99,
		ImageUrl:    "https://example.com/galaxys22.jpg",
	}
	product3 := Product{
		ID:          3,
		Title:       "Google Pixel 6",
		Description: "Experience the best of Google with the Pixel 6.",
		Price:       599.99,
		ImageUrl:    "https://example.com/pixel6.jpg",
	}
	product4 := Product{
		ID:          4,
		Title:       "OnePlus 9",
		Description: "Flagship performance at an affordable price.",
		Price:       729.99,
		ImageUrl:    "https://example.com/oneplus9.jpg",
	}

	productList = append(productList, product1)
	productList = append(productList, product2)
	productList = append(productList, product3)
	productList = append(productList, product4)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(productList)
	w.WriteHeader(http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newProduct)

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	json.NewEncoder(w).Encode(newProduct)
	w.WriteHeader(http.StatusCreated)
}

func globalRouter(mux *http.ServeMux) http.Handler {
	handleGlobalRoutes := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleGlobalRoutes)
}

func main() {
	mux := http.NewServeMux()
	globalRouter := globalRouter(mux)

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-product", http.HandlerFunc(createProduct))

	fmt.Println("Server is running on port 8080...")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
