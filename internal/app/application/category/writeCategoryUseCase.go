package application

import (
	"context"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
)

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
}
type CreateCategoryResponse struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name"`
}

func (w *WriteCategoryUseCase) RegisterCategory(ctx context.Context, request *CreateCategoryRequest) (*response.Response, error) {
	category := &domain.Category{
		CategoryName: request.CategoryName,
	}
	createdCategory, err := w.writeCategoryService.CreateCategory(ctx, *category)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "category_creation_failed",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Category created successfully",
		Data: response.ResponseData{
			Result: &CreateCategoryResponse{
				ID:           createdCategory.ID,
				CategoryName: createdCategory.CategoryName,
			},
		},
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
