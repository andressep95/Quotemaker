package ports

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
)

type ProductRepository interface {
	SaveProduct(ctx context.Context, args domain.Product) (domain.Product, error)
	GetProductByID(ctx context.Context, id int) (*domain.Product, error)
}
