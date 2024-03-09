package domain

import (
	"context"
)

type ProductRepository interface {
	SaveProduct(ctx context.Context, args Product) (Product, error)
	GetProductByID(ctx context.Context, id int) (*Product, error)
	ListProducts(ctx context.Context, limit, offset int) ([]Product, error) // no se va a usar por lo visto
	DeleteProduct(ctx context.Context, id int) error
	UpdateProduct(ctx context.Context, args Product) error
	ListProductsByName(ctx context.Context, limit, offset int, description string) ([]Product, error)
	ListProductByCategory(ctx context.Context, categoryID int) ([]Product, error)
}
