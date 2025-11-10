package cart

import "ecommerce/domain"

type Service interface {
	GetByUserId(userId int) (*domain.Cart, error)
	AddItem(userId, productId, quantity int) error
	UpdateItemQuantity(userId, productId, quantity int) error
	RemoveItem(userId, productId int) error
	ClearCart(userId int) error
}
