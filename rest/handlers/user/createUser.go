package user

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	newUser = newUser.Store()

	utils.SendData(res, newUser, http.StatusCreated)
}
