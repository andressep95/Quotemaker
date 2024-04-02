package domain

import (
	"context"
	"errors"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
)

type WriteQuotationService struct {
	writeQuotationRepo WriteQuotationRepository
	writeProductRepo   domain.WriteProductRepository
}

func NewWriteQuotationService(writeQuotationRepo WriteQuotationRepository, writeProductRepo domain.WriteProductRepository) *WriteQuotationService {
	return &WriteQuotationService{
		writeQuotationRepo: writeQuotationRepo,
		writeProductRepo:   writeProductRepo,
	}
}

func (w *WriteQuotationService) CreateQuotation(ctx context.Context, quotation Quotation) (*Quotation, error) {
	if quotation.TotalPrice <= 0 {
		return nil, errors.New("el valor total de la cotizacion debe de ser mayor a 0")
	}
	quotation, err := w.writeQuotationRepo.SaveQuotation(ctx, quotation)
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}
