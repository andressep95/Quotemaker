package domain

import "context"

type ReadQuotationService struct {
	readQuotationRepo ReadQuotationRepository
}

func NewReadQuotationService(readQuotationRepo ReadQuotationRepository) *ReadQuotationService {
	return &ReadQuotationService{
		readQuotationRepo: readQuotationRepo,
	}
}

func (r *ReadQuotationService) ListQuotations(ctx context.Context, limit, offset int) ([]Quotation, error) {
	return r.readQuotationRepo.ListQuotations(ctx, limit, offset)
}
