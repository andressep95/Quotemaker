package utiltest

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	domainCat "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	domainProd "github.com/Andressep/QuoteMaker/internal/app/domain/product"

	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/config"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/db"
	"github.com/Andressep/QuoteMaker/internal/pkg/util"
	_ "github.com/lib/pq"
)

func SetupTestDB(t *testing.T) *sql.DB {
	config, err := config.LoadConfig("../../../../..")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
	}
	database, err := db.DBConnection(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		log.Fatal("cannot connect to database:", err)
	}
	return database
}

func CreateRandomProduct(t *testing.T) domainProd.Product {
	rand.Seed(time.Now().UnixNano())

	product := domainProd.Product{
		Description: "Product-" + util.RandomString(8),
		CategoryID:  util.RandomInt(1, 100),
		Price:       util.RandomFloat(100, 500),
		Length:      util.RandomFloat(1, 6),
		Weight:      util.RandomFloat(10, 15),
		Code:        "Code-" + util.RandomString(8),
		IsAvailable: true,
	}

	return product
}

func CreateRandomCategory(t *testing.T) domainCat.Category {
	rand.Seed(time.Now().UnixNano())

	category := domainCat.Category{
		CategoryName: util.RandomString(5),
	}
	return category
}
