package wireup

import (
	"database/sql"
	"net/http"

	application "github.com/Andressep/QuoteMaker/internal/app/application/product"
	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/category"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/product"
	controller "github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/controller/product"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetupAppControllers(e *echo.Echo, db *sql.DB) {
	// Repository´s
	productRepo := product.NewProductRepository(db)
	categoryRepo := category.NewCategoryRepository(db)

	// Services´s
	productService := domain.NewProductService(productRepo, categoryRepo)

	// Usecase´s
	createProductUseCase := application.NewCreateProduct(productService)
	listProductUseCase := application.NewListProduct(productService)

	// Controller´s
	productController := controller.NewProductController(createProductUseCase, listProductUseCase)

	// middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// Routes
	productController.ProductRouter(e)
}
