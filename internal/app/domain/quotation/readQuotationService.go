package domain

import (
	"context"

	dto "github.com/Andressep/QuoteMaker/internal/app/dto/quotation"
)

type ReadQuotationService struct {
	readQuotationRepo ReadQuotationRepository
}

func NewReadQuotationService(readQuotationRepo ReadQuotationRepository) *ReadQuotationService {
	return &ReadQuotationService{
		readQuotationRepo: readQuotationRepo,
	}
}

func (r *ReadQuotationService) ListQuotations(ctx context.Context, limit, offset int) ([]dto.QuotationDTO, error) {
	return r.readQuotationRepo.ListQuotations(ctx, limit, offset)
}
