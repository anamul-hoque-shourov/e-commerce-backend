package cart

import "ecommerce/domain"

type cartService struct {
	cartRepo CartRepo
}

func NewCartService(cartRepo CartRepo) CartService {
	return &cartService{
		cartRepo: cartRepo,
	}
}

func (service *cartService) GetByUserId(userId int) (*domain.Cart, error) {
	return service.cartRepo.GetByUserId(userId)
}

func (service *cartService) AddItem(userId, productId, quantity int) error {
	return service.cartRepo.AddItem(userId, productId, quantity)
}

func (service *cartService) UpdateItemQuantity(userId, productId, quantity int) error {
	return service.cartRepo.UpdateItemQuantity(userId, productId, quantity)
}

func (service *cartService) RemoveItem(userId, productId int) error {
	return service.cartRepo.RemoveItem(userId, productId)
}

func (service *cartService) ClearCart(userId int) error {
	return service.cartRepo.ClearCart(userId)
}
