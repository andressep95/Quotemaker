package dto

type CreateProductRequest struct {
	CategoryID  string  `json:"category_id"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"`
	Length      float64 `json:"length"`
	IsAvailable bool    `json:"is_available"`
}
type UpdateProductRequest struct {
	ID          string  `json:"id"`
	CategoryID  string  `json:"category_id,omitempty"`
	Code        string  `json:"code,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
	Length      float64 `json:"length,omitempty"`
	IsAvailable bool    `json:"is_available,omitempty"`
}

// CreateProductResponse define los datos de salida tras crear un producto.
type CreateProductResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id"`
}

type DeleteProductRequest struct {
	ID string `json:"id"`
}

type DeleteProductResponse struct {
	Message string `json:"message"`
}
