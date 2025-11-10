package cart

import "ecommerce/domain"

type service struct {
	cartRepo CartRepo
}

func NewService(cartRepo CartRepo) Service {
	return &service{
		cartRepo: cartRepo,
	}
}

func (s *service) GetByUserId(userId int) (*domain.Cart, error) {
	return s.cartRepo.GetByUserId(userId)
}

func (s *service) AddItem(userId, productId, quantity int) error {
	return s.cartRepo.AddItem(userId, productId, quantity)
}

func (s *service) UpdateItemQuantity(userId, productId, quantity int) error {
	return s.cartRepo.UpdateItemQuantity(userId, productId, quantity)
}

func (s *service) RemoveItem(userId, productId int) error {
	return s.cartRepo.RemoveItem(userId, productId)
}

func (s *service) ClearCart(userId int) error {
	return s.cartRepo.ClearCart(userId)
}
