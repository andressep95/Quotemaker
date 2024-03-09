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

func (c *CreateCategory) RegisterCategory(ctx context.Context, request *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	category := &domain.Category{
		CategoryName: request.CategoryName,
	}
	createdCategory, err := c.categoryService.CreateCategory(ctx, *category)
	if err != nil {
		return nil, err
	}
	return &CreateCategoryResponse{
		ID:           createdCategory.ID,
		CategoryName: createdCategory.CategoryName,
	}, nil
}

type CreateCategory struct {
	categoryService *domain.CategoryService
}

func NewCreateCategory(categoryService *domain.CategoryService) *CreateCategory {
	return &CreateCategory{
		categoryService: categoryService,
	}
}
