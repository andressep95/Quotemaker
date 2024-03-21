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
			return // No establece dbInstance y devuelve el error real
		}
		// Opcional: Verifica la conexión
		if err := dbInstance.Ping(); err != nil {
			dbInstance.Close() // Cierra la conexión
			dbInstance = nil   // Establece dbInstance a nil
			return             // Devuelve el error real
		}
	})
	if err != nil {
		return nil, err // Si hay un error, devuelve nil para la instancia de la base de datos y el error real
	}
	return dbInstance, nil // Si no hay error, devuelve la instancia de la base de datos y nil para el error
}
