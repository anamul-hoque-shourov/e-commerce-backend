package user

import (
	"ecommerce/config"
)

type UserHandler struct {
	config  *config.Config
	userService UserService
}

func NewUserHandler(config *config.Config, userService UserService) *UserHandler {
	return &UserHandler{
		config:  config,
		userService: userService,
	}
}
