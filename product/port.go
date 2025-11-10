package product

import (
	"ecommerce/domain"
	"ecommerce/rest/handlers/product"
)

type Service interface {
	product.Service
}

type ProductRepo interface {
	Create(product domain.Product) (*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(productId int) error
}
