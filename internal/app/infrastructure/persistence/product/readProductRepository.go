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
		SELECT id, category_id, code, description, ROUND(price::numeric, 2), ROUND(weight::numeric, 2), ROUND(length::numeric, 2), is_available
		FROM product
        WHERE lower(description) LIKE lower($1)
        ORDER BY id
        LIMIT $2 OFFSET $3;
    `

// ListProductsByDescription implements domain.ProductRepository.
func (r *readProductRepository) ListProductsByName(ctx context.Context, limit int, offset int, description string) ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.QueryContext(ctx, listProductsByNameQuery, "%"+description+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(
			&product.ID,
			&product.CategoryID,
			&product.Code,
			&product.Description,
			&product.Price,
			&product.Weight,
			&product.Length,
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
		SELECT id, category_id, code, description, ROUND(price::numeric, 2), ROUND(weight::numeric, 2), ROUND(length::numeric, 2), is_available
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
			&product.CategoryID,
			&product.Code,
			&product.Description,
			&product.Price,
			&product.Weight,
			&product.Length,
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
		SELECT id, category_id, code, description, ROUND(price::numeric, 2), ROUND(weight::numeric, 2), ROUND(length::numeric, 2), is_available
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
		var product domain.Product
		err := rows.Scan(
			&product.ID,
			&product.CategoryID,
			&product.Code,
			&product.Description,
			&product.Price,
			&product.Weight,
			&product.Length,
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

const getProductByIDQuery = `
		SELECT id, category_id, code, description, ROUND(price::numeric, 2), ROUND(weight::numeric, 2), ROUND(length::numeric, 2), is_available
		FROM product
		WHERE id = $1;
	`

func (r *readProductRepository) GetProductByID(ctx context.Context, id int) (*domain.Product, error) {
	rows := r.db.QueryRowContext(ctx, getProductByIDQuery, id)
	var product domain.Product

	err := rows.Scan(
		&product.ID,
		&product.CategoryID,
		&product.Code,
		&product.Description,
		&product.Price,
		&product.Weight,
		&product.Length,
		&product.IsAvailable,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("producto con ID %d no encontrado", id)
		}
		return nil, fmt.Errorf("error al recuperar el producto con ID %d: %w", id, err)
	}
	return &product, err
}

func NewReadProductRepository(db *sql.DB) domain.ReadProductRepository {
	return &readProductRepository{
		db: db,
	}
}
