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
	usr, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}

func (s *service) Get(email string, password string) (*domain.User, error) {
	usr, err := s.userRepo.Get(email, password)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
