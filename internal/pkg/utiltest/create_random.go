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
	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
	"github.com/stretchr/testify/require"

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

func CreateRandomProduct(t *testing.T, db *sql.DB) domainProd.Product {
	rand.Seed(time.Now().UnixNano())

	// Primero, crea una categoría aleatoria para satisfacer la restricción de clave foránea.
	category := CreateRandomCategory(t, db)

	product := domainProd.Product{
		Description: "Product-" + util.RandomString(8),
		CategoryID:  category.ID, // Asume que CategoryID es una cadena que representa un UUID.
		Price:       util.RandomFloat(100, 500),
		Length:      util.RandomFloat(1, 6),
		Weight:      util.RandomFloat(10, 15),
		Code:        "Code-" + util.RandomString(8),
		IsAvailable: true,
	}

	// Asegúrate de ajustar los nombres de las columnas y los tipos según tu esquema actualizado.
	query := `INSERT INTO product (id, category_id, code, description, price, weight, length, is_available)
	VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := db.QueryRow(query, product.CategoryID, product.Code, product.Description, product.Price, product.Weight, product.Length, product.IsAvailable).Scan(&product.ID)
	require.NoError(t, err, "CreateRandomProduct: failed to insert product")

	return product
}

func CreateRandomCategory(t *testing.T, db *sql.DB) domainCat.Category {
	rand.Seed(time.Now().UnixNano())
	categoryName := util.RandomString(5)
	var categoryId string

	query := `INSERT INTO category (category_name) VALUES ($1) RETURNING id`
	err := db.QueryRow(query, categoryName).Scan(&categoryId)
	if err != nil {
		t.Fatalf("CreateRandomCategory: failed to insert category: %v", err)
	}
	return domainCat.Category{ID: categoryId, CategoryName: categoryName}
}

func CreateRandomQuoteProduct(t *testing.T, db *sql.DB, quotationId string) domain.QuoteProduct {
	rand.Seed(time.Now().UnixNano())

	product := CreateRandomProduct(t, db)

	quoteProduct := domain.QuoteProduct{
		QuotationID: quotationId,
		ProductID:   product.ID,
		Quantity:    util.RandomInt(1, 10),
	}

	_, err := db.Exec("INSERT INTO quote_product (quotation_id, product_id, quantity) VALUES ($1, $2, $3)", quoteProduct.QuotationID, quoteProduct.ProductID, quoteProduct.Quantity)
	if err != nil {
		t.Fatalf("CreateRandomQuoteProduct: failed to insert quote product: %v", err)
	}

	return quoteProduct
}

func CreateRandomQuotation(t *testing.T, db *sql.DB) domain.Quotation {
	quotation := domain.Quotation{
		CreatedAt:   time.Now(),
		TotalPrice:  util.RandomFloat(100, 500),
		IsPurchased: false,
		IsDelivered: false,
	}

	// Inserta la cotización en la base de datos
	query := `INSERT INTO quotation (created_at, total_price, is_purchased, is_delivered) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(query, quotation.CreatedAt, quotation.TotalPrice, quotation.IsPurchased, quotation.IsDelivered).Scan(&quotation.ID)
	require.NoError(t, err, "CreateRandomQuotation: failed to insert quotation")

	// Ahora que la cotización existe, inserta productos relacionados.
	quotation.Products = make([]domain.QuoteProduct, 0, 5)
	for i := 0; i < 5; i++ {
		quoteProduct := CreateRandomQuoteProduct(t, db, quotation.ID)
		quotation.Products = append(quotation.Products, quoteProduct)
	}

	return quotation
}
