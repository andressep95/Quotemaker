package application

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
)

// CreateProductRequest define los datos de entrada para crear un producto.
type CreateProductRequest struct {
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	Length      float64 `json:"length"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"`
	Code        string  `json:"code"`
	IsAvailable bool    `json:"is_available"`
}

// CreateProductResponse define los datos de salida tras crear un producto.
type CreateProductResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
}

func (c *CreateProduct) RegisterProduct(ctx context.Context, request *CreateProductRequest) (*CreateProductResponse, error) {
	product := &domain.Product{
		Name:        request.Name,
		CategoryID:  request.CategoryID,
		Length:      request.Length,
		Price:       request.Price,
		Weight:      request.Weight,
		Code:        request.Code,
		IsAvailable: request.IsAvailable,
	}
	createdProduct, err := c.productService.CreateProduct(ctx, *product)
	if err != nil {
		return nil, err
	}
	return &CreateProductResponse{
		ID:         createdProduct.ID,
		Name:       createdProduct.Name,
		CategoryID: createdProduct.CategoryID,
	}, nil
}

type CreateProduct struct {
	productService *domain.ProductService
}

func NewCreateProduct(productService *domain.ProductService) *CreateProduct {
	return &CreateProduct{
		productService: productService,
	}
}
