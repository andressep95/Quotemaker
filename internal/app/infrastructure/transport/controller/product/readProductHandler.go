package controller

import (
	"net/http"
	"strconv"

	application "github.com/Andressep/QuoteMaker/internal/app/application/product"
	"github.com/gin-gonic/gin"
)

// ProductRouter registra las rutas para productos en Gin.
func (r *readProductHanndler) ReadProductRouter(g *gin.Engine) {
	g.GET("/products", r.ListProductsByNameHandler)
	g.GET("/products/:id", r.GetProductByIDHandler)
	g.GET("/products/category/:categoryName", r.ListProductsByCategoryHandler)
}

type readProductHanndler struct {
	readProductUseCase *application.ReadProductUseCase
}

func NewReadProductHandler(readProductUseCase *application.ReadProductUseCase) *readProductHanndler {
	return &readProductHanndler{
		readProductUseCase: readProductUseCase,
	}
}

func (r *readProductHanndler) ListProductsByNameHandler(c *gin.Context) {
	name := c.Query("name")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	request := application.ListProductsRequest{
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

func (r *readProductHanndler) ListProductsByCategoryHandler(c *gin.Context) {
	categoryName := c.Param("categoryName")

	request := application.ListProductByCategoryRequest{
		CategoryName: categoryName,
	}

	response, err := r.readProductUseCase.ListProductByCategory(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (r *readProductHanndler) GetProductByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	request := application.GetProductByIDRequest{
		ID: id,
	}

	response, err := r.readProductUseCase.GetProductByID(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
