package domain

import (
	"context"

	dto "github.com/Andressep/QuoteMaker/internal/app/dto/quotation"
)

type ReadQuotationRepository interface {
	GetQuotationByID(ctx context.Context, id string) (*Quotation, error)
	ListQuotations(ctx context.Context, limit, offset int) ([]dto.QuotationDTO, error)
}
