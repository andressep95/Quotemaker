package domain

import (
	"context"
	"fmt"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
)

type ReadProductService struct {
	readProductRepo ReadProductRepository
	categoryRepo    domain.CategoryRepository
}

func NewReadProductService(readProductRepo ReadProductRepository, categoryRepo domain.CategoryRepository) *ReadProductService {
	return &ReadProductService{
		readProductRepo: readProductRepo,
		categoryRepo:    categoryRepo,
	}
}

func (s *ReadProductService) ListProductsByName(ctx context.Context, limit, offset int, name string) ([]Product, error) {
	return s.readProductRepo.ListProductsByName(ctx, limit, offset, name)
}

func (s *ReadProductService) ListProductByCategory(ctx context.Context, categoryName string) ([]Product, error) {
	category, err := s.categoryRepo.GetCategoryByName(ctx, categoryName)
	if err != nil {
		return nil, fmt.Errorf("error al buscar la categoria: %v", err)
	}

	if category.ID < 1 {
		return nil, fmt.Errorf("no se encontró la categoría con el nombre: %s", categoryName)
	}
	return s.readProductRepo.ListProductByCategory(ctx, category.ID)
}

func (s *ReadProductService) GetProductByID(ctx context.Context, productID int) (*Product, error) {
	// Validar que el ID del producto sea válido.
	if productID <= 0 {
		return &Product{}, fmt.Errorf("el ID del producto debe ser mayor que cero")
	}

	// Llamar al repositorio para obtener el producto por su ID.
	product, err := s.readProductRepo.GetProductByID(ctx, productID)
	if err != nil {
		// Manejar cualquier error que pueda surgir durante la obtención del producto.
		return &Product{}, fmt.Errorf("error al obtener el producto por ID %d: %w", productID, err)
	}

	// Devolver el producto obtenido y ningún error.
	return product, nil
}
