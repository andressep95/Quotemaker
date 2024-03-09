package controller

import (
	"net/http"
	"strconv"

	application "github.com/Andressep/QuoteMaker/internal/app/application/category"
	"github.com/gin-gonic/gin"
)

// CategoryRouter registra las rutas para categorías en Gin.
func (cc *CategoryController) CategoryRouter(r *gin.Engine) {
	// Rutas para categorías
	r.POST("/category", cc.CreateCategoryHandler)
	r.GET("/category", cc.ListCategoryHandler)
}

type CategoryController struct {
	createCategoryUseCase *application.CreateCategory
	listCategoryUseCase   *application.ListCategory
}

func NewCategoryController(createCategoryUseCase *application.CreateCategory, listCategoryUseCase *application.ListCategory) *CategoryController {
	return &CategoryController{
		createCategoryUseCase: createCategoryUseCase,
		listCategoryUseCase:   listCategoryUseCase,
	}
}

// CreateCategoryHandler maneja las solicitudes POST para crear nuevas categorías.
func (cc *CategoryController) CreateCategoryHandler(c *gin.Context) {
	var req application.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := cc.createCategoryUseCase.RegisterCategory(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// ListCategoryHandler maneja las solicitudes GET para listar categorías.
func (cc *CategoryController) ListCategoryHandler(c *gin.Context) {
	// Obtener los parámetros de la solicitud
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	request := application.ListCategorysRequest{
		Limit:  limit,
		Offset: offset,
	}

	response, err := cc.listCategoryUseCase.ListCategorys(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
