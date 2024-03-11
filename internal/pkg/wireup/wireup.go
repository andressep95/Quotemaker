package wireup

import (
	"database/sql"

	appCategory "github.com/Andressep/QuoteMaker/internal/app/application/category"
	application "github.com/Andressep/QuoteMaker/internal/app/application/product"

	categoryServ "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"

	categoryRepo "github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/category"
	persistence "github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/product"

	catController "github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/controller/category"
	controller "github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/controller/product"

	"github.com/gin-gonic/gin"
)

func SetupAppControllers(r *gin.Engine, db *sql.DB) {
	// Repository´s
	readProductRepo := persistence.NewReadProductRepository(db)
	writeProductRepo := persistence.NewWriteProductRepository(db)
	categoryRepo := categoryRepo.NewCategoryRepository(db)

	// Services´s
	readProductService := domain.NewReadProductService(readProductRepo, categoryRepo)
	writeProductService := domain.NewWriteProductService(writeProductRepo, categoryRepo)
	categoryService := categoryServ.NewCategoryService(categoryRepo)

	// Usecase´s
	readProductUseCase := application.NewReadProductUseCase(readProductService)
	writeProductUseCase := application.NewWriteProductUseCase(writeProductService, readProductService)
	categoryListUseCase := appCategory.NewListCategory(categoryService)
	categoryCreateUseCase := appCategory.NewCreateCategory(categoryService)

	// Handler´s
	readProductHandler := controller.NewReadProductHandler(readProductUseCase)
	writeProductHandler := controller.NewWriteProductHandler(writeProductUseCase)
	categoryHandler := catController.NewCategoryController(categoryCreateUseCase, categoryListUseCase)

	// Routes
	readProductHandler.ReadProductRouter(r)
	writeProductHandler.WriteProductRouter(r)
	categoryHandler.CategoryRouter(r)
}
