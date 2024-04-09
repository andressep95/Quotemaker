package controller

import (
	"net/http"

	application "github.com/Andressep/QuoteMaker/internal/app/application/product"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/product"
	"github.com/gin-gonic/gin"
)

func (r *writeProductHandler) WriteProductRouter(g *gin.Engine) {
	g.POST("/products", r.CreateProductHandler)
	g.PUT("/products/:id", r.UpdateProductHandler)
	g.DELETE("/products/:id", r.DeleteProductHandler)
}

type writeProductHandler struct {
	writeProductUseCase *application.WriteProductUseCase
}

func NewWriteProductHandler(writeProductUseCase *application.WriteProductUseCase) *writeProductHandler {
	return &writeProductHandler{
		writeProductUseCase: writeProductUseCase,
	}
}

func (w *writeProductHandler) CreateProductHandler(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := w.writeProductUseCase.RegisterProduct(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (w *writeProductHandler) UpdateProductHandler(c *gin.Context) {
	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := w.writeProductUseCase.ModifyProduct(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (w *writeProductHandler) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")

	request := dto.DeleteProductRequest{
		ID: id,
	}

	_, err := w.writeProductUseCase.DeleteProduct(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
