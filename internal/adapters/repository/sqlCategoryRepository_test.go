package repository

import (
	"context"
	"math/rand"
	"testing"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomCategory(t *testing.T) domain.Category {
	rand.Seed(time.Now().UnixNano())

	category := domain.Category{
		CategoryName: util.RandomString(5),
	}

	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCategoryRepository(db)

	savedCategory, err := repo.SaveCategory(ctx, category)

	require.NoError(t, err)
	require.NotEqual(t, 0, savedCategory.ID)
	require.NotEmpty(t, savedCategory)
	require.Equal(t, category.CategoryName, savedCategory.CategoryName)
	require.NotZero(t, savedCategory.ID)

	return savedCategory
}

func TestSaveCategory(t *testing.T) {
	CreateRandomCategory(t)
}

func TestGetCategoryByID(t *testing.T) {
	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCategoryRepository(db)
	newCategory := CreateRandomCategory(t)

	fetchedCategory, err := repo.GetCategoryByID(ctx, newCategory.ID)
	require.NoError(t, err)
	require.NotNil(t, fetchedCategory)
	require.Equal(t, newCategory.ID, fetchedCategory.ID)
	require.Equal(t, newCategory.CategoryName, fetchedCategory.CategoryName)
}

func TestListCategorys(t *testing.T) {
	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCategoryRepository(db)

	for i := 0; i < 5; i++ {
		CreateRandomProduct(t)
	}

	categorys, err := repo.ListCategorys(ctx, 5, 0)
	require.NoError(t, err)

	for _, category := range categorys {
		require.NotEmpty(t, category)
		require.Len(t, categorys, 5)
	}
}

func TestDeleteCategory(t *testing.T) {
	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCategoryRepository(db)
	newCategory := CreateRandomCategory(t)

	// delete product
	err := repo.DeleteCategory(ctx, int(newCategory.ID))
	require.NoError(t, err)

	//verify
	_, err = repo.GetCategoryByID(ctx, int(newCategory.ID))
	require.Error(t, err)
}

func TestGetCategoryByName(t *testing.T) {
	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCategoryRepository(db)
	newCategory := CreateRandomCategory(t)

	// get category
	fetchedCategory, err := repo.GetCategoryByName(ctx, newCategory.CategoryName)

	require.NoError(t, err)
	require.NotNil(t, fetchedCategory)
	require.Equal(t, newCategory.ID, fetchedCategory.ID)
	require.Equal(t, newCategory.CategoryName, fetchedCategory.CategoryName)
}

func TestUpdateCategory(t *testing.T) {
	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCategoryRepository(db)
	originalCategory := CreateRandomCategory(t)

	// Update catefory name
	originalCategory.CategoryName = "New Category Name"
	err := repo.UpdateCategory(ctx, originalCategory)
	require.NoError(t, err)

	// Verify
	updatedCategory, err := repo.GetCategoryByID(ctx, originalCategory.ID)
	require.NoError(t, err)
	require.Equal(t, originalCategory.CategoryName, updatedCategory.CategoryName)
}
