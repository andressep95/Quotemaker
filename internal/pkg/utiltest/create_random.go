package utiltest

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/config"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/db"
	"github.com/stretchr/testify/require"
)

func SetupTestDB(t *testing.T) *sql.DB {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	database, err := db.NewDBConnection(&config.DB)
	require.NoError(t, err)
	return database
}
