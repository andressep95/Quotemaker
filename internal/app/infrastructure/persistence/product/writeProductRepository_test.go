package persistence

import (
	"context"
	"fmt"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func TestSaveProduct(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteProductRepository(db)
	product := utiltest.CreateRandomProduct(t)

	_, err := db.ExecContext(ctx, "INSERT INTO category (id, category_name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING", product.CategoryID, "Test Category")
	if err != nil {
		fmt.Println("error:", err)
	}

	savedProduct, err := writeRepo.SaveProduct(ctx, product)
	require.NoError(t, err)
	require.NotEmpty(t, savedProduct.ID)
	require.Equal(t, product.Name, savedProduct.Name)
	require.Equal(t, product.IsAvailable, savedProduct.IsAvailable)
}

func TestUpdateProduct(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteProductRepository(db)
	readRepo := NewReadProductRepository(db)
	originalProduct := utiltest.CreateRandomProduct(t)
	_, err := db.ExecContext(ctx, "INSERT INTO category (id, category_name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING", originalProduct.CategoryID, "Test Category")
	if err != nil {
		fmt.Println("error:", err)
	}

	// update
	originalProduct.Name = "Product name"
	originalProduct.Code = "124214124"
	_, err = writeRepo.UpdateProduct(ctx, originalProduct)
	require.NoError(t, err)

	// verify
	updateProduct, err := readRepo.GetProductByID(ctx, originalProduct.ID)
	require.NoError(t, err)
	require.Equal(t, originalProduct.Name, updateProduct.Name)
	require.Equal(t, originalProduct.Code, updateProduct.Code)
}

func TestDeleteProduct(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteProductRepository(db)
	readRepo := NewReadProductRepository(db)
	newProduct := utiltest.CreateRandomProduct(t)

	// delete product
	err := writeRepo.DeleteProduct(ctx, int(newProduct.ID))
	require.NoError(t, err)

	//verify
	_, err = readRepo.GetProductByID(ctx, int(newProduct.ID))
	require.Error(t, err)
}
