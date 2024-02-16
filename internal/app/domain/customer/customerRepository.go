package domain

import (
	"context"
)

type CustomerRepository interface {
	SaveCustomer(ctx context.Context, args Customer) (Customer, error)
	UpdateCustomer(ctx context.Context, args Customer) error
	GetCustomerByID(ctx context.Context, id int) (*Customer, error)
	ListCustomers(ctx context.Context, limit, offset int) ([]Customer, error)
	DeleteCustomer(ctx context.Context, id int) error
}
