package product

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (handler *ProductHandler) DeleteProduct(res http.ResponseWriter, req *http.Request) {
	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid product id", http.StatusBadRequest)
		return
	}

	err = handler.productService.Delete(id)
	if err != nil {
		http.Error(res, "Error deleting product", http.StatusInternalServerError)
		return
	}

	utils.SendData(res, "Product deleted successfully", http.StatusOK)
}
