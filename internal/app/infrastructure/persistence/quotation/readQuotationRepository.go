package persistence

import (
	"context"
	"database/sql"
	"fmt"

	domainP "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	domainQ "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
)

type readQuotationRepository struct {
	db *sql.DB
}

const getQuotationByIDQuery = `
SELECT
    q.id, q.created_at, q.updated_at, q.total_price, q.is_purchased, q.purchased_at, q.is_delivered, q.delivered_at,
    p.id AS product_id, p.description, p.category_id, p.length, p.price, p.weight, p.code, p.is_available,
    qp.quantity
	FROM quotation q
	LEFT JOIN quote_product qp ON q.id = qp.quotation_id
	LEFT JOIN product p ON p.id = qp.product_id
	WHERE q.id = $1;
`

func (r *readQuotationRepository) GetQuotationByID(ctx context.Context, id string) (*domainQ.Quotation, error) {
	rows, err := r.db.QueryContext(ctx, getQuotationByIDQuery, id)
	if err != nil {
		return nil, fmt.Errorf("error querying combined quotation data: %w", err)
	}
	defer rows.Close()

	var quotation *domainQ.Quotation
	products := make(map[string]domainP.Product)

	for rows.Next() {
		var q domainQ.Quotation
		var p domainP.Product
		var quantity int

		err := rows.Scan(
			&q.ID, &q.CreatedAt, &q.UpdatedAt, &q.TotalPrice, &q.IsPurchased, &q.PurchasedAt, &q.IsDelivered, &q.DeliveredAt,
			&p.ID, &p.Description, &p.CategoryID, &p.Length, &p.Price, &p.Weight, &p.Code, &p.IsAvailable,
			&quantity,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning combined quotation data: %w", err)
		}

		// Si es la primera vez que vemos esta cotización, inicialízala.
		if quotation == nil {
			quotation = &q
			quotation.Products = []domainQ.QuoteProduct{} // Inicializa el slice de productos.
		}

		// Añade el producto a la cotización si aún no se ha añadido.
		if _, exists := products[p.ID]; !exists && p.ID != "" {
			products[p.ID] = p // Marcamos el producto como visto.
			quotation.Products = append(quotation.Products, domainQ.QuoteProduct{
				ProductID: p.ID,
				Quantity:  quantity,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating combined quotation data: %w", err)
	}

	if quotation == nil {
		return nil, fmt.Errorf("quotation with ID %s does not exist", id)
	}

	return quotation, nil
}

const listQuotationsQuery = `
		SELECT id, created_at, updated_at, total_price, is_purchased, purchased_at, is_delivered, delivered_at
		FROM quotation
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2;
	`

func (r *readQuotationRepository) ListQuotations(ctx context.Context, limit int, offset int) ([]domainQ.Quotation, error) {
	rows, err := r.db.QueryContext(ctx, listQuotationsQuery, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying quotations: %w", err)
	}
	defer rows.Close()

	var quotations []domainQ.Quotation
	for rows.Next() {
		var q domainQ.Quotation
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

func NewReadQuotationRepository(db *sql.DB) domainQ.ReadQuotationRepository {
	return &readQuotationRepository{
		db: db,
	}
}
