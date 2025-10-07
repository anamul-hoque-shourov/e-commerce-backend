package product

import (
	"ecommerce/database"
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

	var product database.Product
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&product)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid json", http.StatusBadRequest)
		return
	}
	product.ID = id
	database.Update(product)

	utils.SendData(res, "Product updated successfully", http.StatusCreated)
}
