package persistence

import (
	"context"
	"database/sql"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
)

type WriteCategoryRepository struct {
	db *sql.DB
}

const saveCategoryQuery = `
INSERT INTO category (category_name)
VALUES ($1)
RETURNING id, category_name;
`

// SaveCategory implements ports.CategoryRepository.
func (w *WriteCategoryRepository) SaveCategory(ctx context.Context, args domain.Category) (domain.Category, error) {
	row := w.db.QueryRowContext(ctx, saveCategoryQuery, args.CategoryName)
	var i domain.Category

	err := row.Scan(
		&i.ID,
		&i.CategoryName,
	)
	return i, err
}

const updateCategoryQuery = `
UPDATE category
SET category_name = $1
WHERE id = $2;
`

// UpdateCategory implements ports.CategoryRepository.
func (w *WriteCategoryRepository) UpdateCategory(ctx context.Context, category domain.Category) error {
	_, err := w.db.ExecContext(ctx, updateCategoryQuery, category.CategoryName, category.ID)
	if err != nil {
		return err
	}
	return nil
}

const deleteCategoryQuery = `
DELETE FROM category
WHERE id = $1;
`

// DeleteCategory implements ports.CategoryRepository.
func (w *WriteCategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	_, err := w.db.ExecContext(ctx, deleteCategoryQuery, id)
	return err
}

func NewWriteCategoryRepository(db *sql.DB) domain.WriteCategoryRepository {
	return &WriteCategoryRepository{
		db: db,
	}
}
