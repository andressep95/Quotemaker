package persistence

import (
	"context"
	"database/sql"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/customer"
)

type sqlCustomerRepository struct {
	db *sql.DB
}

const listCustomerQuery = `
SELECT id, name, rut, address, phone, email
FROM customer
ORDER BY id
LIMIT $1 OFFSET $2;
`

// listCustomers implements ports.CustomerRepository.
func (r *sqlCustomerRepository) ListCustomers(ctx context.Context, limit int, offset int) ([]domain.Customer, error) {
	rows, err := r.db.QueryContext(ctx, listCustomerQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		var i domain.Customer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.Email,
			&i.Phone,
			&i.Phone); err != nil {
			return nil, err
		}
		customers = append(customers, i)
	}

	// Verificar por errores al finalizar la iteración
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
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

const updateCustomerQuery = `
UPDATE customer
SET name = $1, rut = $2, address = $3, phone = $4, email = $5 
where id = $6;
`

// UpdateCustomer implements ports.CustomerRepository.
func (r *sqlCustomerRepository) UpdateCustomer(ctx context.Context, args domain.Customer) error {
	_, err := r.db.ExecContext(ctx, updateCustomerQuery, args.Name, args.Rut, args.Address, args.Phone, args.Email, args.ID)
	return err
}

const getCustomerByIDQuery = `
SELECT id, name, rut, address, phone, email
FROM customer
WHERE id = $1;
`

// GetCustomerByID implements ports.CustomerRepository.
func (r *sqlCustomerRepository) GetCustomerByID(ctx context.Context, id int) (*domain.Customer, error) {
	row := r.db.QueryRowContext(ctx, getCustomerByIDQuery, id)
	var i domain.Customer

	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Rut,
		&i.Address,
		&i.Phone,
		&i.Email,
	)

	return &i, err
}

const deleteCustomerQuery = `
DELETE FROM customer
WHERE id = $1;
`

func (r *sqlCustomerRepository) DeleteCustomer(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, deleteCustomerQuery, id)
	return err
}

func NewCustomerRepository(db *sql.DB) domain.CustomerRepository {
	return &sqlCustomerRepository{
		db: db,
	}
}
