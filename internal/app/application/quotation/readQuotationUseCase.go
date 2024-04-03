package application

import (
	"context"
	"net/http"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
)

type ListQuotationsRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListQuotationsResponse struct {
	Quotations []QuotationDTO `json:"quotations"`
}

type QuotationDTO struct {
	ID          string               `json:"id"`
	CreatedAt   time.Time            `json:"created_at"`
	TotalPrice  float64              `json:"total_price"`
	IsPurchased bool                 `json:"is_purchased"`
	IsDelivered bool                 `json:"is_delivered"`
	Products    []QuoteProductDetail `json:"products"`
}

func (r *ReadQuotationUseCase) ListQuotations(ctx context.Context, request ListQuotationsRequest) (*response.Response, error) {
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

	var quotationsResponse []QuotationDTO
	for _, q := range quotations {
		productsResponse := make([]QuoteProductDetail, len(q.Products))
		for i, prod := range q.Products {
			productsResponse[i] = QuoteProductDetail{
				ProductID: prod.ProductID,
				Quantity:  prod.Quantity,
			}
		}

		quotationsResponse = append(quotationsResponse, QuotationDTO{
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
			Result: ListQuotationsResponse{
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
