package util

import (
	"context"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/config"
	"github.com/Andressep/QuoteMaker/internal/infrastructure/db"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func SetupTestDB(t *testing.T) *sqlx.DB {
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
