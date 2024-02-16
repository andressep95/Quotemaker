package domain

type Customer struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Rut     string `db:"rut"`
	Address string `db:"address"`
	Phone   string `db:"phone"`
	Email   string `db:"email"`
}
