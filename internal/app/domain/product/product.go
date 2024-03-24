package domain

type Product struct {
	ID          int     `db:"id"`
	CategoryID  int     `db:"category_id"`
	Code        string  `db:"code"`
	Description string  `db:"name"`
	Price       float64 `db:"price"`
	Weight      float64 `db:"weight"`
	Length      float64 `db:"length"`
	IsAvailable bool    `db:"is_available"`
}
