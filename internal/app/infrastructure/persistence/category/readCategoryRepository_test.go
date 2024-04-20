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
	newCategory := utiltest.CreateRandomCategory(t, db)

	fetchedCategory, err := readRepo.GetCategoryByID(ctx, newCategory.ID)
	require.NoError(t, err)
	require.NotNil(t, fetchedCategory)
	require.Equal(t, newCategory.ID, fetchedCategory.ID)
	require.Equal(t, newCategory.CategoryName, fetchedCategory.CategoryName)
}

func TestListCategorys(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	readRepo := NewReadCategoryRepository(db)

	// Create some categories
	for i := 0; i < 10; i++ {
		utiltest.CreateRandomCategory(t, db)
	}

	// List categories with limit and offset
	categories, err := readRepo.ListCategorys(ctx, 5, 0)
	require.NoError(t, err)
	require.Len(t, categories, 5)

	// List categories with a larger offset
	categories, err = readRepo.ListCategorys(ctx, 5, 10)
	require.NoError(t, err)
	require.Len(t, categories, 5)
}

func TestGetCategoryByName(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	readRepo := NewReadCategoryRepository(db)
	newCategory := utiltest.CreateRandomCategory(t, db)
	// get category
	fetchedCategory, err := readRepo.GetCategoryByName(ctx, newCategory.CategoryName)

	require.NoError(t, err)
	require.NotNil(t, fetchedCategory)
	require.Equal(t, newCategory.ID, fetchedCategory.ID)
	require.Equal(t, newCategory.CategoryName, fetchedCategory.CategoryName)
}
