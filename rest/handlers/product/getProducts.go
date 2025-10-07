package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProducts(res http.ResponseWriter, req *http.Request) {
	utils.SendData(res, database.List(), http.StatusOK)
}
