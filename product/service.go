package product

import "ecommerce/domain"

type productService struct {
	productRepo ProductRepo
}

func NewProductService(productRepo ProductRepo) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (service *productService) List(page, limit int) ([]*domain.Product, error) {
	return service.productRepo.List(page, limit)
}

func (service *productService) Get(productId int) (*domain.Product, error) {
	return service.productRepo.Get(productId)
}

func (service *productService) Create(product domain.Product) (*domain.Product, error) {
	return service.productRepo.Create(product)
}

func (service *productService) Update(product domain.Product) (*domain.Product, error) {
	return service.productRepo.Update(product)
}

func (service *productService) Delete(productId int) error {
	return service.productRepo.Delete(productId)
}
