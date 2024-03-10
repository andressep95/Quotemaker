package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	"github.com/lib/pq"
)

type ProductService struct {
	productRepo  ProductRepository
	categoryRepo domain.CategoryRepository
}

func NewProductService(productRepo ProductRepository, categoryRepo domain.CategoryRepository) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

// CreateProduct valida y crea un nuevo producto.
func (s *ProductService) CreateProduct(ctx context.Context, product Product) (*Product, error) {
	if product.Name == "" {
		return nil, errors.New("el nombre del producto no puede estar vacío")
	}
	if product.Price <= 0 {
		return nil, errors.New("el precio del producto debe ser positivo")
	}
	if product.Length <= 0 || product.Weight <= 0 {
		return nil, errors.New("la longitud y el peso del producto deben ser positivos")
	}
	if product.Code == "" {
		return nil, errors.New("el código del producto no puede estar vacío")
	}
	if product.CategoryID < 1 {
		category, err := s.categoryRepo.GetCategoryByName(ctx, "No Asignada")
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				createdCategory, createErr := s.categoryRepo.SaveCategory(ctx, domain.Category{CategoryName: "No Asignada"})
				if createErr != nil {
					// Verifica si el error es debido a una clave duplicada
					var pgErr *pq.Error
					if errors.As(createErr, &pgErr) && pgErr.Code == "23505" {
						// Si el error es por clave duplicada, puede ser una condición de carrera, reintenta obtener la categoría
						category, err = s.categoryRepo.GetCategoryByName(ctx, "No Asignada")
						if err != nil {
							// Manejar el error si aún falla
							return nil, fmt.Errorf("error al obtener la categoría 'No Asignada' después de un conflicto de clave duplicada: %w", err)
						}
						product.CategoryID = category.ID
					} else {
						// Si el error es otro, manéjalo adecuadamente
						return nil, fmt.Errorf("error al crear la categoría 'No Asignada': %w", createErr)
					}
				} else {
					product.CategoryID = createdCategory.ID
				}
			} else {
				return nil, fmt.Errorf("error al buscar la categoría 'No Asignada': %w", err)
			}
		} else {
			product.CategoryID = category.ID
		}
	}
	// Continúa con guardar el producto en la base de datos
	id, err := s.productRepo.SaveProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (s *ProductService) GetProductByID(ctx context.Context, productID int) (*Product, error) {
	// Validar que el ID del producto sea válido.
	if productID <= 0 {
		return &Product{}, fmt.Errorf("el ID del producto debe ser mayor que cero")
	}

	// Llamar al repositorio para obtener el producto por su ID.
	product, err := s.productRepo.GetProductByID(ctx, productID)
	if err != nil {
		// Manejar cualquier error que pueda surgir durante la obtención del producto.
		return &Product{}, fmt.Errorf("error al obtener el producto por ID %d: %w", productID, err)
	}

	// Devolver el producto obtenido y ningún error.
	return product, nil
}

func (s *ProductService) ListProductsByName(ctx context.Context, limit, offset int, name string) ([]Product, error) {
	return s.productRepo.ListProductsByName(ctx, limit, offset, name)
}

func (s *ProductService) ListProductByCategory(ctx context.Context, categoryName string) ([]Product, error) {
	category, err := s.categoryRepo.GetCategoryByName(ctx, categoryName)
	if err != nil {
		return nil, fmt.Errorf("error al buscar la categoria: %v", err)
	}

	if category.ID < 1 {
		return nil, fmt.Errorf("no se encontró la categoría con el nombre: %s", categoryName)
	}
	return s.productRepo.ListProductByCategory(ctx, category.ID)
}

func (s *ProductService) UpdateProduct(ctx context.Context, product Product) (*Product, error) {

	if product.ID <= 0 {
		return nil, errors.New("el ID del producto debe ser mayor que cero")
	}
	if product.Name == "" {
		return nil, errors.New("el nombre del producto no puede estar vacío")
	}
	if product.Price < 0 {
		return nil, errors.New("el precio del producto no puede ser negativo")
	}
	// Puedes agregar más validaciones según sea necesario

	// Actualizar el producto en la base de datos a través del repositorio
	updatedProduct, err := s.productRepo.UpdateProduct(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar el producto: %v", err)
	}

	return &updatedProduct, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, productID int) error {
	return s.productRepo.DeleteProduct(ctx, productID)
}
