package application

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
)

type ListProduct struct {
	productService *domain.ProductService
}

func NewListProduct(productService *domain.ProductService) *ListProduct {
	return &ListProduct{
		productService: productService,
	}
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
func (l *ListProduct) ListProductByName(ctx context.Context, request *ListProductsRequest) (*ListProductsResponse, error) {
	products, err := l.productService.ListProductsByName(ctx, request.Limit, request.Offset, request.Name)
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

func (l *ListProduct) ListProductByCategory(ctx context.Context, request ListProductByCategoryRequest) (*ListProductsResponse, error) {
	products, err := l.productService.ListProductByCategory(ctx, request.CategoryName)
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
