package persistence

import (
	"context"
	"database/sql"
	"fmt"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
)

type readQuotationRepository struct {
	db *sql.DB
}

const getQuotationByIDQuery = `
	SELECT id, created_at, updated_at, total_price, is_purchased, purchased_at, is_delivered, delivered_at
	FROM quotation
	WHERE id = $1;
`

const getProductsForQuotationQuery = `
	SELECT 
	p.id, 
	p.name, 
	p.category_id, 
	p.length, 
	p.price, 
	p.weight, 
	p.code, 
	p.is_available,
	qp.quantity
	FROM product p
	JOIN quote_product qp ON p.id = qp.product_id
	WHERE qp.quotation_id = $1;
`

func (r *readQuotationRepository) GetQuotationByID(ctx context.Context, id string) (*domain.Quotation, error) {
	row := r.db.QueryRowContext(ctx, getQuotationByIDQuery, id)
	var quotation domain.Quotation

	err := row.Scan(
		&quotation.ID,
		&quotation.CreatedAt,
		&quotation.UpdatedAt,
		&quotation.TotalPrice,
		&quotation.IsPurchased,
		&quotation.PurchasedAt,
		&quotation.IsDelivered,
		&quotation.DeliveredAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("quotation with ID %s does not exist", id)
		}
		return nil, fmt.Errorf("error querying quotation: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, getProductsForQuotationQuery, quotation.ID)
	if err != nil {
		return nil, fmt.Errorf("error querying quote products: %w", err)
	}
	defer rows.Close()

	var products []domain.QuoteProduct

	for rows.Next() {
		var qp domain.QuoteProduct
		if err := rows.Scan(&qp.ID, &qp.QuotationID, &qp.ProductID, &qp.Quantity); err != nil {
			return nil, fmt.Errorf("error scanning quote product: %w", err)
		}
		products = append(products, qp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating quote products: %w", err)
	}

	quotation.Products = products

	return &quotation, nil
}

const listQuotationsQuery = `
SELECT id, created_at, updated_at, total_price, is_purchased, purchased_at, is_delivered, delivered_at
FROM quotation
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
`

func (r *readQuotationRepository) ListQuotations(ctx context.Context, limit int, offset int) ([]domain.Quotation, error) {
	rows, err := r.db.QueryContext(ctx, listQuotationsQuery, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying quotations: %w", err)
	}
	defer rows.Close()

	var quotations []domain.Quotation
	for rows.Next() {
		var q domain.Quotation
		err := rows.Scan(
			&q.ID,
			&q.CreatedAt,
			&q.UpdatedAt,
			&q.TotalPrice,
			&q.IsPurchased,
			&q.PurchasedAt,
			&q.IsDelivered,
			&q.DeliveredAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning quotation: %w", err)
		}
		quotations = append(quotations, q)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating quotations: %w", err)
	}

	return quotations, nil
}

func NewReadQuotationRepository(db *sql.DB) domain.ReadQuotationRepository {
	return &readQuotationRepository{
		db: db,
	}
}
