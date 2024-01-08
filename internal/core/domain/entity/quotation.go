package domain

import (
	"time"
)

type Quotation struct {
	ID         int32     `db:"id"`
	SellerID   int32     `db:"seller_id"`
	CustomerID int32     `db:"customer_id"`
	Date       time.Time `db:"date"`
	TotalPrice int64     `db:"total_price"`
}
