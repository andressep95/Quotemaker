package application

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
)

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
}
type CreateCategoryResponse struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

func (w *WriteCategoryUseCase) RegisterCategory(ctx context.Context, request *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	category := &domain.Category{
		CategoryName: request.CategoryName,
	}
	createdCategory, err := w.writeCategoryService.CreateCategory(ctx, *category)
	if err != nil {
		return nil, err
	}
	return &CreateCategoryResponse{
		ID:           createdCategory.ID,
		CategoryName: createdCategory.CategoryName,
	}, nil
}

type WriteCategoryUseCase struct {
	writeCategoryService *domain.WriteCategoryService
}

func NewWriteCategoryUseCase(writeCategoryService *domain.WriteCategoryService) *WriteCategoryUseCase {
	return &WriteCategoryUseCase{
		writeCategoryService: writeCategoryService,
	}
}
