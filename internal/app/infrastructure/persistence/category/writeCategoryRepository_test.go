package persistence

import (
	"context"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func TestSaveCategory(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteCategoryRepository(db)
	category := utiltest.CreateRandomCategory(t)

	savesCategory, err := writeRepo.SaveCategory(ctx, category)
	require.NoError(t, err)
	require.NotEmpty(t, savesCategory.ID)
	require.Equal(t, category.CategoryName, savesCategory.CategoryName)

}

func TestUpdateCategory(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteCategoryRepository(db)
	readRepo := NewReadCategoryRepository(db)
	originalCategory := utiltest.CreateRandomCategory(t)

	// Update catefory name
	originalCategory.CategoryName = "New Category Name"
	err := writeRepo.UpdateCategory(ctx, originalCategory)
	require.NoError(t, err)

	// Verify
	updatedCategory, err := readRepo.GetCategoryByID(ctx, originalCategory.ID)
	require.NoError(t, err)
	require.Equal(t, originalCategory.CategoryName, updatedCategory.CategoryName)
}

func TestDeleteCategory(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteCategoryRepository(db)
	readRepo := NewReadCategoryRepository(db)
	newCategory := utiltest.CreateRandomCategory(t)

	// delete product
	err := writeRepo.DeleteCategory(ctx, int(newCategory.ID))
	require.NoError(t, err)

	//verify
	_, err = readRepo.GetCategoryByID(ctx, int(newCategory.ID))
	require.Error(t, err)
}
