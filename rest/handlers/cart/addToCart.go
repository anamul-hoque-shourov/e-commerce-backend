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

func (handler *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var userRequest UserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userRequest)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	err = handler.cartService.AddItem(userRequest.UserId, userRequest.ProductId, userRequest.Quantity)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, "Product added to cart", http.StatusCreated)
}
