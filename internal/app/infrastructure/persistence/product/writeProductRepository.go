package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
)

type writeProductRepository struct {
	db *sql.DB
}

const saveProductQuery = `
INSERT INTO product (category_id, code, description, price, weight, length, is_available)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, category_id, code, description, price, weight, length, is_available;
`

func (r *writeProductRepository) SaveProduct(ctx context.Context, args domain.Product) (domain.Product, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	rows := tx.QueryRowContext(ctx, saveProductQuery, args.CategoryID, args.Code, args.Description, args.Price, args.Weight, args.Length, args.IsAvailable)
	var product domain.Product

	err = rows.Scan(
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
		return domain.Product{}, fmt.Errorf("failed to scan product: %w", err)
	}

	return product, nil
}

const updateProductQuery = `
UPDATE product
SET category_id = $1, code = $2, description = $3, price = $4, weight = $5, length = $6, is_available = $7
WHERE id = $8;
`

func (r *writeProductRepository) UpdateProduct(ctx context.Context, args domain.Product) (domain.Product, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Product{}, fmt.Errorf("could not begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Printf("transaction rolled back: %v", err)
			return
		}
		tx.Commit()
	}()

	_, err = tx.ExecContext(ctx, updateProductQuery, args.CategoryID, args.Code, args.Description, args.Price, args.Weight, args.Length, args.IsAvailable, args.ID)
	if err != nil {
		log.Printf("error updating the product: %v", err)
		return domain.Product{}, err
	}
	return args, nil
}

const deleteProductQuery = `
DELETE FROM product
WHERE id = $1;
`

func (r *writeProductRepository) DeleteProduct(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, deleteProductQuery, id)
	return err
}

func NewWriteProductRepository(db *sql.DB) domain.WriteProductRepository {
	return &writeProductRepository{
		db: db,
	}
}
