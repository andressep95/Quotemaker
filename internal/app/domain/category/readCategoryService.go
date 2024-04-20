package domain

import "context"

type ReadCategoryService struct {
	readCategoryService ReadCategoryRepository
}

func (s *ReadCategoryService) GetCategory(ctx context.Context, id string) (*Category, error) {
	return s.readCategoryService.GetCategoryByID(ctx, id)
}
func (s *ReadCategoryService) GetCategoryByName(ctx context.Context, name string) (Category, error) {
	return s.readCategoryService.GetCategoryByName(ctx, name)
}

func (s *ReadCategoryService) ListCategories(ctx context.Context, limit int, offset int) ([]Category, error) {
	return s.readCategoryService.ListCategorys(ctx, limit, offset)
}
func NewReadCategoryService(readCategoryRepository ReadCategoryRepository) *ReadCategoryService {
	return &ReadCategoryService{
		readCategoryService: readCategoryRepository,
	}
}

func (s *ReadCategoryService) ListCategorys(ctx context.Context, limit, offset int) ([]Category, error) {
	return s.readCategoryService.ListCategorys(ctx, limit, offset)
}
