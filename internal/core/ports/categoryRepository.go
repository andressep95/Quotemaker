package ports

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
)

type CategoryRepository interface {
	SaveCategory(ctx context.Context, args domain.Category) (domain.Category, error)
	GetCategoryByID(ctx context.Context, id int) (*domain.Category, error)
	ListCategorys(ctx context.Context, limit, offset int) ([]domain.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}
