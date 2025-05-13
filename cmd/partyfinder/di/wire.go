//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"

	"github.com/Oskarbayy/partyfinder-backend/internal/products"
	"github.com/google/wire"
)

func MapProductHandler(db *sql.DB) (*products.ProductHandler, error) {
	wire.Build(products.NewProductRepository, products.NewProductService, products.NewProductHandler)
	return nil, nil
}
