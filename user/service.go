package user

import "ecommerce/domain"

type userService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (service *userService) Create(user domain.User) (*domain.User, error) {
	return service.userRepo.Create(user)
}

func (service *userService) Get(email, password string) (*domain.User, error) {
	return service.userRepo.Get(email, password)
}
