package controller

import (
	"net/http"

	application "github.com/Andressep/QuoteMaker/internal/app/application/category"
	"github.com/gin-gonic/gin"
)

type WriteCategoryHandler struct {
	writeCategoryUseCase *application.WriteCategoryUseCase
}

func (rw WriteCategoryHandler) WriteCategoryRouter(r *gin.Engine) {
	r.POST("/category", rw.CreateCategoryHandler)
}

func NewWriteCategoryHandler(writeCategoryUseCase *application.WriteCategoryUseCase) *WriteCategoryHandler {
	return &WriteCategoryHandler{
		writeCategoryUseCase: writeCategoryUseCase,
	}
}

// CreateCategoryHandler maneja las solicitudes POST para crear nuevas categorías.
func (rw *WriteCategoryHandler) CreateCategoryHandler(c *gin.Context) {
	var req application.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := rw.writeCategoryUseCase.RegisterCategory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
