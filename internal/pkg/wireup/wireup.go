package wireup

import (
	"database/sql"

	appCategory "github.com/Andressep/QuoteMaker/internal/app/application/category"
	application "github.com/Andressep/QuoteMaker/internal/app/application/product"
	appQuotation "github.com/Andressep/QuoteMaker/internal/app/application/quotation"

	categoryServ "github.com/Andressep/QuoteMaker/internal/app/domain/category"
	domain "github.com/Andressep/QuoteMaker/internal/app/domain/product"
	quotationServ "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"

	categoryRepo "github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/category"
	persistence "github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/product"
	quotationRepo "github.com/Andressep/QuoteMaker/internal/app/infrastructure/persistence/quotation"

	catController "github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/controller/category"
	controller "github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/controller/product"
	quotationController "github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/controller/quotation"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/transport/middleware"

	"github.com/gin-gonic/gin"
)

// SetupAppControllers configures the application controllers and sets up the routes for handling HTTP requests.
// It takes in a *gin.Engine instance and a *sql.DB instance as parameters.
// The function sets up the necessary middleware, repositories, services, use cases, and handlers for the application.
// It then registers the routes for handling different types of requests, such as reading and writing quotations, products, and categories.
func SetupAppControllers(r *gin.Engine, db *sql.DB) {
	// Middleware
	r.Use(middleware.CORSMiddleware())

	// Repository´s
	readQuotationRepo := quotationRepo.NewReadQuotationRepository(db)
	writeQuotationRepo := quotationRepo.NewWriteQuotationRepository(db)
	readProductRepo := persistence.NewReadProductRepository(db)
	writeProductRepo := persistence.NewWriteProductRepository(db)
	readCategoryRepo := categoryRepo.NewReadCategoryRepository(db)
	writeCategoryRepo := categoryRepo.NewWriteCategoryRepository(db)

	// Services´s
	readQuotationService := quotationServ.NewReadQuotationService(readQuotationRepo)
	writeQuotationService := quotationServ.NewWriteQuotationService(writeQuotationRepo, writeProductRepo)
	readProductService := domain.NewReadProductService(readProductRepo, readCategoryRepo)
	writeProductService := domain.NewWriteProductService(writeProductRepo, writeCategoryRepo, readCategoryRepo)
	readCategoryService := categoryServ.NewReadCategoryService(readCategoryRepo)
	writeCategoryService := categoryServ.NewWriteCategoryService(writeCategoryRepo)

	// Usecase´s
	readQuotationUseCase := appQuotation.NewReadQuotationuseCase(readQuotationService)
	writeQuotationUseCase := appQuotation.NewWriteQuotationUseCase(writeQuotationService, readProductService)
	readProductUseCase := application.NewReadProductUseCase(readProductService)
	writeProductUseCase := application.NewWriteProductUseCase(writeProductService, readProductService)
	readCategegoryUseCase := appCategory.NewReadCategoryUseCase(readCategoryService)
	writeCategoryUseCase := appCategory.NewWriteCategoryUseCase(writeCategoryService)

	// Handler´s
	readQuotationHandler := quotationController.NewReadQuotationUseCase(readQuotationUseCase)
	writeQuotationHandler := quotationController.NewWriteQuotationHandler(writeQuotationUseCase)
	readProductHandler := controller.NewReadProductHandler(readProductUseCase)
	writeProductHandler := controller.NewWriteProductHandler(writeProductUseCase)
	readCategoryHandler := catController.NewReadCategoryHandler(writeCategoryUseCase, readCategegoryUseCase)
	writeCategoryHandler := catController.NewWriteCategoryHandler(writeCategoryUseCase)

	// Routes
	readQuotationHandler.ReadQuotationRouter(r)
	writeQuotationHandler.WriteQuotationRouter(r)
	readProductHandler.ReadProductRouter(r)
	writeProductHandler.WriteProductRouter(r)
	readCategoryHandler.CategoryRouter(r)
	writeCategoryHandler.WriteCategoryRouter(r)
}
