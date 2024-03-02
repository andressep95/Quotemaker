package main

import (
	"log"

	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/config"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/db"
	"github.com/Andressep/QuoteMaker/internal/pkg/wireup"
	"github.com/labstack/echo"
)

func main() {
	// Inicia el servidor
	e := echo.New()

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load the config: ", err)
	}

	// Crea la conexión a la base de datos
	db, err := db.NewDBConnection(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("No se pudo establecer la conexión a la base de datos: %v", err)
	}
	defer db.Close()

	wireup.SetupAppControllers(e, db)
	// Inicia el servidor
	e.Logger.Fatal(e.Start(config.ServerAddress))

}
