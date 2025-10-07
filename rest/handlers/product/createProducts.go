package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(res http.ResponseWriter, req *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid json", http.StatusBadRequest)
		return
	}
	newProduct = database.Store(newProduct)

	utils.SendData(res, newProduct, http.StatusCreated)
}
