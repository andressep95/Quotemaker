package application

import (
	"context"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	domainQ "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/quotation"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
)

func (w *WriteQuotationUseCase) RegisterQuotation(ctx context.Context, request dto.CreateQuotationRequest) (*response.Response, error) {
	// Convertimos los detalles del producto de DTO a entidades de dominio
	var quoteProducts []domainQ.QuoteProduct
	for _, pd := range request.Products {
		quoteProducts = append(quoteProducts, domainQ.QuoteProduct{
			ProductID: pd.Product.ID, // Usamos el ID directamente, asumiendo que es válido y existe
			Quantity:  pd.Quantity,
		})
	}

	// Creamos la entidad de cotización para usarla en la lógica de negocio
	quotation := domainQ.Quotation{
		TotalPrice:  request.TotalPrice,
		IsPurchased: request.IsPurchased,
		IsDelivered: false, // Inicialmente, cuando se crea, no está entregada
		Products:    quoteProducts,
	}

	// Llamada al servicio de dominio para crear la cotización
	createdQuotation, err := w.writeQuotationService.CreateQuotation(ctx, quotation)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "internal_server_error",
			Errors:     []response.ErrorDetail{{Message: err.Error()}},
		}, nil
	}

	// Preparamos los productos para la respuesta DTO
	productsResponse := make([]dto.QuoteProductDetail, len(createdQuotation.Products))
	for i, prod := range createdQuotation.Products {
		productDTO, _ := w.readProductService.GetProductByID(ctx, prod.ProductID) // Asume una función que devuelve detalles del producto
		productsResponse[i] = dto.QuoteProductDetail{
			Product:  *productDTO,
			Quantity: prod.Quantity,
		}
	}

	// Construimos la respuesta final usando DTO
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
	writeQuotationService *domainQ.WriteQuotationService
	readProductService    *domain.ReadProductService
}

func NewWriteQuotationUseCase(writeQuotationService *domainQ.WriteQuotationService, readProductService *domain.ReadProductService) *WriteQuotationUseCase {
	return &WriteQuotationUseCase{
		writeQuotationService: writeQuotationService,
		readProductService:    readProductService,
	}
}
