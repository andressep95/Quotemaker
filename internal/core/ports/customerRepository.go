package ports

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
)

type CustomerRepository interface {
	SaveCustomer(ctx context.Context, args domain.Customer) (domain.Customer, error)
	GetCustomerByID(ctx context.Context, id int) (*domain.Customer, error)
}
