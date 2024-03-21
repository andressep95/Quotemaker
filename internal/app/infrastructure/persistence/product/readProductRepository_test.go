package persistence

import (
	"context"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func TestListProducts(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	readRepo := NewReadProductRepository(db)

	for i := 0; i < 5; i++ {
		utiltest.CreateRandomProduct(t)
	}

	products, err := readRepo.ListProducts(ctx, 5, 0)
	require.NoError(t, err)

	for _, product := range products {
		require.NotEmpty(t, product)
		require.Len(t, products, 5)
	}
}

func TestListProductsByName(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	readRepo := NewReadProductRepository(db)

	for i := 0; i < 20; i++ {
		utiltest.CreateRandomProduct(t)
	}

	products, err := readRepo.ListProductsByName(ctx, 5, 0, "Pro")
	require.NoError(t, err)

	for _, product := range products {
		require.NotEmpty(t, product)
		require.True(t, true, len(products) <= 5)
	}
}
