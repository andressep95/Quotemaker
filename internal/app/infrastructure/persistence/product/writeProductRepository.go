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
INSERT INTO product (name, category_id, length, price, weight, code, is_available)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, name, category_id, length, price, weight, code, is_available;
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

	row := tx.QueryRowContext(ctx, saveProductQuery, args.Name, args.CategoryID, args.Length, args.Price, args.Weight, args.Code, args.IsAvailable)
	var product domain.Product

	err = row.Scan(
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
		return domain.Product{}, fmt.Errorf("failed to scan product: %w", err)
	}

	return product, nil
}

const updateProductQuery = `
UPDATE product
SET name = $1, category_id = $2, length = $3, price = $4, weight = $5, code = $6, is_available = $7
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

	_, err = tx.ExecContext(ctx, updateProductQuery, args.Name, args.CategoryID, args.Length, args.Price, args.Weight, args.Code, args.IsAvailable, args.ID)
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

func (r *writeProductRepository) DeleteProduct(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, deleteProductQuery, id)
	return err
}

func NewWriteProductRepository(db *sql.DB) domain.WriteProductRepository {
	return &writeProductRepository{
		db: db,
	}
}
