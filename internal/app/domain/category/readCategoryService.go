package domain

import "context"

type ReadCategoryService struct {
	readCategoryService ReadCategoryRepository
}

func NewReadCategoryService(readCategoryRepository ReadCategoryRepository) *ReadCategoryService {
	return &ReadCategoryService{
		readCategoryService: readCategoryRepository,
	}
}

func (s *ReadCategoryService) ListCategorys(ctx context.Context, limit, offset int) ([]Category, error) {
	return s.readCategoryService.ListCategorys(ctx, limit, offset)
}
