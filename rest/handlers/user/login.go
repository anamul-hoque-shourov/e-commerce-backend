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

func (h *Handler) Login(res http.ResponseWriter, req *http.Request) {
	var loginReq LoginRequest
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.service.Get(loginReq.Email, loginReq.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Could not login", http.StatusInternalServerError)
		return
	}
	if user == nil {
		utils.SendError(res, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.CreateJwt(h.config.JwtSecret, utils.Payload{
		ID:          user.ID,
		FistName:    user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Could not create JWT", http.StatusInternalServerError)
		return
	}

	utils.SendData(res, accessToken, http.StatusOK)
}
