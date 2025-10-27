package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(res http.ResponseWriter, req *http.Request) {
	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid product id", http.StatusBadRequest)
		return
	}

	var product domain.Product
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&product)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid json", http.StatusBadRequest)
		return
	}
	product.ID = id
	_, err = h.service.Update(product)
	if err != nil {
		http.Error(res, "Error updating product", http.StatusInternalServerError)
		return
	}

	utils.SendData(res, "Product updated successfully", http.StatusCreated)
}
