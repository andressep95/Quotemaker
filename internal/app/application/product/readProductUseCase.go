package application

import (
	"context"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
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

type GetProductByIDRequest struct {
	ID string `json:"id"`
}
type GetProductByIDResponse struct {
	ID          string  `json:"id"`
	CategoryID  string  `json:"category_id"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"`
	Length      float64 `json:"length"`
	IsAvailable bool    `json:"is_available"`
}

// ListProductsRequest define los datos de entrada para listar productos.
type ListProductsRequest struct {
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

// ListProductsResponse define los datos de salida tras listar productos.
type ListProductsResponse struct {
	Products []ProductDTO `json:"products"`
}

type ListProductByCategoryRequest struct {
	CategoryName string `json:"category_name"`
}

type ProductDTO struct {
	ID          string  `json:"id"`
	CategoryID  string  `json:"category_id"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"`
	Length      float64 `json:"length"`
	IsAvailable bool    `json:"is_available"`
}

// Execute ejecuta la lógica del caso de uso de listar productos.
func (r *ReadProductUseCase) ListProductByName(ctx context.Context, request *ListProductsRequest) (*response.Response, error) {
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

	productDTOs := make([]ProductDTO, len(products))
	for i, p := range products {
		productDTOs[i] = ProductDTO{
			ID:          p.ID,
			CategoryID:  p.CategoryID,
			Code:        p.Code,
			Description: p.Description,
			Price:       p.Price,
			Weight:      p.Weight,
			Length:      p.Length,
			IsAvailable: p.IsAvailable,
		}
	}

	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Products listed successfully",
		Data: response.ResponseData{
			Result: &ListProductsResponse{
				Products: productDTOs,
			},
		},
	}, nil
}

func (r *ReadProductUseCase) ListProductByCategory(ctx context.Context, request ListProductByCategoryRequest) (*response.Response, error) {
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
	productDTOs := make([]ProductDTO, len(products))
	for i, p := range products {
		productDTOs[i] = ProductDTO{
			ID:          p.ID,
			CategoryID:  p.CategoryID,
			Code:        p.Code,
			Description: p.Description,
			Price:       p.Price,
			Weight:      p.Weight,
			Length:      p.Length,
			IsAvailable: p.IsAvailable,
		}
	}
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Products listed successfully",
		Data: response.ResponseData{
			Result: &ListProductsResponse{
				Products: productDTOs,
			},
		},
	}, nil
}

func (r *ReadProductUseCase) GetProductByID(ctx context.Context, request *GetProductByIDRequest) (*response.Response, error) {
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

	productDTO := GetProductByIDResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Code:        product.Code,
		Description: product.Description,
		Price:       product.Price,
		Weight:      product.Weight,
		Length:      product.Length,
		IsAvailable: product.IsAvailable,
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
