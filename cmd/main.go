package main

import (
	"log"

	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/config"
	"github.com/Andressep/QuoteMaker/internal/app/infrastructure/db"
	"github.com/Andressep/QuoteMaker/internal/pkg/wireup"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicia el servidor
	r := gin.Default()

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

	wireup.SetupAppControllers(r, db)
	// Inicia el servidor
	r.Run(config.ServerAddress)

}
