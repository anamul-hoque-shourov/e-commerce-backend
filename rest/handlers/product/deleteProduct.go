package product

import (
	"ecommerce/database"
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

	database.Delete(id)

	utils.SendData(res, "Product deleted successfully", http.StatusOK)
}
