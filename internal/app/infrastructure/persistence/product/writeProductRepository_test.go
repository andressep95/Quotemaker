package persistence

import (
	"context"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func TestSaveProduct(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteProductRepository(db)
	product := utiltest.CreateRandomProduct(t, db)

	savedProduct, err := writeRepo.SaveProduct(ctx, product)
	require.NoError(t, err)
	require.NotEmpty(t, savedProduct.ID)
	require.Equal(t, product.Description, savedProduct.Description)
	require.Equal(t, product.IsAvailable, savedProduct.IsAvailable)
}

func TestUpdateProduct(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteProductRepository(db)
	originalProduct := utiltest.CreateRandomProduct(t, db)

	newProduct, err := writeRepo.SaveProduct(ctx, originalProduct)
	if err != nil {
		return
	}

	// update
	newProduct.Description = "Product name"
	newProduct.Code = "124214124"
	updatedProduct, err := writeRepo.UpdateProduct(ctx, newProduct)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, newProduct.Description, updatedProduct.Description)
	require.Equal(t, newProduct.Code, updatedProduct.Code)
}

func TestDeleteProduct(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteProductRepository(db)
	readRepo := NewReadProductRepository(db)
	originalProduct := utiltest.CreateRandomProduct(t, db)

	newProduct, err := writeRepo.SaveProduct(ctx, originalProduct)
	if err != nil {
		return
	}
	// delete product
	_ = writeRepo.DeleteProduct(ctx, newProduct.ID)
	require.NoError(t, err)

	//verify
	_, err = readRepo.GetProductByID(ctx, newProduct.ID)
	require.Error(t, err)
}
