package ports

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
)

type CustomerRepository interface {
	SaveCustomer(ctx context.Context, args domain.Customer) (domain.Customer, error)
	UpdateCustomer(ctx context.Context, args domain.Customer) error
	GetCustomerByID(ctx context.Context, id int) (*domain.Customer, error)
	ListCustomers(ctx context.Context, limit, offset int) ([]domain.Customer, error)
	DeleteCustomer(ctx context.Context, id int) error
}
