package domain

import (
	"context"
)

type SellerRepository interface {
	SaveSeller(ctx context.Context, args Seller) (Seller, error)
	GetSellerByID(ctx context.Context, id int) (*Seller, error)
	ListSellers(ctx context.Context, limit, offset int) ([]Seller, error)
	DeleteSeller(ctx context.Context, id int) error
	UpdateSeller(ctx context.Context, args Seller) error
}
