package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestCreateProduct struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var requestedProduct RequestCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestedProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	createdProduct, err := handler.productService.Create(domain.Product{
		Title:       requestedProduct.Title,
		Description: requestedProduct.Description,
		Price:       requestedProduct.Price,
		ImageUrl:    requestedProduct.ImageUrl,
	})
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, createdProduct, http.StatusCreated)
}
