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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Please send GET request", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(productList)
	w.WriteHeader(http.StatusOK)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Please send POST request", http.StatusBadRequest)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newProduct)

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	json.NewEncoder(w).Encode(newProduct)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-product", createProduct)

	fmt.Println("Server is running on port 8080...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
