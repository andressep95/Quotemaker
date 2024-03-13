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
	readCategoryRepo := categoryRepo.NewReadCategoryRepository(db)
	writeCategoryRepo := categoryRepo.NewWriteCategoryRepository(db)

	// Services´s
	readProductService := domain.NewReadProductService(readProductRepo, readCategoryRepo)
	writeProductService := domain.NewWriteProductService(writeProductRepo, writeCategoryRepo, readCategoryRepo)
	readCategoryService := categoryServ.NewReadCategoryService(readCategoryRepo)
	writeCategoryService := categoryServ.NewWriteCategoryService(writeCategoryRepo)

	// Usecase´s
	readProductUseCase := application.NewReadProductUseCase(readProductService)
	writeProductUseCase := application.NewWriteProductUseCase(writeProductService, readProductService)
	readCategegoryUseCase := appCategory.NewReadCategoryUseCase(readCategoryService)
	writeCategoryUseCase := appCategory.NewWriteCategoryUseCase(writeCategoryService)

	// Handler´s
	readProductHandler := controller.NewReadProductHandler(readProductUseCase)
	writeProductHandler := controller.NewWriteProductHandler(writeProductUseCase)
	readCategoryHandler := catController.NewReadCategoryHandler(writeCategoryUseCase, readCategegoryUseCase)
	writeCategoryHandler := catController.NewWriteCategoryHandler(writeCategoryUseCase)

	// Routes
	readProductHandler.ReadProductRouter(r)
	writeProductHandler.WriteProductRouter(r)
	categoryHandler.CategoryRouter(r)
}
