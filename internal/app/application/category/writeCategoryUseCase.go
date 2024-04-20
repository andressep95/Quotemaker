package application

import (
	"context"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/category"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
)

func (w *WriteCategoryUseCase) RegisterCategory(ctx context.Context, request *dto.CreateCategoryRequest) (*response.Response, error) {
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
			Result: &dto.CreateCategoryResponse{
				ID:           createdCategory.ID,
				CategoryName: createdCategory.CategoryName,
			},
		},
	}, nil
}
func (w *WriteCategoryUseCase) UpdateCategory(ctx context.Context, request *dto.UpdateCategoryRequest) (*response.Response, error) {
	category := &domain.Category{
		ID:           request.ID,
		CategoryName: request.CategoryName,
	}
	updatedCategory, err := w.writeCategoryService.UpdateCategory(ctx, *category)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "category_update_failed",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Category updated successfully",
		Data: response.ResponseData{
			Result: &dto.UpdateCategoryResponse{
				ID:           updatedCategory.ID,
				CategoryName: updatedCategory.CategoryName,
			},
		},
	}, nil
}

func (w *WriteCategoryUseCase) DeleteCategory(ctx context.Context, request *dto.DeleteCategoryRequest) (*response.Response, error) {
	err := w.writeCategoryService.DeleteCategory(ctx, request.ID)
	if err != nil {
		return &response.Response{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ErrorCode:  "category_deletion_failed",
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Category deleted successfully",
		Data:       response.ResponseData{},
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
