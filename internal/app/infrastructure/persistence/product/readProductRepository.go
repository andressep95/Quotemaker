package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
)

type readProductRepository struct {
	db *sql.DB
}

const listProductsByNameQuery = `
        SELECT id, name, category_id, ROUND(length::numeric, 2), ROUND(price::numeric, 2), ROUND(weight::numeric, 2), code, is_available
        FROM product
        WHERE lower(name) LIKE lower($1)
        ORDER BY id
        LIMIT $2 OFFSET $3;
    `

// ListProductsByDescription implements domain.ProductRepository.
func (r *readProductRepository) ListProductsByName(ctx context.Context, limit int, offset int, name string) ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.QueryContext(ctx, listProductsByNameQuery, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.CategoryID,
			&product.Length,
			&product.Price,
			&product.Weight,
			&product.Code,
			&product.IsAvailable,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

const listProductsByCategoryQuery = `
        SELECT id, name, category_id, ROUND(length::numeric, 2), ROUND(price::numeric, 2), ROUND(weight::numeric, 2), code, is_available
        FROM product
        WHERE category_id = $1
        ORDER BY name;
    `

// ListProductByCategory implements domain.ProductRepository.
func (r *readProductRepository) ListProductByCategory(ctx context.Context, categoryID int) ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.QueryContext(ctx, listProductsByCategoryQuery, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.CategoryID,
			&product.Length,
			&product.Price,
			&product.Weight,
			&product.Code,
			&product.IsAvailable,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

const listProductsQuery = `
SELECT id, name, code, category_id, ROUND(length::numeric, 2) as length, ROUND(price::numeric, 2) as price, ROUND(weight::numeric, 2) as weight, is_available
FROM product
ORDER BY id
LIMIT $1 OFFSET $2;
`

func (r *readProductRepository) ListProducts(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	rows, err := r.db.QueryContext(ctx, listProductsQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var i domain.Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.CategoryID,
			&i.Length,
			&i.Price,
			&i.Weight,
			&i.IsAvailable,
		); err != nil {
			return nil, err
		}
		products = append(products, i)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

const getProductByIDQuery = `
SELECT id, name, category_id, length, price, weight, code, is_available
FROM product
WHERE id = $1;
`

func (r *readProductRepository) GetProductByID(ctx context.Context, id int) (*domain.Product, error) {
	row := r.db.QueryRowContext(ctx, getProductByIDQuery, id)
	var i domain.Product

	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryID,
		&i.Length,
		&i.Price,
		&i.Weight,
		&i.Code,
		&i.IsAvailable,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("producto con ID %d no encontrado", id)
		}
		return nil, fmt.Errorf("error al recuperar el producto con ID %d: %w", id, err)
	}
	return &i, err
}

func NewReadProductRepository(db *sql.DB) domain.ReadProductRepository {
	return &readProductRepository{
		db: db,
	}
}
