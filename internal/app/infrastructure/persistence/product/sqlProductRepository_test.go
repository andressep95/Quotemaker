package product

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	"github.com/Andressep/QuoteMaker/internal/pkg/util"
	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func CreateRandomProduct(t *testing.T) domain.Product {
	rand.Seed(time.Now().UnixNano())

	product := domain.Product{
		Name:        "Product-" + util.RandomString(8),
		CategoryID:  util.RandomInt(1, 100),
		Price:       util.RandomFloat(100, 500),
		Length:      util.RandomFloat(1, 6),
		Weight:      util.RandomFloat(10, 15),
		Code:        "Code-" + util.RandomString(8),
		IsAvailable: true,
	}

	db := utiltest.SetupTestDB(t)
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO category (id, category_name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING", product.CategoryID, "Test Category")
	if err != nil {
		fmt.Println("error:", err)
	}

	repo := NewProductRepository(db)

	savedProduct, err := repo.SaveProduct(ctx, product)

	require.NoError(t, err)
	require.NotEqual(t, 0, savedProduct.ID) // Asegúrate de que se generó un ID
	require.NotEmpty(t, savedProduct)
	require.Equal(t, product.Code, savedProduct.Code)
	require.Equal(t, product.Name, savedProduct.Name)
	require.Equal(t, product.Price, savedProduct.Price)
	require.NotZero(t, savedProduct.ID)

	return savedProduct
}

func TestSaveProduct(t *testing.T) {
	CreateRandomProduct(t)
}

func TestDeleteProduct(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewProductRepository(db)
	newProduct := CreateRandomProduct(t)

	// delete product
	err := repo.DeleteProduct(ctx, int(newProduct.ID))
	require.NoError(t, err)

	//verify
	_, err = repo.GetProductByID(ctx, int(newProduct.ID))
	require.Error(t, err)
}

func TestListProducts(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewProductRepository(db)

	for i := 0; i < 5; i++ {
		CreateRandomProduct(t)
	}

	products, err := repo.ListProducts(ctx, 5, 0)
	require.NoError(t, err)

	for _, product := range products {
		require.NotEmpty(t, product)
		require.Len(t, products, 5)
	}
}

func TestUpdateProduct(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewProductRepository(db)
	originalProduct := CreateRandomProduct(t)

	// update
	originalProduct.Name = "Product name"
	originalProduct.Code = "124214124"
	_, err := repo.UpdateProduct(ctx, originalProduct)
	require.NoError(t, err)

	// verify
	updateProduct, err := repo.GetProductByID(ctx, originalProduct.ID)
	require.NoError(t, err)
	require.Equal(t, originalProduct.Name, updateProduct.Name)
	require.Equal(t, originalProduct.Code, updateProduct.Code)
}

func TestListProductsByName(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewProductRepository(db)

	for i := 0; i < 20; i++ {
		CreateRandomProduct(t)
	}

	products, err := repo.ListProductsByName(ctx, 5, 0, "Pro")
	require.NoError(t, err)

	for _, product := range products {
		require.NotEmpty(t, product)
		require.True(t, true, len(products) <= 5)
	}
}
