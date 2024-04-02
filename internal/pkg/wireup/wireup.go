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

	"github.com/gin-gonic/gin"
)

func SetupAppControllers(r *gin.Engine, db *sql.DB) {
	// Repository´s
	//readQuotationRepo := quotationRepo.NewReadQuotationRepository(db)
	writeQuotationRepo := quotationRepo.NewWriteQuotationRepository(db)
	readProductRepo := persistence.NewReadProductRepository(db)
	writeProductRepo := persistence.NewWriteProductRepository(db)
	readCategoryRepo := categoryRepo.NewReadCategoryRepository(db)
	writeCategoryRepo := categoryRepo.NewWriteCategoryRepository(db)

	// Services´s
	writeQuotationService := quotationServ.NewWriteQuotationService(writeQuotationRepo, writeProductRepo)
	readProductService := domain.NewReadProductService(readProductRepo, readCategoryRepo)
	writeProductService := domain.NewWriteProductService(writeProductRepo, writeCategoryRepo, readCategoryRepo)
	readCategoryService := categoryServ.NewReadCategoryService(readCategoryRepo)
	writeCategoryService := categoryServ.NewWriteCategoryService(writeCategoryRepo)

	// Usecase´s
	writeQuotationUseCase := appQuotation.NewWriteQuotationUseCase(writeQuotationService)
	readProductUseCase := application.NewReadProductUseCase(readProductService)
	writeProductUseCase := application.NewWriteProductUseCase(writeProductService, readProductService)
	readCategegoryUseCase := appCategory.NewReadCategoryUseCase(readCategoryService)
	writeCategoryUseCase := appCategory.NewWriteCategoryUseCase(writeCategoryService)

	// Handler´s
	writeQuotationHandler := quotationController.NewWriteQuotationHandler(writeQuotationUseCase)
	readProductHandler := controller.NewReadProductHandler(readProductUseCase)
	writeProductHandler := controller.NewWriteProductHandler(writeProductUseCase)
	readCategoryHandler := catController.NewReadCategoryHandler(writeCategoryUseCase, readCategegoryUseCase)
	writeCategoryHandler := catController.NewWriteCategoryHandler(writeCategoryUseCase)

	// Routes
	writeQuotationHandler.WriteQuotationRouter(r)
	readProductHandler.ReadProductRouter(r)
	writeProductHandler.WriteProductRouter(r)
	readCategoryHandler.CategoryRouter(r)
	writeCategoryHandler.WriteCategoryRouter(r)
}
