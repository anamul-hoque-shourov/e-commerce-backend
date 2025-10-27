package user

import (
	"ecommerce/config"
)

type Handler struct {
	config  *config.Config
	service Service
}

func NewHandler(config *config.Config, service Service) *Handler {
	return &Handler{
		config:  config,
		service: service,
	}
}
