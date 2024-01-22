package repository

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type sqlQuotationRepository struct {
	db *sqlx.DB
}

const insertQuotationQuery = `
INSERT INTO quotation (seller_id, customer_id, total_price, is_purchased, is_delivered)
VALUES ($1, $2, $3, $4, $5)
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
		SellerID:    args.SellerID,
		CustomerID:  args.CustomerID,
		TotalPrice:  args.TotalPrice,
		IsPurchased: args.IsPurchased,
		IsDelivered: args.IsDelivered,
		Products:    args.Products,
	}

	// Inserta la cotización y recupera datos relevantes
	row := tx.QueryRowContext(ctx, insertQuotationQuery, quotation.SellerID, quotation.CustomerID, quotation.TotalPrice, quotation.IsPurchased, quotation.IsDelivered)
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

// DeleteQuotation implements ports.QuotationRepository.
func (*sqlQuotationRepository) DeleteQuotation(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetQuotationByID implements ports.QuotationRepository.
func (*sqlQuotationRepository) GetQuotationByID(ctx context.Context, id int) (*domain.Quotation, error) {
	panic("unimplemented")
}

// ListQuotations implements ports.QuotationRepository.
func (*sqlQuotationRepository) ListQuotations(ctx context.Context, limit int, offset int) ([]domain.Quotation, error) {
	panic("unimplemented")
}

func NewQuotationRepository(db *sqlx.DB) ports.QuotationRepository {
	return &sqlQuotationRepository{
		db: db,
	}
}
