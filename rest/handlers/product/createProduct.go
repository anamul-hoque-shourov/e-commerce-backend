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

func (h *Handler) CreateProduct(res http.ResponseWriter, req *http.Request) {
	var requestedProduct RequestCreateProduct
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&requestedProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid json", http.StatusBadRequest)
		return
	}
	createdProduct, err := h.service.Create(domain.Product{
		Title:       requestedProduct.Title,
		Description: requestedProduct.Description,
		Price:       requestedProduct.Price,
		ImageUrl:    requestedProduct.ImageUrl,
	})
	if err != nil {
		http.Error(res, "Error creating product", http.StatusInternalServerError)
		return
	}

	utils.SendData(res, createdProduct, http.StatusCreated)
}
