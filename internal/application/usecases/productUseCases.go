package usecases

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/core/domain/service"
)

// application/usecases/product_usecase.go
type ProductUsecase struct {
	productService *service.ProductService
}

func NewProductUsecase(productService *service.ProductService) *ProductUsecase {
	return &ProductUsecase{productService: productService}
}

func (u *ProductUsecase) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	return u.productService.CreateProduct(ctx, product)
}
