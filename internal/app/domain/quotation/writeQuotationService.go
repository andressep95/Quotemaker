package domain

import (
	"context"
	"errors"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
)

type WriteQuotationService struct {
	writeQuotationRepo WriteQuotationRepository
	writeProductRepo   domain.WriteProductRepository
	readProductRepo    domain.ReadProductRepository
}

func NewWriteQuotationService(writeQuotationRepo WriteQuotationRepository, writeProductRepo domain.WriteProductRepository, readProductRepo domain.ReadProductRepository) *WriteQuotationService {
	return &WriteQuotationService{
		writeQuotationRepo: writeQuotationRepo,
		writeProductRepo:   writeProductRepo,
		readProductRepo:    readProductRepo,
	}
}

func (w *WriteQuotationService) CreateQuotation(ctx context.Context, quotation Quotation) (*Quotation, error) {
	totalPrice := 0.0
	for _, prod := range quotation.Products {
		productDetails, err := w.readProductRepo.GetProductByID(ctx, prod.ProductID)
		if err != nil {
			return &Quotation{}, err
		}
		totalPrice += productDetails.Price * float64(prod.Quantity)
	}
	quotation.TotalPrice = totalPrice
	if quotation.TotalPrice <= 0 {
		return nil, errors.New("el valor total de la cotizacion debe de ser mayor a 0")
	}
	quotation, err := w.writeQuotationRepo.SaveQuotation(ctx, quotation)
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}
