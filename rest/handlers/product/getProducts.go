package product

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProducts(res http.ResponseWriter, req *http.Request) {
	products, err := h.service.List()
	if err != nil {
		http.Error(res, "Error fetching products", http.StatusInternalServerError)
		return
	}
	utils.SendData(res, products, http.StatusOK)
}
