package controller

import (
	"net/http"

	application "github.com/Andressep/QuoteMaker/internal/app/application/quotation"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/quotation"
	"github.com/gin-gonic/gin"
)

// WriteQuotationRouter define las rutas para las operaciones de escritura de cotizaciones
func (r *writeQuotationHandler) WriteQuotationRouter(g *gin.Engine) {
	g.POST("/quotations", r.CreateQuotationHandler)

}

// writeQuotationHandler maneja las solicitudes HTTP relacionadas con cotizaciones
type writeQuotationHandler struct {
	writeQuotationUseCase *application.WriteQuotationUseCase
}

// NewWriteQuotationHandler crea una nueva instancia de writeQuotationHandler
func NewWriteQuotationHandler(writeQuotationUseCase *application.WriteQuotationUseCase) *writeQuotationHandler {
	return &writeQuotationHandler{
		writeQuotationUseCase: writeQuotationUseCase,
	}
}

// CreateQuotationHandler maneja la creación de cotizaciones
func (w *writeQuotationHandler) CreateQuotationHandler(c *gin.Context) {
	var req dto.CreateQuotationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := w.writeQuotationUseCase.RegisterQuotation(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
