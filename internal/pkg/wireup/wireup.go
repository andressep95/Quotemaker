package wireup

import (
	"database/sql"

	catCases "github.com/Andressep/QuoteMaker/internal/app/application/category"
	prodCases "github.com/Andressep/QuoteMaker/internal/app/application/product"
	catService "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	prodService "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	"github.com/gin-gonic/gin"

	catRep "github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/category"
	prodRep "github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/product"
	catContrl "github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/controller/category"
	prodContrl "github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/controller/product"
)

func SetupAppControllers(r *gin.Engine, db *sql.DB) {
	// Repository´s
	categoryRepo := catRep.NewCategoryRepository(db)
	productRepo := prodRep.NewProductRepository(db)

	// Services´s
	categoryService := catService.NewCategoryService(categoryRepo)
	productService := prodService.NewProductService(productRepo, categoryRepo)

	// Usecase´s
	createCategoryUseCase := catCases.NewCreateCategory(categoryService)
	listCategoryUseCase := catCases.NewListCategory(categoryService)
	createProductUseCase := prodCases.NewCreateProduct(productService)
	listProductUseCase := prodCases.NewListProduct(productService)

	// Controller´s
	categoryController := catContrl.NewCategoryController(createCategoryUseCase, listCategoryUseCase)
	productController := prodContrl.NewProductController(createProductUseCase, listProductUseCase)

	// Routes
	productController.ProductRouter(r)
	categoryController.CategoryRouter(r)

}
