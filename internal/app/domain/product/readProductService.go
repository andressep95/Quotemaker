package domain

import (
	"context"
	"fmt"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/product"
)

type ReadProductService struct {
	readProductRepo  ReadProductRepository
	readCategoryRepo domain.ReadCategoryRepository
}

func NewReadProductService(readProductRepo ReadProductRepository, readCategoryRepo domain.ReadCategoryRepository) *ReadProductService {
	return &ReadProductService{
		readProductRepo:  readProductRepo,
		readCategoryRepo: readCategoryRepo,
	}
}

func (s *ReadProductService) ListProductsByName(ctx context.Context, limit, offset int, name string) ([]dto.ProductDTO, error) {
	return s.readProductRepo.ListProductsByName(ctx, limit, offset, name)
}

func (s *ReadProductService) ListProductByCategory(ctx context.Context, categoryName string) ([]dto.ProductDTO, error) {
	category, err := s.readCategoryRepo.GetCategoryByName(ctx, categoryName)
	if err != nil {
		return nil, fmt.Errorf("error al buscar la categoria: %v", err)
	}

	if category.ID == "" {
		return nil, fmt.Errorf("no se encontró la categoría con el nombre: %s", categoryName)
	}
	return s.readProductRepo.ListProductByCategory(ctx, category.CategoryName)
}

func (s *ReadProductService) GetProductByID(ctx context.Context, productID string) (*dto.ProductDTO, error) {
	// Validar que el ID del producto sea válido.
	if productID == "" {
		return &dto.ProductDTO{}, fmt.Errorf("el ID del producto debe ser mayor que cero")
	}

	// Llamar al repositorio para obtener el producto por su ID.
	product, err := s.readProductRepo.GetProductByID(ctx, productID)
	if err != nil {
		// Manejar cualquier error que pueda surgir durante la obtención del producto.
		return &dto.ProductDTO{}, fmt.Errorf("error al obtener el producto por ID %s: %w", productID, err)
	}

	// Devolver el producto obtenido y ningún error.
	return product, nil
}
