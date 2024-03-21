package domain

import (
	"context"
)

type WriteCategoryRepository interface {
	SaveCategory(ctx context.Context, args Category) (Category, error)
	UpdateCategory(ctx context.Context, category Category) error
	DeleteCategory(ctx context.Context, id int) error
}
