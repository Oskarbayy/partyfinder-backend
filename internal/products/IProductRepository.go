package products

import "context"

type IProductRepository interface {
	AddProduct(ctx context.Context, product Product) error
}
