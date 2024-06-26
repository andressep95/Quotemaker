package domain

import (
	"context"
)

type ReadCategoryRepository interface {
	GetCategoryByID(ctx context.Context, id string) (*Category, error)
	ListCategorys(ctx context.Context, limit, offset int) ([]Category, error)
	GetCategoryByName(ctx context.Context, name string) (Category, error)
}
