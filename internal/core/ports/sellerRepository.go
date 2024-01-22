package ports

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
)

type SellerRepository interface {
	SaveSeller(ctx context.Context, args domain.Seller) (domain.Seller, error)
	GetSellerByID(ctx context.Context, id int) (*domain.Seller, error)
	ListSellers(ctx context.Context, limit, offset int) ([]domain.Seller, error)
	DeleteSeller(ctx context.Context, id int) error
}
