package application

import (
	"context"
	"errors"
	"net/http"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/category"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/response"
	"github.com/Andressep/QuoteMaker/internal/pkg/errs"
)

type ReadCategoryUseCase struct {
	readCategoryService *domain.ReadCategoryService
}

func NewReadCategoryUseCase(readCategoryService *domain.ReadCategoryService) *ReadCategoryUseCase {
	return &ReadCategoryUseCase{
		readCategoryService: readCategoryService,
	}
}

func (l *ReadCategoryUseCase) ListCategorys(ctx context.Context, request *dto.ListCategorysRequest) (*response.Response, error) {
	cats, err := l.readCategoryService.ListCategories(ctx, request.Limit, request.Offset)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorCode := "category_list_failed"
		if errors.Is(err, errs.ErrCategoryNotFound) {
			statusCode = http.StatusNotFound
			errorCode = "category_not_found"
		}
		return &response.Response{
			Status:     "error",
			StatusCode: statusCode,
			Message:    err.Error(),
			ErrorCode:  errorCode,
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}

	categoryDTOs := make([]dto.CategoryDTO, len(cats))
	for i, p := range cats {
		categoryDTOs[i] = dto.CategoryDTO{
			ID:           p.ID,
			CategoryName: p.CategoryName,
		}
	}

	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Categories listed successfully",
		Data: response.ResponseData{
			Result: dto.ListCategorysResponse{
				Category: categoryDTOs,
			},
		},
	}, nil
}
func (l *ReadCategoryUseCase) GetCategory(ctx context.Context, request *dto.GetCategoryRequest) (*response.Response, error) {
	cat, err := l.readCategoryService.GetCategory(ctx, request.ID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorCode := "category_get_failed"
		if errors.Is(err, errs.ErrCategoryNotFound) {
			statusCode = http.StatusNotFound
			errorCode = "category_not_found"
		}
		return &response.Response{
			Status:     "error",
			StatusCode: statusCode,
			Message:    err.Error(),
			ErrorCode:  errorCode,
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}
	categoryDTO := dto.CategoryDTO{
		ID:           cat.ID,
		CategoryName: cat.CategoryName,
	}
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Category retrieved successfully",
		Data: response.ResponseData{
			Result: dto.GetCategoryResponse{
				Category: categoryDTO,
			},
		},
	}, nil
}

func (l *ReadCategoryUseCase) SearchCategoryByName(ctx context.Context, request *dto.GetCategoryByNameRequest) (*response.Response, error) {
	cat, err := l.readCategoryService.GetCategoryByName(ctx, request.Name)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorCode := "category_search_failed"
		return &response.Response{
			Status:     "error",
			StatusCode: statusCode,
			Message:    err.Error(),
			ErrorCode:  errorCode,
			Errors: []response.ErrorDetail{
				{Message: err.Error()},
			},
			Data: response.ResponseData{},
		}, nil
	}
	categoryDTO := dto.CategoryDTO{
		ID:           cat.ID,
		CategoryName: cat.CategoryName,
	}
	return &response.Response{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Category searched successfully",
		Data: response.ResponseData{
			Result: dto.GetCategoryResponse{
				Category: categoryDTO,
			},
		},
	}, nil
}
