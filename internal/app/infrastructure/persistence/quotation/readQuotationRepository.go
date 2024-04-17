package persistence

import (
	"context"
	"database/sql"
	"fmt"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
	dtoP "github.com/Andressep/QuoteMaker/internal/app/dto/product"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/quotation"
)

type readQuotationRepository struct {
	db *sql.DB
}

// GetQuotationByID implements domain.ReadQuotationRepository.
func (r *readQuotationRepository) GetQuotationByID(ctx context.Context, id string) (*domain.Quotation, error) {
	panic("unimplemented")
}

const listQuotationsQuery = `
    SELECT 
    q.id, q.created_at, q.updated_at, q.total_price, q.is_purchased, q.purchased_at, q.is_delivered, q.delivered_at,
    p.id AS product_id, p.code, p.description, cat.category_name, p.price, p.weight, p.length, p.is_available, qp.quantity
    FROM quotation q
    LEFT JOIN quote_product qp ON q.id = qp.quotation_id
    LEFT JOIN product p ON p.id = qp.product_id
    LEFT JOIN category cat ON cat.id = p.category_id
    ORDER BY q.created_at DESC, q.id, p.id
    LIMIT $1 OFFSET $2;
`

func (r *readQuotationRepository) ListQuotations(ctx context.Context, limit int, offset int) ([]dto.QuotationDTO, error) {
	rows, err := r.db.QueryContext(ctx, listQuotationsQuery, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying quotations: %w", err)
	}
	defer rows.Close()

	var quotations []dto.QuotationDTO
	var lastID string
	var currentQuotation dto.QuotationDTO

	for rows.Next() {
		var q dto.QuotationDTO
		var prod dtoP.ProductDTO
		var quantity int

		err := rows.Scan(
			&q.ID,
			&q.CreatedAt,
			&q.PurchasedAt,
			&q.TotalPrice,
			&q.IsPurchased,
			&q.PurchasedAt,
			&q.IsDelivered,
			&q.DeliveredAt,
			&prod.ID,
			&prod.Code,
			&prod.Description,
			&prod.CategoryName,
			&prod.Price,
			&prod.Weight,
			&prod.Length,
			&prod.IsAvailable,
			&quantity,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning quotation: %w", err)
		}

		if q.ID != lastID {
			if lastID != "" {
				quotations = append(quotations, currentQuotation)
			}
			currentQuotation = q
			currentQuotation.Products = []dto.QuoteProductDetail{}
			lastID = q.ID
		}

		if prod.ID != "" {
			currentQuotation.Products = append(currentQuotation.Products, dto.QuoteProductDetail{
				Product:  prod,
				Quantity: quantity,
			})
		}
	}
	if lastID != "" { // add the last quotation if any
		quotations = append(quotations, currentQuotation)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating quotations: %w", err)
	}

	return quotations, nil
}

func NewReadQuotationRepository(db *sql.DB) domain.ReadQuotationRepository {
	return &readQuotationRepository{
		db: db,
	}
}
