package seller

import (
	"context"
	"database/sql"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/seller"
)

type sqlSellerRepository struct {
	db *sql.DB
}

const saveSellerQuery = `
INSERT INTO seller (name)
VALUES ($1)
RETURNING id, name;
`

// SaveSeller implements ports.SellerRepository.
func (r *sqlSellerRepository) SaveSeller(ctx context.Context, args domain.Seller) (domain.Seller, error) {
	row := r.db.QueryRowContext(ctx, saveSellerQuery, args.Name)
	var i domain.Seller

	err := row.Scan(
		&i.ID,
		&i.Name,
	)
	return i, err
}

const updateSellerQuery = `
UPDATE seller
SET name = $1
WHERE id = $2;
`

// UpdateSeller implements ports.SellerRepository.
func (r *sqlSellerRepository) UpdateSeller(ctx context.Context, args domain.Seller) error {
	_, err := r.db.ExecContext(ctx, updateSellerQuery, args.Name, args.ID)
	if err != nil {
		return err
	}
	return nil
}

const getSellerByID = `
SELECT id, name
FROM seller
WHERE id = $1;
`

// GetSellerByID implements ports.SellerRepository.
func (r *sqlSellerRepository) GetSellerByID(ctx context.Context, id int) (*domain.Seller, error) {
	row := r.db.QueryRowContext(ctx, getSellerByID, id)
	var i domain.Seller

	err := row.Scan(
		&i.ID,
		&i.Name,
	)
	return &i, err
}

const deleteSellerQuery = `
DELETE FROM seller
WHERE id = $1;
`

// DeleteSeller implements ports.SellerRepository.
func (r *sqlSellerRepository) DeleteSeller(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, deleteSellerQuery, id)
	return err
}

const listSellersQuery = `
SELECT id, name
FROM seller
ORDER BY id
LIMIT $1 OFFSET $2;
`

// ListSellers implements ports.SellerRepository.
func (r *sqlSellerRepository) ListSellers(ctx context.Context, limit int, offset int) ([]domain.Seller, error) {
	rows, err := r.db.QueryContext(ctx, listSellersQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sellers []domain.Seller
	for rows.Next() {
		var i domain.Seller
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		sellers = append(sellers, i)
	}

	// Verificar por errores al finalizar la iteración
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sellers, nil
}

func NewSellerRepository(db *sql.DB) domain.SellerRepository {
	return &sqlSellerRepository{
		db: db,
	}
}
