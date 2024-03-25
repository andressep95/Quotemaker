package domain

import "context"

type WriteQuotationRepository interface {
	SaveQuotation(ctx context.Context, args Quotation) (Quotation, error)
	DeleteQuotation(ctx context.Context, id int) error
	UpdateQuotation(ctx context.Context, args Quotation) (Quotation, error)
}
