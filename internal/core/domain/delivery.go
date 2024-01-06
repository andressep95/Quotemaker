package domain

type Delivery struct {
	ID      int32  `db:"id"`
	Address string `db:"address"`
	Weight  string `db:"weight"`
	Cost    string `db:"cost"`
}
