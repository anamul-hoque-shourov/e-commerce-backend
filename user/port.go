package user

import (
	"ecommerce/domain"
	"ecommerce/rest/handlers/user"
)

type UserService interface {
	user.UserService
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Get(email, password string) (*domain.User, error)
}
