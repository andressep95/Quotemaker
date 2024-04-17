package application

import (
	"context"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/quotation"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
)

// RegisterQuotation maneja la lógica de negocio para registrar una nueva cotización
func (w *WriteQuotationUseCase) RegisterQuotation(ctx context.Context, request dto.CreateQuotationRequest) (*response.Response, error) {
	// Convierte los productos del DTO a entidades de dominio
	var quoteProducts []domain.QuoteProduct
	for _, p := range request.Products {
		quoteProducts = append(quoteProducts, domain.QuoteProduct{
			ProductID: p.Product.ID,
			Quantity:  p.Quantity,
		})
	}

	// Crea una entidad de cotización con los productos convertidos
	quotation := domain.Quotation{
		TotalPrice:  request.TotalPrice,
		IsPurchased: request.IsPurchased,
		Products:    quoteProducts,
	}

	// Llama al servicio de dominio para crear la cotización
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

	// Convierte los productos de la cotización creada a DTOs para la respuesta
	productsResponse := make([]dto.QuoteProductDetail, len(createdQuotation.Products))
	for i, prod := range createdQuotation.Products {
		productDTO, _ := w.writeQuotationService.GetProductDetails(ctx, prod.ProductID) // Asume una función en el servicio que devuelve detalles del producto
		productsResponse[i] = dto.QuoteProductDetail{
			Product:  productDTO,
			Quantity: prod.Quantity,
		}
	}

	// Prepara la respuesta final utilizando DTOs
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Quotation created successfully",
		Data: response.ResponseData{
			Result: dto.CreateQuotationResponse{
				ID:          createdQuotation.ID,
				CreatedAt:   createdQuotation.CreatedAt,
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
