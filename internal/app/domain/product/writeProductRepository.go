package domain

import (
	"context"
)

type WriteProductRepository interface {
	SaveProduct(ctx context.Context, args Product) (Product, error)
	DeleteProduct(ctx context.Context, id int) error
	UpdateProduct(ctx context.Context, args Product) (Product, error)
}
