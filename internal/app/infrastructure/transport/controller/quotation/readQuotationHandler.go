package controller

import (
	"net/http"
	"strconv"

	application "github.com/Andressep/QuoteMaker/internal/app/application/quotation"
	"github.com/gin-gonic/gin"
)

func (r *readQuotationHandler) ReadQuotationRouter(g *gin.Engine) {
	g.GET("/quotations", r.ListQuotationsHandler)
}

func (r *readQuotationHandler) ListQuotationsHandler(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
		return
	}

	// Crea el request de listado de cotizaciones.
	request := application.ListQuotationsRequest{
		Limit:  limit,
		Offset: offset,
	}

	// Llama al caso de uso de listado.
	resp, err := r.readQuotationUseCase.ListQuotations(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

type readQuotationHandler struct {
	readQuotationUseCase *application.ReadQuotationUseCase
}

func NewReadQuotationUseCase(readQuotationUseCase *application.ReadQuotationUseCase) *readQuotationHandler {
	return &readQuotationHandler{
		readQuotationUseCase: readQuotationUseCase,
	}
}
