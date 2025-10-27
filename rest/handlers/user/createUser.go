package user

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestCreateUser struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

func (h *Handler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var requestedUser RequestCreateUser
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&requestedUser)

	if err != nil {
		fmt.Println(err)
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	newUser, err := h.service.Create(domain.User{
		FirstName:   requestedUser.FirstName,
		LastName:    requestedUser.LastName,
		Email:       requestedUser.Email,
		Password:    requestedUser.Password,
		IsShopOwner: requestedUser.IsShopOwner,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(res, "Could not create user", http.StatusInternalServerError)
		return
	}

	utils.SendData(res, newUser, http.StatusCreated)
}
