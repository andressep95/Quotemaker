package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
)

type writeQuotationRepository struct {
	db *sql.DB
}

func insertQuoteProductsInBatch(ctx context.Context, tx *sql.Tx, quotationID string, products []domain.QuoteProduct) error {
	if len(products) == 0 {
		return nil // No products to insert
	}

	valueStrings := make([]string, 0, len(products))
	valueArgs := make([]interface{}, 0, len(products)*3) // 3 fields per product
	for i, qp := range products {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
		valueArgs = append(valueArgs, quotationID, qp.ProductID, qp.Quantity)
	}

	stmt := fmt.Sprintf("INSERT INTO quote_product (quotation_id, product_id, quantity) VALUES %s", strings.Join(valueStrings, ","))
	_, err := tx.ExecContext(ctx, stmt, valueArgs...)
	return err
}

const insertQuotationQuery = `
		INSERT INTO quotation (created_at, total_price, is_purchased, is_delivered)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, total_price;
	`

func (r *writeQuotationRepository) SaveQuotation(ctx context.Context, args domain.Quotation) (domain.Quotation, error) {
	var err error
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Quotation{}, err
	}

	defer func() {
		// Si err no es nil al final, hacemos rollback
		if p := recover(); p != nil || err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Se crean valores predeterminados para los campos que pueden ser nil
	now := time.Now()
	quotation := domain.Quotation{
		CreatedAt:   now,
		TotalPrice:  args.TotalPrice,
		IsPurchased: args.IsPurchased,
		IsDelivered: args.IsDelivered,
		Products:    args.Products,
	}

	err = tx.QueryRowContext(ctx, insertQuotationQuery, now, quotation.TotalPrice, false, false).Scan(
		&quotation.ID,
		&quotation.CreatedAt,
		&quotation.TotalPrice,
	)

	if err != nil {
		return domain.Quotation{}, err
	}

	if err = insertQuoteProductsInBatch(ctx, tx, quotation.ID, quotation.Products); err != nil {
		return domain.Quotation{}, err
	}

	return quotation, nil
}

const updateQuotationQuery = `
    UPDATE quotation
    SET updated_at = $1, total_price = $2, is_purchased = $3, purchased_at = $4, is_delivered = $5, delivered_at = $6
    WHERE id = $7;
`

const upsertQuoteProductQuery = `
    INSERT INTO quote_product (quotation_id, product_id, quantity)
    SELECT $1, $2, $3
    WHERE NOT EXISTS (
        SELECT 1 FROM quote_product
        WHERE quotation_id = $1 AND product_id = $2
    )
    ON CONFLICT (quotation_id, product_id)
    DO UPDATE SET quantity = EXCLUDED.quantity;
`

func (r *writeQuotationRepository) UpdateQuotation(ctx context.Context, args domain.Quotation) (domain.Quotation, error) {
	var err error
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Quotation{}, err
	}

	// Uso de defer para manejar el rollback de manera centralizada
	defer func() {
		if p := recover(); p != nil || err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	now := time.Now()
	args.UpdatedAt = &now // Actualizamos UpdatedAt directamente en args

	_, err = tx.ExecContext(ctx, updateQuotationQuery,
		args.UpdatedAt,
		args.TotalPrice,
		args.IsPurchased,
		args.PurchasedAt,
		args.IsDelivered,
		args.DeliveredAt,
		args.ID)
	if err != nil {
		return domain.Quotation{}, err
	}

	// Actualizamos los productos asociados
	for i, product := range args.Products {
		_, err = tx.ExecContext(ctx, upsertQuoteProductQuery, args.ID, product.ProductID, product.Quantity)
		if err != nil {
			return domain.Quotation{}, err
		}
		args.Products[i] = product
	}

	return args, nil
}

const deleteQuotationQuery = `
	DELETE FROM quotation
	WHERE id = $1;
`

func (r *writeQuotationRepository) DeleteQuotation(ctx context.Context, id string) error {
	result, err := r.db.ExecContext(ctx, deleteQuotationQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting quotation: %w", err)
	}

	// Verifica si algún registro fue afectado
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("quotation with ID %s does not exist", id)
	}

	return nil
}
func NewWriteQuotationRepository(db *sql.DB) domain.WriteQuotationRepository {
	return &writeQuotationRepository{
		db: db,
	}
}
