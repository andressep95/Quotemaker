package domain

import "context"

type ReadQuotationRepository interface {
	GetQuotationByID(ctx context.Context, id int) (*Quotation, error)
	ListQuotations(ctx context.Context, limit, offset int) ([]Quotation, error)
}
