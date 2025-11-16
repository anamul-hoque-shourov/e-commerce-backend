package user

import (
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := handler.userService.Get(loginReq.Email, loginReq.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not login", http.StatusInternalServerError)
		return
	}
	if user == nil {
		utils.SendError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.GenerateToken(handler.config.JwtSecret, utils.Payload{
		Id:          user.Id,
		FistName:    user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not create JWT", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, accessToken, http.StatusOK)
}
