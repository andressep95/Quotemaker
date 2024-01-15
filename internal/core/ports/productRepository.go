package ports

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
)

type ProductRepository interface {
	SaveProduct(ctx context.Context, args domain.Product) (domain.Product, error)
	GetProductByID(ctx context.Context, id int) (domain.Product, error)
	ListProducts(ctx context.Context, limit, offset int) ([]domain.Product, error)
	DeleteProduct(ctx context.Context, id int) error
}
