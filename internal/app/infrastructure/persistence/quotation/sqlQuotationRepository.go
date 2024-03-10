package quotation

import (
	"context"
	"database/sql"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
)

type sqlQuotationRepository struct {
	db *sql.DB
}

const insertQuotationQuery = `
INSERT INTO quotation (total_price, is_purchased, is_delivered)
VALUES ($1, $2, $3)
RETURNING id, created_at;
`
const insertProductQuery = `
INSERT INTO quote_product (quotation_id, product_id, quantity)
VALUES ($1, $2, $3);
`

// SaveQuptation implements ports.QuotationRepository.
func (r *sqlQuotationRepository) SaveQuotation(ctx context.Context, args domain.Quotation) (domain.Quotation, error) {
	// Inicia una transacción
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.Quotation{}, err
	}

	// Manejo de Rollback en caso de error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Completa los campos de la cotización con los datos de args
	quotation := domain.Quotation{
		TotalPrice:  args.TotalPrice,
		IsPurchased: args.IsPurchased,
		IsDelivered: args.IsDelivered,
		Products:    args.Products,
	}

	// Inserta la cotización y recupera datos relevantes
	row := tx.QueryRowContext(ctx, insertQuotationQuery, quotation.TotalPrice, quotation.IsPurchased, quotation.IsDelivered)
	err = row.Scan(
		&quotation.ID,
		&quotation.CreatedAt,
	)
	if err != nil {
		return domain.Quotation{}, err // El Rollback se maneja mediante defer
	}

	// Inserta cada producto asociado en la tabla 'quote_product'
	for _, quotationProducts := range quotation.Products {
		if _, err = tx.ExecContext(ctx, insertProductQuery, quotation.ID, quotationProducts.ProductID, quotationProducts.Quantity); err != nil {
			return domain.Quotation{}, err // El Rollback se maneja mediante defer
		}
	}

	// Compromete la transacción
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

// UpdateQuotation implements ports.QuotationRepository.
func (r *sqlQuotationRepository) UpdateQuotation(ctx context.Context, args domain.Quotation) error {
	now := time.Now()
	if args.UpdatedAt == nil {
		args.UpdatedAt = &now
	}
	_, err := r.db.ExecContext(ctx, updateQuotationQuery,
		args.UpdatedAt,
		args.TotalPrice,
		args.IsPurchased,
		args.PurchasedAt,
		args.IsDelivered,
		args.DeliveredAt,
		args.ID)
	if err != nil {
		return err
	}
	return nil
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

// GetQuotationByID implements ports.QuotationRepository.
func (r *sqlQuotationRepository) GetQuotationByID(ctx context.Context, id int) (*domain.Quotation, error) {
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
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, getProductsForQuotationQuery, quotation.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.QuoteProduct

	for rows.Next() {
		var qp domain.QuoteProduct
		if err := rows.Scan(&qp.ID, &qp.QuotationID, &qp.ProductID, &qp.Quantity); err != nil {
			return nil, err
		}
		products = append(products, qp)
	}

	// Verifica si hubo errores durante la iteración
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Asigna los productos a la cotización
	quotation.Products = products

	return &quotation, nil
}

const deleteQuotationQuery = `
DELETE FROM quotation
WHERE id = $1;
`

// DeleteQuotation implements ports.QuotationRepository.
func (r *sqlQuotationRepository) DeleteQuotation(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, deleteQuotationQuery, id)
	return err
}

const listQuotationsQuery = `
SELECT id, created_at, updated_at, total_price, is_purchased, purchased_at, is_delivered, delivered_at
FROM quotation
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
`

// ListQuotations implements ports.QuotationRepository.
func (r *sqlQuotationRepository) ListQuotations(ctx context.Context, limit int, offset int) ([]domain.Quotation, error) {
	rows, err := r.db.QueryContext(ctx, listQuotationsQuery, limit, offset)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		quotations = append(quotations, q)
	}

	// Check for any iteration errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quotations, nil
}

func NewQuotationRepository(db *sql.DB) domain.QuotationRepository {
	return &sqlQuotationRepository{
		db: db,
	}
}
