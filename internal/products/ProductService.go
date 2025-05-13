package products

import "context"

type ProductService struct {
	repo IProductRepository
}

func NewProductService(repo IProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) AddProduct(ctx context.Context, product Product) error {
	return s.repo.AddProduct(ctx, product)
}
