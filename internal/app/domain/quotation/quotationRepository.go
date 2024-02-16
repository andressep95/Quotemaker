package domain

import (
	"context"
)

type QuotationRepository interface {
	SaveQuotation(ctx context.Context, args Quotation) (Quotation, error)
	GetQuotationByID(ctx context.Context, id int) (*Quotation, error)
	ListQuotations(ctx context.Context, limit, offset int) ([]Quotation, error)
	DeleteQuotation(ctx context.Context, id int) error
	UpdateQuotation(ctx context.Context, args Quotation) error
}
