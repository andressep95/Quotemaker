package domain

type Delivery struct {
	ID      int     `db:"id"`
	Address string  `db:"address"`
	Weight  float64 `db:"weight"`
	Cost    float64 `db:"cost"`
}
