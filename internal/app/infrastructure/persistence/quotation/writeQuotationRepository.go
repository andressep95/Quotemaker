package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
)

type writeQuotationRepository struct {
	db *sql.DB
}

const insertQuotationQuery = `
		INSERT INTO quotation (created_at, total_price)
		VALUES ($1, $2)
		RETURNING id, created_at, total_price;
	`
const insertProductQuery = `
		INSERT INTO quote_product (quotation_id, product_id, quantity)
		VALUES ($1, $2, $3);
	`

func (r *writeQuotationRepository) SaveQuotation(ctx context.Context, args domain.Quotation) (domain.Quotation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Quotation{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	quotation := domain.Quotation{
		CreatedAt:  time.Now(),
		TotalPrice: args.TotalPrice,
		Products:   args.Products,
	}

	row := tx.QueryRowContext(ctx, insertQuotationQuery, quotation.TotalPrice)
	err = row.Scan(
		&quotation.ID,
		&quotation.CreatedAt,
		&quotation.TotalPrice,
	)
	if err != nil {
		return domain.Quotation{}, err
	}

	for _, quotationProducts := range quotation.Products {
		if _, err = tx.ExecContext(ctx, insertProductQuery, quotation.ID, quotationProducts.ProductID, quotationProducts.Quantity); err != nil {
			return domain.Quotation{}, err
		}
	}

	if err = tx.Commit(); err != nil {
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
	VALUES ($1, $2, $3)
	ON CONFLICT (quotation_id, product_id)
	DO UPDATE SET quantity = EXCLUDED.quantity;
`

const selectQuotationQuery = `
	SELECT * FROM quotation WHERE id = $1;
`

const selectQuoteProductsQuery = `
	SELECT * FROM quote_product WHERE quotation_id = $1;
`

func (r *writeQuotationRepository) UpdateQuotation(ctx context.Context, args domain.Quotation) (domain.Quotation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Quotation{}, err
	}

	now := time.Now()
	if args.UpdatedAt == nil {
		args.UpdatedAt = &now
	}

	_, err = tx.ExecContext(ctx, updateQuotationQuery,
		args.UpdatedAt,
		args.TotalPrice,
		args.IsPurchased,
		args.PurchasedAt,
		args.IsDelivered,
		args.DeliveredAt,
		args.ID)
	if err != nil {
		tx.Rollback()
		return domain.Quotation{}, err
	}

	for _, product := range args.Products {
		_, err = tx.ExecContext(ctx, upsertQuoteProductQuery, args.ID, product.ProductID, product.Quantity)
		if err != nil {
			tx.Rollback()
			return domain.Quotation{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return domain.Quotation{}, err
	}

	row := r.db.QueryRowContext(ctx, selectQuotationQuery, args.ID)
	quotation := domain.Quotation{}
	err = row.Scan(&quotation.ID, &quotation.CreatedAt, &quotation.UpdatedAt, &quotation.TotalPrice, &quotation.IsPurchased, &quotation.PurchasedAt, &quotation.IsDelivered, &quotation.DeliveredAt)
	if err != nil {
		return domain.Quotation{}, err
	}

	rows, err := r.db.QueryContext(ctx, selectQuoteProductsQuery, args.ID)
	if err != nil {
		return domain.Quotation{}, fmt.Errorf("error querying quote products: %w", err)
	}
	defer rows.Close()

	quotation.Products = []domain.QuoteProduct{}
	for rows.Next() {
		product := domain.QuoteProduct{}
		err = rows.Scan(&product.ID, &product.QuotationID, &product.ProductID, &product.Quantity)
		if err != nil {
			return domain.Quotation{}, fmt.Errorf("error scanning quote product: %w", err)
		}
		quotation.Products = append(quotation.Products, product)
	}

	if err = rows.Err(); err != nil {
		return domain.Quotation{}, fmt.Errorf("error after iterating quote products: %w", err)
	}

	return quotation, nil
}

const deleteQuotationQuery = `
	DELETE FROM quotation
	WHERE id = $1;
`

func (r *writeQuotationRepository) DeleteQuotation(ctx context.Context, id int) error {
	row := r.db.QueryRowContext(ctx, selectQuotationQuery, id)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("quotation with ID %d does not exist", id)
		}
		return fmt.Errorf("error checking existence of quotation: %w", err)
	}

	_, err := r.db.ExecContext(ctx, deleteQuotationQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting quotation: %w", err)
	}

	return nil
}
func NewWriteQuotationRepository(db *sql.DB) domain.WriteQuotationRepository {
	return &writeQuotationRepository{
		db: db,
	}
}
