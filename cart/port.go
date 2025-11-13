package cart

import (
	"ecommerce/domain"
	"ecommerce/rest/handlers/cart"
)

type CartService interface {
	cart.CartService
}

type CartRepo interface {
	GetByUserId(userId int) (*domain.Cart, error)
	AddItem(userId, productId, quantity int) error
	UpdateItemQuantity(userId, productId, quantity int) error
	RemoveItem(userId, productId int) error
	ClearCart(userId int) error
}
