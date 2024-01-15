package db

import (
	"context"
	"fmt"

	"github.com/Andressep/QuoteMaker/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(ctx context.Context, s *config.Config) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)

	return sqlx.ConnectContext(ctx, "postgres", connectionString)
}
