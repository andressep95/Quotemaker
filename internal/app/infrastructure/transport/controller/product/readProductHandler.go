package controller

import (
	"net/http"
	"strconv"

	application "github.com/Andressep/QuoteMaker/internal/app/application/product"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/product"
	"github.com/gin-gonic/gin"
)

// ProductRouter registra las rutas para productos en Gin.
func (r *readProductHandler) ReadProductRouter(g *gin.Engine) {
	g.GET("/products", r.ListProductsByNameHandler)
	g.GET("/products/:id", r.GetProductByIDHandler)
	g.GET("/products/category/:categoryName", r.ListProductsByCategoryHandler)
}

type readProductHandler struct {
	readProductUseCase *application.ReadProductUseCase
}

func NewReadProductHandler(readProductUseCase *application.ReadProductUseCase) *readProductHandler {
	return &readProductHandler{
		readProductUseCase: readProductUseCase,
	}
}

func (r *readProductHandler) ListProductsByNameHandler(c *gin.Context) {
	name := c.Query("name")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit value"})
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset value"})
		return
	}

	request := dto.ListProductsRequest{
		Name:   name,
		Limit:  limit,
		Offset: offset,
	}

	response, err := r.readProductUseCase.ListProductByName(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *readProductHandler) ListProductsByCategoryHandler(c *gin.Context) {
	categoryName := c.Param("categoryName")

	request := dto.ListProductByCategoryRequest{
		CategoryName: categoryName,
	}

	response, err := r.readProductUseCase.ListProductByCategory(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *readProductHandler) GetProductByIDHandler(c *gin.Context) {
	id := c.Param("id")
	// No conversion needed for UUID.

	request := dto.GetProductByIDRequest{
		ID: id,
	}

	response, err := r.readProductUseCase.GetProductByID(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
