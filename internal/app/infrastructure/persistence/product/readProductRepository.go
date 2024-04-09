package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/product"
)

type readProductRepository struct {
	db *sql.DB
}

const listProductsByNameQuery = `
	SELECT p.id, p.code, p.description, ROUND(p.price::numeric, 2), ROUND(p.weight::numeric, 2), ROUND(p.length::numeric, 2), p.is_available, c.category_name
	FROM product p
	INNER JOIN category c ON p.category_id = c.id
	WHERE lower(p.description) LIKE lower($1)
	ORDER BY p.id
	LIMIT $2 OFFSET $3;
	`

// ListProductsByDescription implements domain.ProductRepository.
func (r *readProductRepository) ListProductsByName(ctx context.Context, limit int, offset int, description string) ([]dto.ProductDTO, error) {
	var products []dto.ProductDTO
	rows, err := r.db.QueryContext(ctx, listProductsByNameQuery, "%"+description+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product dto.ProductDTO
		err := rows.Scan(
			&product.ID,
			&product.Code,
			&product.Description,
			&product.Price,
			&product.Weight,
			&product.Length,
			&product.IsAvailable,
			&product.CategoryName,
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
	SELECT p.id, p.code, p.description, ROUND(p.price::numeric, 2), ROUND(p.weight::numeric, 2), ROUND(p.length::numeric, 2), p.is_available, c.category_name
	FROM product p
	INNER JOIN category c ON p.category_id = c.id
	ORDER BY p.id
	LIMIT $1 OFFSET $2;
	`

// ListProductByCategory implements domain.ProductRepository.
func (r *readProductRepository) ListProductByCategory(ctx context.Context, categoryName string) ([]dto.ProductDTO, error) {
	var products []dto.ProductDTO
	rows, err := r.db.QueryContext(ctx, listProductsByCategoryQuery, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product dto.ProductDTO
		err := rows.Scan(
			&product.ID,
			&product.Code,
			&product.Description,
			&product.Price,
			&product.Weight,
			&product.Length,
			&product.IsAvailable,
			&product.CategoryName,
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
	SELECT p.id, p.code, p.description, ROUND(p.price::numeric, 2), ROUND(p.weight::numeric, 2), ROUND(p.length::numeric, 2), p.is_available, c.category_name
	FROM product p
	INNER JOIN category c ON p.category_id = c.id
	ORDER BY p.id
	LIMIT $1 OFFSET $2;
	`

func (r *readProductRepository) ListProducts(ctx context.Context, limit, offset int) ([]dto.ProductDTO, error) {
	rows, err := r.db.QueryContext(ctx, listProductsQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []dto.ProductDTO
	for rows.Next() {
		var product dto.ProductDTO
		err := rows.Scan(
			&product.ID,
			&product.Code,
			&product.Description,
			&product.Price,
			&product.Weight,
			&product.Length,
			&product.IsAvailable,
			&product.CategoryName,
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
	SELECT p.id, p.code, p.description, ROUND(p.price::numeric, 2), ROUND(p.weight::numeric, 2), ROUND(p.length::numeric, 2), p.is_available, c.category_name
	FROM product p
	INNER JOIN category c ON p.category_id = c.id
	WHERE p.id = $1;
	`

func (r *readProductRepository) GetProductByID(ctx context.Context, id string) (*dto.ProductDTO, error) {
	rows := r.db.QueryRowContext(ctx, getProductByIDQuery, id)
	var product dto.ProductDTO

	err := rows.Scan(
		&product.ID,
		&product.Code,
		&product.Description,
		&product.Price,
		&product.Weight,
		&product.Length,
		&product.IsAvailable,
		&product.CategoryName,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("producto con ID %s no encontrado", id)
		}
		return nil, fmt.Errorf("error al recuperar el producto con ID %s: %w", id, err)
	}
	return &product, err
}

func NewReadProductRepository(db *sql.DB) domain.ReadProductRepository {
	return &readProductRepository{
		db: db,
	}
}
