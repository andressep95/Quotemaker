package ports

import (
	"context"

	"github.com/Andressep/QuoteMaker/internal/core/domain"
)

type ProductRepository interface {
	SaveProduct(ctx context.Context, args domain.Product) (domain.Product, error)
	GetProductByID(ctx context.Context, id int) (*domain.Product, error)
}
