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
type UpdateProductRequest struct {
	ID          int     `json:"id"`
	Name        string  `json:"name,omitempty"`
	CategoryID  int     `json:"category_id,omitempty"`
	Length      float64 `json:"length,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
	Code        string  `json:"code,omitempty"`
	IsAvailable bool    `json:"is_available,omitempty"`
}

// CreateProductResponse define los datos de salida tras crear un producto.
type CreateProductResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
}

type DeleteProductRequest struct {
	ID int `json:"id"`
}

type DeleteProductResponse struct {
	Message string `json:"message"`
}

func (w *WriteProductUseCase) RegisterProduct(ctx context.Context, request *CreateProductRequest) (*CreateProductResponse, error) {
	product := &domain.Product{
		Name:        request.Name,
		CategoryID:  request.CategoryID,
		Length:      request.Length,
		Price:       request.Price,
		Weight:      request.Weight,
		Code:        request.Code,
		IsAvailable: request.IsAvailable,
	}
	createdProduct, err := w.writeProductService.CreateProduct(ctx, *product)
	if err != nil {
		return nil, err
	}
	return &CreateProductResponse{
		ID:         createdProduct.ID,
		Name:       createdProduct.Name,
		CategoryID: createdProduct.CategoryID,
	}, nil
}

func (w *WriteProductUseCase) ModifyProduct(ctx context.Context, request *UpdateProductRequest) (*CreateProductResponse, error) {
	// Primero, obtén el producto existente que se desea modificar
	existingProduct, err := w.readProductService.GetProductByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	// Actualiza los campos del producto existente con los valores proporcionados en la solicitud
	existingProduct.Name = request.Name
	existingProduct.CategoryID = request.CategoryID
	existingProduct.Length = request.Length
	existingProduct.Price = request.Price
	existingProduct.Weight = request.Weight
	existingProduct.Code = request.Code
	existingProduct.IsAvailable = request.IsAvailable

	// Llama al servicio de dominio para modificar el producto en la base de datos
	updatedProduct, err := w.writeProductService.UpdateProduct(ctx, *existingProduct)
	if err != nil {
		return nil, err
	}

	// Devuelve una respuesta con los detalles del producto modificado
	return &CreateProductResponse{
		ID:         updatedProduct.ID,
		Name:       updatedProduct.Name,
		CategoryID: updatedProduct.CategoryID,
	}, nil
}

// Execute ejecuta la lógica del caso de uso de eliminar un producto.
func (w *WriteProductUseCase) DeleteProduct(ctx context.Context, request *DeleteProductRequest) (*DeleteProductResponse, error) {
	// Llama al servicio de dominio para eliminar el producto.
	err := w.writeProductService.DeleteProduct(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	// Puedes agregar lógica adicional aquí si es necesario.

	return &DeleteProductResponse{
		Message: "Product deleted successfully",
	}, nil
}

type WriteProductUseCase struct {
	writeProductService *domain.WriteProductService
	readProductService  *domain.ReadProductService
}

func NewWriteProductUseCase(writeProductService *domain.WriteProductService, readProductService *domain.ReadProductService) *WriteProductUseCase {
	return &WriteProductUseCase{
		writeProductService: writeProductService,
		readProductService:  readProductService,
	}
}
