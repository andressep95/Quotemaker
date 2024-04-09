package application

import (
	"context"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/product"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
)

type ReadProductUseCase struct {
	readProductUseCase *domain.ReadProductService
}

func NewReadProductUseCase(readProductUseCase *domain.ReadProductService) *ReadProductUseCase {
	return &ReadProductUseCase{
		readProductUseCase: readProductUseCase,
	}
}

// Execute ejecuta la lógica del caso de uso de listar productos.
func (r *ReadProductUseCase) ListProductByName(ctx context.Context, request *dto.ListProductsRequest) (*response.Response, error) {
	products, err := r.readProductUseCase.ListProductsByName(ctx, request.Limit, request.Offset, request.Name)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "product_list_failed",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}

	productDTOs := make([]dto.ProductDTO, len(products))
	for i, p := range products {
		productDTOs[i] = dto.ProductDTO{
			ID:           p.ID,
			CategoryName: p.CategoryName,
			Code:         p.Code,
			Description:  p.Description,
			Price:        p.Price,
			Weight:       p.Weight,
			Length:       p.Length,
			IsAvailable:  p.IsAvailable,
		}
	}

	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Products listed successfully",
		Data: response.ResponseData{
			Result: &dto.ListProductsResponse{
				Products: productDTOs,
			},
		},
	}, nil
}

func (r *ReadProductUseCase) ListProductByCategory(ctx context.Context, request dto.ListProductByCategoryRequest) (*response.Response, error) {
	products, err := r.readProductUseCase.ListProductByCategory(ctx, request.CategoryName)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "product_list_failed",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}
	// Mapea los productos a DTOs para el response.
	productDTOs := make([]dto.ProductDTO, len(products))
	for i, p := range products {
		productDTOs[i] = dto.ProductDTO{
			ID:           p.ID,
			CategoryName: p.CategoryName,
			Code:         p.Code,
			Description:  p.Description,
			Price:        p.Price,
			Weight:       p.Weight,
			Length:       p.Length,
			IsAvailable:  p.IsAvailable,
		}
	}
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Products listed successfully",
		Data: response.ResponseData{
			Result: &dto.ListProductsResponse{
				Products: productDTOs,
			},
		},
	}, nil
}

func (r *ReadProductUseCase) GetProductByID(ctx context.Context, request *dto.GetProductByIDRequest) (*response.Response, error) {
	product, err := r.readProductUseCase.GetProductByID(ctx, request.ID)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			ErrorCode:  "product_not_found",
			Errors: []response.ErrorDetail{
				{Message: "No product found with the provided ID"},
			},
			Data: response.ResponseData{},
		}, nil
	}

	productDTO := dto.GetProductByIDResponse{
		ID:           product.ID,
		CategoryName: product.CategoryName,
		Code:         product.Code,
		Description:  product.Description,
		Price:        product.Price,
		Weight:       product.Weight,
		Length:       product.Length,
		IsAvailable:  product.IsAvailable,
	}

	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Product retrieved successfully",
		Data: response.ResponseData{
			Result: &productDTO,
		},
	}, nil
}
