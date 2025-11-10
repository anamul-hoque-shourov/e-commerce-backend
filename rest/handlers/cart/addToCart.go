package cart

import (
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserRequest struct {
	UserId    int `json:"userId"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

func (h *Handler) AddToCart(res http.ResponseWriter, req *http.Request) {
	var userRequest UserRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&userRequest)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Please provide valid json", http.StatusBadRequest)
		return
	}
	err = h.service.AddItem(userRequest.UserId, userRequest.ProductId, userRequest.Quantity)
	if err != nil {
		http.Error(res, "Error creating product", http.StatusInternalServerError)
		return
	}

	utils.SendData(res, "Product added to cart", http.StatusCreated)
}
