package ports

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
)

type QuotationRepository interface {
	SaveQuotation(ctx context.Context, args domain.Quotation) (domain.Quotation, error)
	GetQuotationByID(ctx context.Context, id int) (*domain.Quotation, error)
	ListQuotations(ctx context.Context, limit, offset int) ([]domain.Quotation, error)
	DeleteQuotation(ctx context.Context, id int) error
}
