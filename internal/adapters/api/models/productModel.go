package models

type CreateProductParams struct {
	Name       string  `json:"name"`
	CategoryID int32   `json:"category_id"`
	Length     float32 `json:"length"`
	Price      float64 `json:"price"`
	Weight     float32 `json:"weight"`
	Code       string  `json:"code"`
}
