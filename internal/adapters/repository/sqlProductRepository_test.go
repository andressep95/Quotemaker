package repository

import (
	"context"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/config"
	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/infraestructure/db"
	"github.com/Andressep/QuoteMaker/internal/util"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	testConfig := &config.Config{
		DB: config.DatabaseConfig{
			User:     "root",
			Password: "secret",
			Host:     "localhost", // O la dirección IP de Docker si estás en un entorno diferente
			Port:     5432,
			Name:     "quote_maker", // El nombre de la base de datos; asegúrate de que esto sea correcto
		},
	}

	database, err := db.New(context.Background(), testConfig)
	require.NoError(t, err)

	return database
}

func createRandomProduct(t *testing.T) domain.Product {
	name := util.RandomString(10)
	categoryID := int32(util.RandomInt(1, 10))
	length := float32(util.RandomInt(5, 30))
	price := float64(util.RandomInt(1000, 10000))
	weight := float64(util.RandomInt(5, 12))
	code := util.RandomString(8)

	product := domain.Product{
		Name:       name,
		CategoryID: categoryID,
		Length:     length,
		Price:      price,
		Weight:     weight,
		Code:       code,
	}

	return product
}

func TestSaveProduct(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close() // Asegúrate de cerrar la conexión de la base de datos al finalizar la prueba

	repo := New(db)

	ctx := context.Background()
	product := domain.Product{
		// Inicializa tu producto con datos de prueba
		Name:       "Test Product",
		CategoryID: 1,
		Length:     10.0,
		Price:      100.0,
		Weight:     5.0,
		Code:       "TEST123",
	}

	savedProduct, err := repo.SaveProduct(ctx, product)
	require.NoError(t, err)
	require.NotEmpty(t, savedProduct)

	require.Equal(t, product.Name, savedProduct.Name)
	require.Equal(t, product.CategoryID, savedProduct.CategoryID)
	require.Equal(t, product.Length, savedProduct.Length)
	require.Equal(t, product.Price, savedProduct.Price)
	require.Equal(t, product.Weight, savedProduct.Weight)
	require.Equal(t, product.Code, savedProduct.Code)

}
