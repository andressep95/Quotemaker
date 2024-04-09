package domain

import (
	"context"

	dto "github.com/Andressep/QuoteMaker/internal/app/dto/product"
)

type ReadProductRepository interface {
	GetProductByID(ctx context.Context, id string) (*dto.ProductDTO, error)
	ListProducts(ctx context.Context, limit, offset int) ([]dto.ProductDTO, error) // no se va a usar por lo visto
	ListProductsByName(ctx context.Context, limit, offset int, description string) ([]dto.ProductDTO, error)
	ListProductByCategory(ctx context.Context, categoryName string) ([]dto.ProductDTO, error)
}
