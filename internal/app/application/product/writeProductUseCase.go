package application

import (
	"context"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/product"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
)

func (w *WriteProductUseCase) RegisterProduct(ctx context.Context, request *dto.CreateProductRequest) (*response.Response, error) {
	product := &domain.Product{
		CategoryID:  request.CategoryID,
		Code:        request.Code,
		Description: request.Description,
		Price:       request.Price,
		Weight:      request.Weight,
		Length:      request.Length,
		IsAvailable: request.IsAvailable,
	}
	createdProduct, err := w.writeProductService.CreateProduct(ctx, *product)
	if err != nil {
		return &response.Response{
			Status:     "error creating product",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "internal_server_error",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Product created successfully",
		Data: response.ResponseData{
			Result: &dto.CreateProductResponse{
				ID:          createdProduct.ID,
				Description: createdProduct.Description,
				CategoryID:  createdProduct.CategoryID,
			},
		},
	}, nil
}
func (w *WriteProductUseCase) ModifyProduct(ctx context.Context, request *dto.UpdateProductRequest) (*response.Response, error) {
	// Primero, obtén el producto existente que se desea modificar
	existingProduct, err := w.readProductService.GetProductByID(ctx, request.ID)
	if err != nil {
		return &response.Response{
			Status:     "error searching the product by ID",
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			ErrorCode:  "product not found",
			Errors: []response.ErrorDetail{
				{Message: "No product found with the provided ID"},
			},
			Data: response.ResponseData{},
		}, nil
	}

	existingProduct.Code = request.Code
	existingProduct.Description = request.Description
	existingProduct.Price = request.Price
	existingProduct.Weight = request.Weight
	existingProduct.Length = request.Length
	existingProduct.IsAvailable = request.IsAvailable

	product := domain.Product{
		Code:        existingProduct.Code,
		CategoryID:  request.CategoryID,
		Description: existingProduct.Description,
		Price:       existingProduct.Price,
		Weight:      existingProduct.Weight,
		Length:      existingProduct.Length,
		IsAvailable: existingProduct.IsAvailable,
	}

	// Llama al servicio de dominio para modificar el producto en la base de datos
	updatedProduct, err := w.writeProductService.UpdateProduct(ctx, product)
	if err != nil {
		return &response.Response{
			Status:     "error updating product",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "internal_server_error",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}

	// Devuelve una respuesta con los detalles del producto modificado
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Product created successfully",
		Data: response.ResponseData{
			Result: &dto.CreateProductResponse{
				ID:          updatedProduct.ID,
				Description: updatedProduct.Description,
				CategoryID:  updatedProduct.CategoryID,
			},
		},
	}, nil
}

// Execute ejecuta la lógica del caso de uso de eliminar un producto.
func (w *WriteProductUseCase) DeleteProduct(ctx context.Context, request *dto.DeleteProductRequest) (*response.Response, error) {
	err := w.writeProductService.DeleteProduct(ctx, request.ID)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "product_deletion_failed",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}

	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Product deleted successfully",
		Data: response.ResponseData{
			Result: &dto.DeleteProductResponse{
				Message: "Product deleted successfully",
			},
		},
	}, nil
}

type WriteProductUseCase struct {
	writeProductService *domain.WriteProductService
	readProductService  *domain.ReadProductService
}

func NewWriteProductUseCase(writeProductService *domain.WriteProductService, readProductService *domain.ReadProductService) *WriteProductUseCase {
	return &WriteProductUseCase{
		writeProductService: writeProductService,
		readProductService:  readProductService,
	}
}
