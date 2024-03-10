package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
)

type sqlProductRepository struct {
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
func (r *sqlProductRepository) ListProductsByName(ctx context.Context, limit int, offset int, name string) ([]domain.Product, error) {
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
func (r *sqlProductRepository) ListProductByCategory(ctx context.Context, categoryID int) ([]domain.Product, error) {
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

func (r *sqlProductRepository) ListProducts(ctx context.Context, limit, offset int) ([]domain.Product, error) {
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

const saveProductQuery = `
INSERT INTO product (name, category_id, length, price, weight, code, is_available)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, name, category_id, length, price, weight, code, is_available;
`

func (r *sqlProductRepository) SaveProduct(ctx context.Context, args domain.Product) (domain.Product, error) {
	// Iniciar la transacción
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

	// Ejecutar la consulta dentro de la transacción
	row := tx.QueryRowContext(ctx, saveProductQuery, args.Name, args.CategoryID, args.Length, args.Price, args.Weight, args.Code, args.IsAvailable)
	var product domain.Product

	// Escanear el resultado de la consulta en una estructura de producto
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

// UpdateProduct implements ports.ProductRepository.
func (r *sqlProductRepository) UpdateProduct(ctx context.Context, args domain.Product) (domain.Product, error) {
	// Iniciar la transacción
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
	// Manejar posibles errores.
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("producto con ID %d no encontrado", id)
		}
		return nil, fmt.Errorf("error al recuperar el producto con ID %d: %w", id, err)
	}
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

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &sqlProductRepository{
		db: db,
	}
}
