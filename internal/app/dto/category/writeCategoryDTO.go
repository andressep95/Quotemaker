package dto

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
}
type CreateCategoryResponse struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name"`
}
type UpdateCategoryRequest struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name"`
}

type UpdateCategoryResponse struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name"`
}

type DeleteCategoryRequest struct {
	ID string `json:"id"`
}

type DeleteCategoryResponse struct {
	ID string `json:"id"`
}
