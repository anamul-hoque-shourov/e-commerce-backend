package utils

import (
	"net/http"
)

type PaginatedData struct {
	Data       any         `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	Limit       int `json:"limit"`
	TotalItems  int `json:"totalItems"`
	TotalPages  int `json:"totalPages"`
}

func SendPage(w http.ResponseWriter, data any, page int, limit int, count int) {
	totalPages := (count + limit - 1) / limit
	if totalPages == 0 {
		totalPages = 1
	}

	paginatedData := &PaginatedData{
		Data: data,
		Pagination: &Pagination{
			Limit:       limit,
			CurrentPage: page,
			TotalItems:  count,
			TotalPages:  totalPages,
		},
	}
	SendData(w, paginatedData, http.StatusOK)
}
