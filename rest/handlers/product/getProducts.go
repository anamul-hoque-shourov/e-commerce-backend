package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"net/http"
	"strconv"
	"sync"
)

var products []*domain.Product
var count int

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

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		prdcts, err := handler.productService.List(page, limit)
		if err != nil {
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			return
		}
		products = prdcts
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cnt, err := handler.productService.Count()
		if err != nil {
			http.Error(w, "Error fetching count", http.StatusInternalServerError)
			return
		}
		count = cnt
	}()

	wg.Wait()
	
	utils.SendPage(w, products, page, limit, count)
}
