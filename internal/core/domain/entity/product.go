package domain

type Product struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	CategoryID int     `db:"category_id"`
	Length     float64 `db:"length"`
	Price      float64 `db:"price"`
	Weight     float64 `db:"weight"`
	Code       string  `db:"code"`
}
