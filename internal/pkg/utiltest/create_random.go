package utiltest

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/config"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/db"
	_ "github.com/lib/pq"
)

func SetupTestDB(t *testing.T) *sql.DB {
	config, err := config.LoadConfig("../../../../..")
	if err != nil {
		log.Fatal("cannot load the config: ", err)
	}

	database, err := db.NewDBConnection(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	return database
}
