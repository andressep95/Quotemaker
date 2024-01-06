package domain

type Category struct {
	ID           int32  `db:"id"`
	CategoryName string `db:"category_name"`
}
