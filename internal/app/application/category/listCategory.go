package application

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
)

type ListCategory struct {
	categoryService *domain.CategoryService
}

func NewListCategory(categoryService *domain.CategoryService) *ListCategory {
	return &ListCategory{categoryService: categoryService}
}

type ListCategorysRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
type ListCategorysResponse struct {
	Category []CategoryDTO `json:"category"`
}
type CategoryDTO struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

func (l *ListCategory) ListCategorys(ctx context.Context, request *ListCategorysRequest) (*ListCategorysResponse, error) {
	categorys, err := l.categoryService.ListCategorys(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, err
	}

	categoryDTOs := make([]CategoryDTO, len(categorys))
	for i, p := range categorys {
		categoryDTOs[i] = CategoryDTO{
			ID:           p.ID,
			CategoryName: p.CategoryName,
		}
	}

	return &ListCategorysResponse{
		Category: categoryDTOs,
	}, nil
}
