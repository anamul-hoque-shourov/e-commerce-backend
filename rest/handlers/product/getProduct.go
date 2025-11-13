package product

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (handler *ProductHandler) GetProduct(res http.ResponseWriter, req *http.Request) {
	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid product id", http.StatusBadRequest)
		return
	}

	product, err := handler.productService.Get(id)
	if err != nil {
		http.Error(res, "Error fetching product", http.StatusInternalServerError)
		return
	}
	if product == nil {
		utils.SendError(res, "Product not found", http.StatusNotFound)
		return
	}

	utils.SendData(res, product, http.StatusOK)
}
