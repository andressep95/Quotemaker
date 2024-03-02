package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDBConnection(driver, source string) (*sql.DB, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	// Opcional: Verifica la conexión
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
