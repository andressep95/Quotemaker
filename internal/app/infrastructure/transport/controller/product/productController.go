package controller

import (
	"net/http"
	"strconv"

	application "github.com/Andressep/QuoteMaker/internal/app/application/product"
	"github.com/labstack/echo"
)

func (pc *ProductController) ProductRouter(e *echo.Echo) {
	// Registra las rutas para productos aquí
	e.POST("/products", pc.CreateProductHandler)
	e.GET("/products", pc.ListProductsHandler)

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
func (pc *ProductController) CreateProductHandler(c echo.Context) error {
	req := new(application.CreateProductRequest)
	if err := c.Bind(req); err != nil {
		// Echo automáticamente responde con un error 400 si el bind falla.
		return err
	}

	resp, err := pc.createProductUseCase.RegisterProduct(c.Request().Context(), req)
	if err != nil {
		// Aquí puedes manejar diferentes tipos de errores y responder adecuadamente.
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

// ListProductsHandler maneja las solicitudes para listar productos.
func (pc *ProductController) ListProductsHandler(c echo.Context) error {
	// Obtener los parámetros de la solicitud
	name := c.QueryParam("name")
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10 // Valor por defecto si no se especifica
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0 // Valor por defecto si no se especifica
	}

	// Crear y ejecutar la solicitud del caso de uso
	request := application.ListProductsRequest{
		Name:   name,
		Limit:  limit,
		Offset: offset,
	}
	response, err := pc.listProductUseCase.ListProductByName(c.Request().Context(), &request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Devolver la respuesta
	return c.JSON(http.StatusOK, response)
}
