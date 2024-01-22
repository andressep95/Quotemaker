package service

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/core/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, p domain.Product) (domain.Product, error) {
	// Lógica para crear un producto
	return s.repo.SaveProduct(ctx, p)
}

func (s *ProductService) GetProduct(ctx context.Context, id int) (*domain.Product, error) {
	// Lógica para obtener un producto
	return s.repo.GetProductByID(ctx, id)
}
