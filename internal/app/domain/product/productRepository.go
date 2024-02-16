package domain

import (
	"context"
)

type ProductRepository interface {
	SaveProduct(ctx context.Context, args Product) (Product, error)
	GetProductByID(ctx context.Context, id int) (*Product, error)
	ListProducts(ctx context.Context, limit, offset int) ([]Product, error)
	DeleteProduct(ctx context.Context, id int) error
	UpdateProduct(ctx context.Context, args Product) error
}
