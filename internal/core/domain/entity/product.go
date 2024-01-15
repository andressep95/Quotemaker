package domain

type Product struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	CategoryID int     `db:"category_id"`
	Length     int     `db:"length"`
	Price      int     `db:"price"`
	Weight     float64 `db:"weight"`
	Code       string  `db:"code"`
}
