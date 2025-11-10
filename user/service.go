package user

import "ecommerce/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Create(user domain.User) (*domain.User, error) {
	return s.userRepo.Create(user)
}

func (s *service) Get(email, password string) (*domain.User, error) {
	return s.userRepo.Get(email, password)
}
