package repository

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type sqlCategoryRepository struct {
	db *sqlx.DB
}

const saveCategoryQuery = `
INSERT INTO category (category_name)
VALUES ($1)
RETURNING id, category_name;
`

// SaveCategory implements ports.CategoryRepository.
func (r *sqlCategoryRepository) SaveCategory(ctx context.Context, args domain.Category) (domain.Category, error) {
	row := r.db.QueryRowContext(ctx, saveCategoryQuery, args.CategoryName)
	var i domain.Category

	err := row.Scan(
		&i.ID,
		&i.CategoryName,
	)
	return i, err

}

// GetCategoryByID implements ports.CategoryRepository.
func (*sqlCategoryRepository) GetCategoryByID(ctx context.Context, id int) (domain.Category, error) {
	panic("unimplemented")
}

// ListCategorys implements ports.CategoryRepository.
func (*sqlCategoryRepository) ListCategorys(ctx context.Context, limit int, offset int) ([]domain.Category, error) {
	panic("unimplemented")
}

// DeleteCategory implements ports.CategoryRepository.
func (*sqlCategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	panic("unimplemented")
}

func NewCategoryRepository(db *sqlx.DB) ports.CategoryRepository {
	return &sqlCategoryRepository{
		db: db,
	}
}
