package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	"github.com/lib/pq"
)

type WriteProductService struct {
	writeProductRepo  WriteProductRepository
	writeCategoryRepo domain.WriteCategoryRepository
	readCategoryRepo  domain.ReadCategoryRepository
}

func NewWriteProductService(writeProductRepo WriteProductRepository, writeCategoryRepo domain.WriteCategoryRepository, readCategoryRepo domain.ReadCategoryRepository) *WriteProductService {
	return &WriteProductService{
		writeProductRepo:  writeProductRepo,
		writeCategoryRepo: writeCategoryRepo,
		readCategoryRepo:  readCategoryRepo,
	}
}

func (s *WriteProductService) CreateProduct(ctx context.Context, product Product) (*Product, error) {
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
		category, err := s.readCategoryRepo.GetCategoryByName(ctx, "No Asignada")
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				createdCategory, createErr := s.writeCategoryRepo.SaveCategory(ctx, domain.Category{CategoryName: "No Asignada"})
				if createErr != nil {
					// Verifica si el error es debido a una clave duplicada
					var pgErr *pq.Error
					if errors.As(createErr, &pgErr) && pgErr.Code == "23505" {
						// Si el error es por clave duplicada, puede ser una condición de carrera, reintenta obtener la categoría
						category, err = s.readCategoryRepo.GetCategoryByName(ctx, "No Asignada")
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
	id, err := s.writeProductRepo.SaveProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (s *WriteProductService) UpdateProduct(ctx context.Context, product Product) (*Product, error) {

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
	updatedProduct, err := s.writeProductRepo.UpdateProduct(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar el producto: %v", err)
	}

	return &updatedProduct, nil
}

func (s *WriteProductService) DeleteProduct(ctx context.Context, productID int) error {
	return s.writeProductRepo.DeleteProduct(ctx, productID)
}
