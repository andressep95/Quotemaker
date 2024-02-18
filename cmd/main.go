package main

import (
	"log"

	application "github.com/Andressep/QuoteMaker/internal/app/application/product"
	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/config"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/db"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/category"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/product"
	controller "github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/controller/product"
	"github.com/labstack/echo"
)

func main() {
	// Inicia el servidor
	e := echo.New()
	// Suponiendo que tienes una función para cargar tu configuración
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}
	// Crea la conexión a la base de datos
	db, err := db.NewDBConnection(&cfg.DB)
	if err != nil {
		log.Fatalf("No se pudo establecer la conexión a la base de datos: %v", err)
	}
	defer db.Close()

	productRepo := product.NewProductRepository(db)
	categoryRepo := category.NewCategoryRepository(db)
	productService := domain.NewProductService(productRepo, categoryRepo)
	productUseCase := application.NewCreateProduct(productService)
	productController := controller.NewProductController(productUseCase)

	controller.ProductRouter(e, productController)
	// Inicia el servidor
	e.Logger.Fatal(e.Start(":8080"))

}
