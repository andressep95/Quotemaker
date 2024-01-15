package domain

type QuoteProduct struct {
	ID          int `db:"id"`
	QuotationID int `db:"quotation_id"`
	ProductID   int `db:"product_id"`
	Quantity    int `db:"quantity"`
	DeliveryID  int `db:"delivery_id"`
}
