package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

type PaginatedData struct {
	Data       []*domain.Product `json:"data"`
	Pagination *Pagination       `json:"pagination"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	Limit       int `json:"limit"`
	TotalItems  int `json:"totalItems"`
	TotalPages  int `json:"totalPages"`
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

	totalPages := (count + limit - 1) / limit
	if totalPages == 0 {
		totalPages = 1
	}

	paginatedData := &PaginatedData{
		Data: products,
		Pagination: &Pagination{
			Limit:       limit,
			CurrentPage: page,
			TotalItems:  count,
			TotalPages:  totalPages,
		},
	}

	utils.SendData(w, paginatedData, http.StatusOK)
}
