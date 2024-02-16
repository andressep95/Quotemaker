package domain

import (
	"context"
)

type CategoryRepository interface {
	SaveCategory(ctx context.Context, args Category) (Category, error)
	UpdateCategory(ctx context.Context, category Category) error
	GetCategoryByID(ctx context.Context, id int) (*Category, error)
	ListCategorys(ctx context.Context, limit, offset int) ([]Category, error)
	DeleteCategory(ctx context.Context, id int) error
	GetCategoryByName(ctx context.Context, name string) (Category, error)
}
