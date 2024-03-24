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
	originalProduct := utiltest.CreateRandomProduct(t)

	_, err := db.ExecContext(ctx, "INSERT INTO category (id, category_name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING", originalProduct.CategoryID, "Test Category")
	if err != nil {
		fmt.Println("error:", err)
	}
	newProduct, err := writeRepo.SaveProduct(ctx, originalProduct)
	if err != nil {
		return
	}

	// update
	newProduct.Name = "Product name"
	newProduct.Code = "124214124"
	updatedProduct, err := writeRepo.UpdateProduct(ctx, newProduct)
	require.NoError(t, err)

	require.NoError(t, err)
	require.Equal(t, newProduct.Name, updatedProduct.Name)
	require.Equal(t, newProduct.Code, updatedProduct.Code)
}

func TestDeleteProduct(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	writeRepo := NewWriteProductRepository(db)
	readRepo := NewReadProductRepository(db)
	originalProduct := utiltest.CreateRandomProduct(t)

	_, err := db.ExecContext(ctx, "INSERT INTO category (id, category_name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING", originalProduct.CategoryID, "Test Category")
	if err != nil {
		fmt.Println("error:", err)
	}
	newProduct, err := writeRepo.SaveProduct(ctx, originalProduct)
	if err != nil {
		return
	}
	// delete product
	_ = writeRepo.DeleteProduct(ctx, int(newProduct.ID))
	require.NoError(t, err)

	//verify
	_, err = readRepo.GetProductByID(ctx, int(newProduct.ID))
	require.Error(t, err)
}
