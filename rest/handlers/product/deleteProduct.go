package product

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(res http.ResponseWriter, req *http.Request) {
	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid product id", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(res, "Error deleting product", http.StatusInternalServerError)
		return
	}

	utils.SendData(res, "Product deleted successfully", http.StatusOK)
}
