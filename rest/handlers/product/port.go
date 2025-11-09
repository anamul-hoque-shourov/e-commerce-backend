package product

import "ecommerce/domain"

type Service interface {
	Create(product domain.Product) (*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	List(page int, limit int) ([]*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(productId int) error
}
