package controller

import (
	"net/http"
	"strconv"

	application "github.com/Andressep/QuoteMaker/internal/app/application/product"
	"github.com/gin-gonic/gin"
)

// ProductRouter registra las rutas para productos en Gin.
func (pc *ProductController) ProductRouter(r *gin.Engine) {
	// Rutas para productos
	r.POST("/products", pc.CreateProductHandler)
	r.GET("/products", pc.ListProductsByNameHandler)
	r.GET("/products/category/:categoryName", pc.ListProductsByCategoryHandler)
	r.GET("/products/:id", pc.GetProductByIDHandler)
	r.PUT("/products/:id", pc.UpdateProductHandler)
	r.DELETE("/products/:id", pc.DeleteProductHandler)

}

type ProductController struct {
	createProductUseCase *application.CreateProduct
	listProductUseCase   *application.ListProduct
}

func NewProductController(createProductUseCase *application.CreateProduct, listProductUseCase *application.ListProduct) *ProductController {
	return &ProductController{
		createProductUseCase: createProductUseCase,
		listProductUseCase:   listProductUseCase,
	}
}

// CreateProductHandler maneja las solicitudes POST para crear nuevos productos.
func (pc *ProductController) CreateProductHandler(c *gin.Context) {
	var req application.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := pc.createProductUseCase.RegisterProduct(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// ListProductsByNameHandler maneja las solicitudes GET para listar productos por nombre.
func (pc *ProductController) ListProductsByNameHandler(c *gin.Context) {
	name := c.Query("name")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	request := application.ListProductsRequest{
		Name:   name,
		Limit:  limit,
		Offset: offset,
	}

	response, err := pc.listProductUseCase.ListProductByName(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListProductsByCategoryHandler maneja las solicitudes GET para listar productos por categoría.
func (pc *ProductController) ListProductsByCategoryHandler(c *gin.Context) {
	categoryName := c.Param("categoryName")

	request := application.ListProductByCategoryRequest{
		CategoryName: categoryName,
	}

	response, err := pc.listProductUseCase.ListProductByCategory(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetProductByIDHandler maneja las solicitudes GET para obtener un producto por su ID.
func (pc *ProductController) GetProductByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	request := application.GetProductByIDRequest{
		ID: id,
	}

	response, err := pc.listProductUseCase.GetProductByID(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (pc *ProductController) UpdateProductHandler(c *gin.Context) {
	var req application.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := pc.createProductUseCase.ModifyProduct(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteProductHandler maneja las solicitudes DELETE para eliminar un producto.
func (pc *ProductController) DeleteProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	request := application.DeleteProductRequest{
		ID: id,
	}

	_, err = pc.createProductUseCase.DeleteProduct(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
