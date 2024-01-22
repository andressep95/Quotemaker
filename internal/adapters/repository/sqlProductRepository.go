package repository

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type sqlProductRepository struct {
	db *sqlx.DB
}

const listProductsQuery = `
SELECT id, name, category_id, length, price, weight, code, is_available
FROM product
ORDER BY id
LIMIT $1 OFFSET $2;
`

func (r *sqlProductRepository) ListProducts(ctx context.Context, limit, offset int) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.SelectContext(ctx, &products, listProductsQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

const saveProductQuery = `
INSERT INTO product (name, category_id, length, price, weight, code, is_available)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, name, category_id, length, price, weight, code, is_available;
`

func (r *sqlProductRepository) SaveProduct(ctx context.Context, args domain.Product) (domain.Product, error) {
	row := r.db.QueryRowContext(ctx, saveProductQuery, args.Name, args.CategoryID, args.Length, args.Price, args.Weight, args.Code, args.IsAvailable)
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
	return i, err
}

const getProductByIDQuery = `
SELECT id, name, category_id, length, price, weight, code, is_available
FROM product
WHERE id = $1;
`

func (r *sqlProductRepository) GetProductByID(ctx context.Context, id int) (*domain.Product, error) {
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

	return &i, err
}

const deleteProductQuery = `
DELETE FROM product
WHERE id = $1;
`

func (r *sqlProductRepository) DeleteProduct(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, deleteProductQuery, id)
	return err
}

func NewProductRepository(db *sqlx.DB) ports.ProductRepository {
	return &sqlProductRepository{
		db: db,
	}
}
