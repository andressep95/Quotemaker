package application

import (
	"context"
	"net/http"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
)

// Para la creación de una cotización, incluyendo productos asociados.
type CreateQuotationRequest struct {
	IsPurchased bool                 `json:"is_purchased"`
	Products    []QuoteProductDetail `json:"products"`
	TotalPrice  float64              `json:"total_price"`
}

// Respuesta tras crear una cotización, reflejando los principales campos de interés.
type CreateQuotationResponse struct {
	ID          string               `json:"id"`
	CreatedAt   time.Time            `json:"created_at"`
	TotalPrice  float64              `json:"total_price"`
	IsPurchased bool                 `json:"is_purchased"`
	IsDelivered bool                 `json:"is_delivered"`
	Products    []QuoteProductDetail `json:"products"`
}

// Para actualizar información de la cotización, permitiendo modificar ciertos campos.
type UpdateQuotationRequest struct {
	ID          string     `json:"id"`
	TotalPrice  *float64   `json:"total_price,omitempty"`
	IsPurchased *bool      `json:"is_purchased,omitempty"`
	IsDelivered *bool      `json:"is_delivered,omitempty"`
	PurchasedAt *time.Time `json:"purchased_at,omitempty"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	// Considerar si se permite actualizar los productos en esta solicitud.
}

type QuoteProductDetail struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (w *WriteQuotationUseCase) RegisterQuotation(ctx context.Context, request CreateQuotationRequest) (*response.Response, error) {
	// llenamos una variable de QuoteProduct con los productos que vienen en la request.
	var quoteProducts []domain.QuoteProduct
	for _, p := range request.Products {
		quoteProducts = append(quoteProducts, domain.QuoteProduct{
			ProductID: p.ProductID,
			Quantity:  p.Quantity,
		})
	}

	quotation := domain.Quotation{
		TotalPrice:  request.TotalPrice,
		IsPurchased: request.IsPurchased,
		Products:    quoteProducts,
	}
	// Utiliza el servicio para crear la cotización.
	createdQuotation, err := w.writeQuotationService.CreateQuotation(ctx, quotation)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "internal_server_error",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
		}, nil
	}
	// Convertir los productos de domain.QuoteProduct a QuoteProductDetail para la respuesta.
	productsResponse := make([]QuoteProductDetail, len(createdQuotation.Products))
	for i, prod := range createdQuotation.Products {
		productsResponse[i] = QuoteProductDetail{
			ProductID: prod.ProductID,
			Quantity:  prod.Quantity,
		}
	}

	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Quotation created successfully",
		Data: response.ResponseData{
			Result: CreateQuotationResponse{
				ID:          createdQuotation.ID,
				CreatedAt:   createdQuotation.CreatedAt, // Asumiendo que el servicio asigna la fecha de creación
				TotalPrice:  createdQuotation.TotalPrice,
				IsPurchased: createdQuotation.IsPurchased,
				IsDelivered: createdQuotation.IsDelivered,
				Products:    productsResponse,
			},
		},
	}, nil
}

type WriteQuotationUseCase struct {
	writeQuotationService *domain.WriteQuotationService
}

func NewWriteQuotationUseCase(writeQuotationService *domain.WriteQuotationService) *WriteQuotationUseCase {
	return &WriteQuotationUseCase{
		writeQuotationService: writeQuotationService,
	}
}
