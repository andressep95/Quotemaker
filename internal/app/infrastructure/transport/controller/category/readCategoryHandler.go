package controller

import (
	"net/http"
	"strconv"

	application "github.com/Andressep/QuoteMaker/internal/app/application/category"
	"github.com/gin-gonic/gin"
)

// CategoryRouter registra las rutas para categorías en Gin.
func (rc *ReadCategoryHandler) CategoryRouter(r *gin.Engine) {
	r.GET("/category", rc.ListCategoryHandler)
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

	request := application.ListCategorysRequest{
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
