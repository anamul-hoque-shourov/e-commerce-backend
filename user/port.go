package user

import (
	"ecommerce/domain"
	"ecommerce/rest/handlers/user"
)

type Service interface {
	user.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Get(email string, password string) (*domain.User, error)
}
