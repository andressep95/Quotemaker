package persistence

import (
	"context"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func TestGetCategoryByID(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	readRepo := NewReadCategoryRepository(db)
	writeRepo := NewWriteCategoryRepository(db)
	newCategory := utiltest.CreateRandomCategory(t)

	saveCategory, err := writeRepo.SaveCategory(ctx, newCategory)
	if err != nil {
		return
	}

	fetchedCategory, err := readRepo.GetCategoryByID(ctx, saveCategory.ID)
	require.NoError(t, err)
	require.NotNil(t, fetchedCategory)
	require.Equal(t, saveCategory.ID, fetchedCategory.ID)
	require.Equal(t, saveCategory.CategoryName, fetchedCategory.CategoryName)
}

func TestListCategorys(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	readRepo := NewReadCategoryRepository(db)

	for i := 0; i < 5; i++ {
		utiltest.CreateRandomCategory(t)
	}

	categorys, err := readRepo.ListCategorys(ctx, 5, 0)
	require.NoError(t, err)

	for _, category := range categorys {
		require.NotEmpty(t, category)
		require.Len(t, categorys, 5)
	}
}

func TestGetCategoryByName(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	readRepo := NewReadCategoryRepository(db)
	writeRepo := NewWriteCategoryRepository(db)
	newCategory, err := writeRepo.SaveCategory(ctx, utiltest.CreateRandomCategory(t))
	if err != nil {
		return
	}

	// get category
	fetchedCategory, err := readRepo.GetCategoryByName(ctx, newCategory.CategoryName)

	require.NoError(t, err)
	require.NotNil(t, fetchedCategory)
	require.Equal(t, newCategory.ID, fetchedCategory.ID)
	require.Equal(t, newCategory.CategoryName, fetchedCategory.CategoryName)
}
