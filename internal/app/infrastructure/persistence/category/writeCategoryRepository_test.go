package persistence

import (
	"context"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func TestSaveCategory(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	category := utiltest.CreateRandomCategory(t, db)
	require.NotEmpty(t, category.ID)
}

func TestUpdateCategory(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteCategoryRepository(db)
	category := utiltest.CreateRandomCategory(t, db)

	// Update category name
	category.CategoryName = "New Category Name"
	updatedCategory, err := writeRepo.UpdateCategory(ctx, category)
	require.NoError(t, err)
	require.Equal(t, category.CategoryName, updatedCategory.CategoryName)

	// Try to update a category that doesn't exist
	nonExistentCategory := category
	nonExistentCategory.ID = "non-existent-id"
	_, err = writeRepo.UpdateCategory(ctx, nonExistentCategory)
	require.Error(t, err)
}

func TestDeleteCategory(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteCategoryRepository(db)
	readRepo := NewReadCategoryRepository(db)
	newCategory := utiltest.CreateRandomCategory(t, db)

	// delete product
	err := writeRepo.DeleteCategory(ctx, newCategory.ID)
	require.NoError(t, err)

	//verify
	_, err = readRepo.GetCategoryByID(ctx, newCategory.ID)
	require.Error(t, err)
}
