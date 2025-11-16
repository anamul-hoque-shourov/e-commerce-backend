package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

type Pagination struct {
	Data        []*domain.Product `json:"data"`
	CurrentPage int               `json:"currentPage"`
	Limit       int               `json:"limit"`
	TotalItems  int               `json:"totalItems"`
	TotalPages  int               `json:"totalPages"`
}

func (handler *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	pageQuery := r.URL.Query().Get("page")
	limitQuery := r.URL.Query().Get("limit")

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
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	count, err := handler.productService.Count()
	if err != nil {
		http.Error(w, "Error fetching count", http.StatusInternalServerError)
		return
	}

	paginatedProducts := &Pagination{
		Data:        products,
		Limit:       limit,
		CurrentPage: page,
		TotalItems:  count,
		TotalPages:  count / limit,
	}

	utils.SendData(w, paginatedProducts, http.StatusOK)
}
