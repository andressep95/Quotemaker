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
	category := utiltest.CreateRandomCategory(t)
	newCategory, err := writeRepo.SaveCategory(ctx, category)
	if err != nil {
		return
	}

	// Update catefory name
	newCategory.CategoryName = "New Category Name"
	updatedCategory, err := writeRepo.UpdateCategory(ctx, newCategory)
	if err != nil {
		return
	}
	require.NoError(t, err)

	// Verify
	require.NoError(t, err)
	require.Equal(t, newCategory.CategoryName, updatedCategory.CategoryName)
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
