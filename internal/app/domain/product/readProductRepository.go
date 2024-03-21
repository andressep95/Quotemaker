package domain

import (
	"context"
)

type ReadProductRepository interface {
	GetProductByID(ctx context.Context, id int) (*Product, error)
	ListProducts(ctx context.Context, limit, offset int) ([]Product, error) // no se va a usar por lo visto
	ListProductsByName(ctx context.Context, limit, offset int, description string) ([]Product, error)
	ListProductByCategory(ctx context.Context, categoryID int) ([]Product, error)
}
