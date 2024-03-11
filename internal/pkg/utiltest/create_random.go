package utiltest

import (
	"database/sql"
	"log"
	"math/rand"
	"testing"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/config"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/db"
	"github.com/Andressep/QuoteMaker/internal/pkg/util"
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
	return product
}
