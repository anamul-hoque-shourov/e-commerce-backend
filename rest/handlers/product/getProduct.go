package product

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid product id", http.StatusBadRequest)
		return
	}

	product, err := handler.productService.Get(id)
	if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		return
	}
	if product == nil {
		utils.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	utils.SendData(w, product, http.StatusOK)
}
