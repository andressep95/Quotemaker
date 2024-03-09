package domain

import "context"

type CategoryService struct {
	categoryRepo CategoryRepository
}

func NewCategoryService(categoryRepo CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, category Category) (*Category, error) {
	createdCategory, err := s.categoryRepo.SaveCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return &createdCategory, nil
}

func (s *CategoryService) ListCategorys(ctx context.Context, limit, offset int) ([]Category, error) {
	return s.categoryRepo.ListCategorys(ctx, limit, offset)
}
