package dto

type GetProductByIDRequest struct {
	ID string `json:"id"`
}
type GetProductByIDResponse struct {
	ID           string  `json:"id"`
	CategoryName string  `json:"category_name"`
	Code         string  `json:"code"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Weight       float64 `json:"weight"`
	Length       float64 `json:"length"`
	IsAvailable  bool    `json:"is_available"`
}

// ListProductsRequest define los datos de entrada para listar productos.
type ListProductsRequest struct {
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

// ListProductsResponse define los datos de salida tras listar productos.
type ListProductsResponse struct {
	Products []ProductDTO `json:"products"`
}

type ListProductByCategoryRequest struct {
	CategoryName string `json:"category_name"`
}

type ProductDTO struct {
	ID           string  `json:"id"`
	CategoryName string  `json:"category_name"`
	Code         string  `json:"code"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Weight       float64 `json:"weight"`
	Length       float64 `json:"length"`
	IsAvailable  bool    `json:"is_available"`
}
