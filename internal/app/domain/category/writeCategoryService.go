package domain

import (
	"context"
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
