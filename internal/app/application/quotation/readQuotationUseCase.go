package application

import (
	"context"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/quotation"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
)

func (r *ReadQuotationUseCase) ListQuotations(ctx context.Context, request dto.ListQuotationsRequest) (*response.Response, error) {
	quotations, err := r.readQuotationService.ListQuotations(ctx, request.Limit, request.Offset)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "internal_server_error",
			Errors:     []response.ErrorDetail{{Message: err.Error()}},
		}, nil
	}

	var quotationsResponse []dto.QuotationDTO
	for _, q := range quotations {
		productsResponse := make([]dto.QuoteProductDetail, len(q.Products))
		for i, prod := range q.Products {
			productsResponse[i] = dto.QuoteProductDetail{
				Product:  prod.Product,
				Quantity: prod.Quantity,
			}
		}

		quotationsResponse = append(quotationsResponse, dto.QuotationDTO{
			ID:          q.ID,
			CreatedAt:   q.CreatedAt,
			TotalPrice:  q.TotalPrice,
			IsPurchased: q.IsPurchased,
			IsDelivered: q.IsDelivered,
			Products:    productsResponse,
		})
	}

	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Quotations listed successfully",
		Data: response.ResponseData{
			Result: dto.ListQuotationsResponse{
				Quotations: quotationsResponse,
			},
		},
	}, nil
}

type ReadQuotationUseCase struct {
	readQuotationService *domain.ReadQuotationService
}

func NewReadQuotationuseCase(readQuotationService *domain.ReadQuotationService) *ReadQuotationUseCase {
	return &ReadQuotationUseCase{
		readQuotationService: readQuotationService,
	}
}
