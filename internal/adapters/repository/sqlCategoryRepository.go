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

const getCategoryByIDQuery = `
SELECT id, category_name
FROM category
WHERE id = $1;
`

// GetCategoryByID implements ports.CategoryRepository.
func (r *sqlCategoryRepository) GetCategoryByID(ctx context.Context, id int) (*domain.Category, error) {
	row := r.db.QueryRowContext(ctx, getCategoryByIDQuery, id)
	var i domain.Category

	err := row.Scan(
		&i.ID,
		&i.CategoryName,
	)
	return &i, err
}

const getCategoryByNameQuery = `
SELECT id, category_name
FROM category
WHERE category_name = $1;
`

// GetCategoryByName implements ports.CategoryRepository.
func (r *sqlCategoryRepository) GetCategoryByName(ctx context.Context, name string) (*domain.Category, error) {
	row := r.db.QueryRowContext(ctx, getCategoryByNameQuery, name)
	var i domain.Category

	err := row.Scan(
		&i.ID,
		&i.CategoryName,
	)

	return &i, err
}

const listCategoryQuery = `
SELECT category_name
FROM category
ORDER BY id
LIMIT $1 OFFSET $2;
`

// ListCategorys implements ports.CategoryRepository.
func (r *sqlCategoryRepository) ListCategorys(ctx context.Context, limit int, offset int) ([]domain.Category, error) {
	var categorys []domain.Category
	err := r.db.SelectContext(ctx, &categorys, listCategoryQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

const deleteCategoryQuery = `
DELETE FROM category
WHERE id = $1;
`

// DeleteCategory implements ports.CategoryRepository.
func (r *sqlCategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, deleteCategoryQuery, id)
	return err
}

func NewCategoryRepository(db *sqlx.DB) ports.CategoryRepository {
	return &sqlCategoryRepository{
		db: db,
	}
}
