package domain

import (
	"context"
	"errors"
)

type ReadCategoryService struct {
	readCategoryService ReadCategoryRepository
}

func (s *ReadCategoryService) GetCategory(ctx context.Context, id string) (*Category, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	return s.readCategoryService.GetCategoryByID(ctx, id)
}

func (s *ReadCategoryService) GetCategoryByName(ctx context.Context, name string) (Category, error) {
	if name == "" {
		return Category{}, errors.New("name cannot be empty")
	}
	return s.readCategoryService.GetCategoryByName(ctx, name)
}

func (s *ReadCategoryService) ListCategories(ctx context.Context, limit int, offset int) ([]Category, error) {
	if limit < 0 {
		return nil, errors.New("limit cannot be negative")
	}
	if offset < 0 {
		return nil, errors.New("offset cannot be negative")
	}
	return s.readCategoryService.ListCategorys(ctx, limit, offset)
}
func NewReadCategoryService(readCategoryRepository ReadCategoryRepository) *ReadCategoryService {
	return &ReadCategoryService{
		readCategoryService: readCategoryRepository,
	}
}
