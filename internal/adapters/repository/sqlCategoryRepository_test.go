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

	return category
}

func TestSaveCategory(t *testing.T) {
	CreateRandomCategory(t)
}
