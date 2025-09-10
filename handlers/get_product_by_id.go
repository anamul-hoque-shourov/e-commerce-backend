package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")

	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please provide valid product id", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			utils.SendData(w, product, http.StatusOK)
			return
		}
	}

	utils.SendData(w, "No data found", http.StatusNotFound)
}
