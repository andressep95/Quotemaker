package domain

type Seller struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
