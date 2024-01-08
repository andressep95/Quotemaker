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

const saveProductQuery = `
INSERT INTO product (name, category, length, price, weight, code)
VALUES ($1, $2, $3, $4, $5, $6);
`

// saveProduct implements ports.ProductRepo.
func (r *sqlProductRepository) SaveProduct(ctx context.Context, args domain.Product) (domain.Product, error) {
	row := r.db.QueryRowContext(ctx, saveProductQuery, args.Name, args.CategoryID, args.Length, args.Price, args.Weight, args.Code)
	var i domain.Product

	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CategoryID,
		&i.Length,
		&i.Price,
		&i.Weight,
		&i.Code,
	)
	return i, err
}

const getProductByIDQuery = `
SELECT id, name, category, length, price, weight, code
FROM product
WHERE id = ?;
`

// getProductByID implements ports.ProductRepository.
func (r *sqlProductRepository) GetProductByID(ctx context.Context, id int) (*domain.Product, error) {
	product := &domain.Product{}
	err := r.db.GetContext(ctx, product, getProductByIDQuery, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func New(db *sqlx.DB) ports.ProductRepository {
	return &sqlProductRepository{
		db: db,
	}
}
