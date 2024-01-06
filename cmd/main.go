package main

import (
	"context"
	"fmt"

	"github.com/Andressep/QuoteMaker/internal/config"
	"github.com/Andressep/QuoteMaker/internal/infraestructure/db"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			config.New,
			db.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM customer")
				if err != nil {
					fmt.Println("all right")
					panic(err)
				}
			},
		),
	)

	app.Run()
}
