package domain

type Product struct {
	ID         int32   `db:"id"`
	Name       string  `db:"name"`
	CategoryID int32   `db:"category_id"`
	Length     float32 `db:"length"`
	Price      float64 `db:"price"`
	Weight     float64 `db:"weight"`
	Code       string  `db:"code"`
}
