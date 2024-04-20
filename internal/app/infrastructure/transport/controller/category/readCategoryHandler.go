package controller

import (
	"net/http"
	"strconv"

	application "github.com/Andressep/QuoteMaker/internal/app/application/category"
	dto "github.com/Andressep/QuoteMaker/internal/app/dto/category"
	"github.com/gin-gonic/gin"
)

// CategoryRouter registra las rutas para categorías en Gin.
func (rc *ReadCategoryHandler) CategoryRouter(r *gin.Engine) {
	r.GET("/category", rc.ListCategoryHandler)
	r.GET("/category/:id", rc.GetCategoryByIdHandler)
}

type ReadCategoryHandler struct {
	writeCategoryUseCase *application.WriteCategoryUseCase
	ReadCategoryUseCase  *application.ReadCategoryUseCase
}

func NewReadCategoryHandler(writeCategoryUseCase *application.WriteCategoryUseCase, ReadCategoryUseCase *application.ReadCategoryUseCase) *ReadCategoryHandler {
	return &ReadCategoryHandler{
		writeCategoryUseCase: writeCategoryUseCase,
		ReadCategoryUseCase:  ReadCategoryUseCase,
	}
}

// ListCategoryHandler maneja las solicitudes GET para listar categorías.
func (rc *ReadCategoryHandler) ListCategoryHandler(c *gin.Context) {
	// Obtener los parámetros de la solicitud
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	request := dto.ListCategorysRequest{
		Limit:  limit,
		Offset: offset,
	}

	response, err := rc.ReadCategoryUseCase.ListCategorys(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetCategoryByIdHandler maneja las solicitudes GET para obtener una categoría por su ID.
func (rc *ReadCategoryHandler) GetCategoryByIdHandler(c *gin.Context) {
	// Obtener el ID de la categoría de los parámetros de la solicitud
	categoryID := c.Param("id")

	// Crear la solicitud para obtener la categoría por su ID
	request := dto.GetCategoryRequest{
		ID: categoryID,
	}

	// Llamar al caso de uso para obtener la categoría por su ID
	response, err := rc.ReadCategoryUseCase.GetCategory(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// SearchCategoryByNameHandler maneja las solicitudes GET para buscar una categoría por su nombre.
func (rc *ReadCategoryHandler) SearchCategoryByNameHandler(c *gin.Context) {
	// Obtener el nombre de la categoría de los parámetros de la solicitud
	categoryName := c.Query("name")
	// Crear la solicitud para buscar la categoría por su nombre
	request := dto.GetCategoryByNameRequest{
		Name: categoryName,
	}
	// Llamar al caso de uso para buscar la categoría por su nombre
	response, err := rc.ReadCategoryUseCase.SearchCategoryByName(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
