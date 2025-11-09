package product

import "ecommerce/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (s *service) List(page int, limit int) ([]*domain.Product, error) {
	return s.productRepo.List(page, limit)
}

func (s *service) Get(productId int) (*domain.Product, error) {
	return s.productRepo.Get(productId)
}

func (s *service) Create(product domain.Product) (*domain.Product, error) {
	return s.productRepo.Create(product)
}

func (s *service) Update(product domain.Product) (*domain.Product, error) {
	return s.productRepo.Update(product)
}

func (s *service) Delete(productId int) error {
	return s.productRepo.Delete(productId)
}
