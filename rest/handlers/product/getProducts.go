package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (handler *ProductHandler) GetProducts(res http.ResponseWriter, req *http.Request) {

	pageQuery := req.URL.Query().Get("page")
	limitQuery := req.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageQuery)
	limit, _ := strconv.Atoi(limitQuery)

	if page <= 0 {
		page = 1
	}

	if limit <= 10 {
		limit = 10
	}

	products, err := handler.productService.List(page, limit)
	if err != nil {
		http.Error(res, "Error fetching products", http.StatusInternalServerError)
		return
	}
	utils.SendData(res, products, http.StatusOK)
}
