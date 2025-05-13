package products

import (
	"context"
	"database/sql"
)

type ProductRepositoryPostgres struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) IProductRepository {
	return &ProductRepositoryPostgres{db: db}
}

func (p *ProductRepositoryPostgres) AddProduct(ctx context.Context, product Product) error {
	print("Test")
	return nil
}
