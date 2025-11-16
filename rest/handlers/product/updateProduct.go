package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid product id", http.StatusBadRequest)
		return
	}

	var product domain.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&product)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	product.Id = id
	_, err = handler.productService.Update(product)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, "Product updated successfully", http.StatusCreated)
}
