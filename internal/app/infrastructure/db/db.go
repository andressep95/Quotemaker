package db

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var dbInstance *sql.DB
var once sync.Once
var err error

func DBConnection(driver, source string) (*sql.DB, error) {
	once.Do(func() {
		dbInstance, err = sql.Open(driver, source)
		if err != nil {
			return
		}
		// Opcional: Verifica la conexión
		if err := dbInstance.Ping(); err != nil {
			return
		}
	})
	return dbInstance, nil
}
