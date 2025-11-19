package product

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (handler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	productId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid product id", http.StatusBadRequest)
		return
	}

	err = handler.productService.Delete(productId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, "Product deleted successfully", http.StatusOK)
}
