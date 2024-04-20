package controller

import (
	"net/http"

	application "github.com/Andressep/QuoteMaker/internal/app/application/category"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/category"
	"github.com/gin-gonic/gin"
)

type WriteCategoryHandler struct {
	writeCategoryUseCase *application.WriteCategoryUseCase
}

func (rw *WriteCategoryHandler) WriteCategoryRouter(r *gin.Engine) {
	r.POST("/category", rw.CreateCategoryHandler)
	r.PUT("/category", rw.UpdateCategoryHandler)
	r.DELETE("/category/:id", rw.DeleteCategoryHandler)
}

func NewWriteCategoryHandler(writeCategoryUseCase *application.WriteCategoryUseCase) *WriteCategoryHandler {
	return &WriteCategoryHandler{
		writeCategoryUseCase: writeCategoryUseCase,
	}
}

// CreateCategoryHandler maneja las solicitudes POST para crear nuevas categorías.
func (rw *WriteCategoryHandler) CreateCategoryHandler(c *gin.Context) {
	var req dto.CreateCategoryRequest
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

// UpdateCategoryHandler maneja las solicitudes PUT para actualizar categorías existentes.
func (rw *WriteCategoryHandler) UpdateCategoryHandler(c *gin.Context) {
	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := rw.writeCategoryUseCase.UpdateCategory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteCategoryHandler maneja las solicitudes DELETE para eliminar categorías existentes.
func (rw *WriteCategoryHandler) DeleteCategoryHandler(c *gin.Context) {
	categoryID := c.Param("id")
	req := &dto.DeleteCategoryRequest{
		ID: categoryID,
	}
	_, err := rw.writeCategoryUseCase.DeleteCategory(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
