package dto

import "time"

type ListQuotationsRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListQuotationsResponse struct {
	Quotations []QuotationDTO `json:"quotations"`
}

type QuotationDTO struct {
	ID          string               `json:"id"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   *time.Time           `json:"updated_at,omitempty"` // Incluido para manejar la fecha de actualización
	TotalPrice  float64              `json:"total_price"`
	IsPurchased bool                 `json:"is_purchased"`
	PurchasedAt *time.Time           `json:"purchased_at,omitempty"` // Incluido para manejar la fecha de compra
	IsDelivered bool                 `json:"is_delivered"`
	DeliveredAt *time.Time           `json:"delivered_at,omitempty"` // Incluido para manejar la fecha de entrega
	Products    []QuoteProductDetail `json:"products"`
}
