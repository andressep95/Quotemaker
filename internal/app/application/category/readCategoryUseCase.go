package application

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
)

type ReadCategoryUseCase struct {
	readCategoryService *domain.ReadCategoryService
}

func NewReadCategoryUseCase(readCategoryService *domain.ReadCategoryService) *ReadCategoryUseCase {
	return &ReadCategoryUseCase{
		readCategoryService: readCategoryService,
	}
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

func (l *ReadCategoryUseCase) ListCategorys(ctx context.Context, request *ListCategorysRequest) (*ListCategorysResponse, error) {
	categorys, err := l.readCategoryService.ListCategorys(ctx, request.Limit, request.Offset)
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
