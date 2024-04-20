package domain

import (
	"context"
	"errors"
)

type WriteCategoryService struct {
	writeCategoryRepo WriteCategoryRepository
}

func NewWriteCategoryService(writeCategoryRepo WriteCategoryRepository) *WriteCategoryService {
	return &WriteCategoryService{
		writeCategoryRepo: writeCategoryRepo,
	}
}

func (s *WriteCategoryService) CreateCategory(ctx context.Context, category Category) (*Category, error) {
	createdCategory, err := s.writeCategoryRepo.SaveCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return &createdCategory, nil
}
func (s *WriteCategoryService) UpdateCategory(ctx context.Context, category Category) (*Category, error) {
	if category.ID == "" {
		return nil, errors.New("category ID is required")
	}

	updatedCategory, err := s.writeCategoryRepo.UpdateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return &updatedCategory, nil
}

func (s *WriteCategoryService) DeleteCategory(ctx context.Context, categoryID string) error {
	if categoryID == "" {
		return errors.New("category ID is required")
	}

	err := s.writeCategoryRepo.DeleteCategory(ctx, categoryID)
	if err != nil {
		return err
	}
	return nil
}
