package user

import "ecommerce/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Get(email string, password string) (*domain.User, error)
}

func NewService() {}
