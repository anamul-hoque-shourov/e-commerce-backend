package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(res http.ResponseWriter, req *http.Request) {
	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid product id", http.StatusBadRequest)
		return
	}

	product := database.Get(id)
	if product == nil {
		utils.SendError(res, "Product not found", http.StatusNotFound)
		return
	}

	utils.SendData(res, product, http.StatusOK)
}
