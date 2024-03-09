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
	r.GET("/products/category", pc.ListProductsByCategoryHandler)
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
	categoryName := c.Query("category_name")

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
