package domain

type Category struct {
	ID           string `db:"id"`
	CategoryName string `db:"category_name"`
}
