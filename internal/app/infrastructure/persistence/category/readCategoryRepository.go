package persistence

import (
	"context"
	"database/sql"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
)

type ReadCategoryRepository struct {
	db *sql.DB
}

const getCategoryByIDQuery = `
SELECT id, category_name
FROM category
WHERE id = $1;
`

// GetCategoryByID implements ports.CategoryRepository.
func (r *ReadCategoryRepository) GetCategoryByID(ctx context.Context, id string) (*domain.Category, error) {
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
func (r *ReadCategoryRepository) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {
	row := r.db.QueryRowContext(ctx, getCategoryByNameQuery, name)
	var i domain.Category

	err := row.Scan(
		&i.ID,
		&i.CategoryName,
	)

	return i, err
}

const listCategoryQuery = `
SELECT id, category_name
FROM category
ORDER BY id
LIMIT $1 OFFSET $2;
`

// ListCategorys implements ports.CategoryRepository.
func (r *ReadCategoryRepository) ListCategorys(ctx context.Context, limit int, offset int) ([]domain.Category, error) {
	rows, err := r.db.QueryContext(ctx, listCategoryQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorys []domain.Category
	for rows.Next() {
		var i domain.Category
		if err := rows.Scan(
			&i.ID,
			&i.CategoryName); err != nil {
			return nil, err
		}
		categorys = append(categorys, i)
	}

	// Verificar por errores al finalizar la iteración
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categorys, nil
}

func NewReadCategoryRepository(db *sql.DB) domain.ReadCategoryRepository {
	return &ReadCategoryRepository{
		db: db,
	}
}
