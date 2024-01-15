package repository

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type sqlCustomerRepository struct {
	db *sqlx.DB
}

const saveCustomerQuery = `
INSERT INTO customer (name, rut, address, phone, email)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, rut, address, phone, email;
`

func (r *sqlCustomerRepository) SaveCustomer(ctx context.Context, args domain.Customer) (domain.Customer, error) {
	row := r.db.QueryRowContext(ctx, saveCustomerQuery, args.Name, args.Rut, args.Address, args.Phone, args.Email)
	var i domain.Customer

	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Rut,
		&i.Address,
		&i.Phone,
		&i.Email,
	)

	return i, err
}

const getCustomerByIDQuery = `
SELECT id, name, rut, address, phone, email
FROM customer
WHERE id = $1;
`

// GetCustomerByID implements ports.CustomerRepository.
func (r *sqlCustomerRepository) GetCustomerByID(ctx context.Context, id int) (*domain.Customer, error) {
	customer := &domain.Customer{}
	err := r.db.GetContext(ctx, customer, getCustomerByIDQuery)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func NewCustomerRepository(db *sqlx.DB) ports.CustomerRepository {
	return &sqlCustomerRepository{
		db: db,
	}
}
