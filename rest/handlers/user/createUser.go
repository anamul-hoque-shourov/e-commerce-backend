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

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestedUser RequestCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestedUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newUser, err := handler.userService.Create(domain.User{
		FirstName:   requestedUser.FirstName,
		LastName:    requestedUser.LastName,
		Email:       requestedUser.Email,
		Password:    requestedUser.Password,
		IsShopOwner: requestedUser.IsShopOwner,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, newUser, http.StatusCreated)
}
