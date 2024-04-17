package dto

import (
	"time"

	dto "github.com/Andressep/QuoteMaker/internal/app/dto/product"
)

// Para la creación de una cotización, incluyendo productos asociados.
type CreateQuotationRequest struct {
	IsPurchased bool                 `json:"is_purchased"`
	Products    []QuoteProductDetail `json:"products"`
	TotalPrice  float64              `json:"total_price"`
}

// Respuesta tras crear una cotización, reflejando los principales campos de interés.
type CreateQuotationResponse struct {
	ID          string               `json:"id"`
	CreatedAt   time.Time            `json:"created_at"`
	TotalPrice  float64              `json:"total_price"`
	IsPurchased bool                 `json:"is_purchased"`
	IsDelivered bool                 `json:"is_delivered"`
	Products    []QuoteProductDetail `json:"products"`
}

// Para actualizar información de la cotización, permitiendo modificar ciertos campos.
type UpdateQuotationRequest struct {
	ID          string     `json:"id"`
	TotalPrice  *float64   `json:"total_price,omitempty"`
	IsPurchased *bool      `json:"is_purchased,omitempty"`
	IsDelivered *bool      `json:"is_delivered,omitempty"`
	PurchasedAt *time.Time `json:"purchased_at,omitempty"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	// Considerar si se permite actualizar los productos en esta solicitud.
}

type QuoteProductDetail struct {
	Product  dto.ProductDTO `json:"product"`
	Quantity int            `json:"quantity"`
}
