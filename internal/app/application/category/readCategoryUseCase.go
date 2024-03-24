package application

import (
	"context"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
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

func (l *ReadCategoryUseCase) ListCategorys(ctx context.Context, request *ListCategorysRequest) (*response.Response, error) {
	categorys, err := l.readCategoryService.ListCategorys(ctx, request.Limit, request.Offset)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "category_list_failed",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}

	categoryDTOs := make([]CategoryDTO, len(categorys))
	for i, p := range categorys {
		categoryDTOs[i] = CategoryDTO{
			ID:           p.ID,
			CategoryName: p.CategoryName,
		}
	}

	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Categories listed successfully",
		Data: response.ResponseData{
			Result: &ListCategorysResponse{
				Category: categoryDTOs,
			},
		},
	}, nil
}
