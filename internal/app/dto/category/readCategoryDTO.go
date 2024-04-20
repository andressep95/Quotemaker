package dto

type ListCategorysRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
type ListCategorysResponse struct {
	Category []CategoryDTO `json:"category"`
}
type CategoryDTO struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name"`
}
type GetCategoryRequest struct {
	ID string `json:"id"`
}

type GetCategoryResponse struct {
	Category CategoryDTO `json:"category"`
}
type GetCategoryByNameRequest struct {
	Name string `json:"category_name"`
}
