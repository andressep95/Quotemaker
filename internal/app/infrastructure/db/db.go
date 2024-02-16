package db

import (
	"database/sql"
	"fmt"

	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/config"
	_ "github.com/lib/pq"
)

func NewDBConnection(s *config.DBConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		s.User,
		s.Password,
		s.Host,
		s.Port,
		s.Name,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	// Opcional: Verifica la conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
