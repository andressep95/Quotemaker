package application

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
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
	ID int `json:"id"`
}
type GetProductByIDResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	Length      float64 `json:"length"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"`
	Code        string  `json:"code"`
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
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	Length      float64 `json:"length"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"`
	Code        string  `json:"code"`
	IsAvailable bool    `json:"is_available"`
}

// Execute ejecuta la lógica del caso de uso de listar productos.
func (r *ReadProductUseCase) ListProductByName(ctx context.Context, request *ListProductsRequest) (*ListProductsResponse, error) {
	products, err := r.readProductUseCase.ListProductsByName(ctx, request.Limit, request.Offset, request.Name)
	if err != nil {
		return nil, err
	}

	productDTOs := make([]ProductDTO, len(products))
	for i, p := range products {
		productDTOs[i] = ProductDTO{
			ID:          p.ID,
			Name:        p.Name,
			CategoryID:  p.CategoryID,
			Length:      p.Length,
			Price:       p.Price,
			Weight:      p.Weight,
			Code:        p.Code,
			IsAvailable: p.IsAvailable,
		}
	}

	return &ListProductsResponse{
		Products: productDTOs,
	}, nil
}

func (r *ReadProductUseCase) ListProductByCategory(ctx context.Context, request ListProductByCategoryRequest) (*ListProductsResponse, error) {
	products, err := r.readProductUseCase.ListProductByCategory(ctx, request.CategoryName)
	if err != nil {
		return nil, err
	}
	// Mapea los productos a DTOs para el response.
	productDTOs := make([]ProductDTO, len(products))
	for i, p := range products {
		productDTOs[i] = ProductDTO{
			ID:          p.ID,
			Name:        p.Name,
			CategoryID:  p.CategoryID,
			Length:      p.Length,
			Price:       p.Price,
			Weight:      p.Weight,
			Code:        p.Code,
			IsAvailable: p.IsAvailable,
		}
	}
	return &ListProductsResponse{
		Products: productDTOs,
	}, nil
}

func (r *ReadProductUseCase) GetProductByID(ctx context.Context, request *GetProductByIDRequest) (*GetProductByIDResponse, error) {
	product, err := r.readProductUseCase.GetProductByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	productDTO := GetProductByIDResponse{
		ID:          product.ID,
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Length:      product.Length,
		Price:       product.Price,
		Weight:      product.Weight,
		Code:        product.Code,
		IsAvailable: product.IsAvailable,
	}

	return &productDTO, nil
}
