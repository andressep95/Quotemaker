package domain

import (
	"time"
)

type Quotation struct {
	ID         int       `db:"id"`
	SellerID   int       `db:"seller_id"`
	CustomerID int       `db:"customer_id"`
	CreatedAt  time.Time `db:"created_at"`
	TotalPrice int       `db:"total_price"`
}
