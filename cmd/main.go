package main

import (
	"context"

	"github.com/Andressep/QuoteMaker/internal/config"
	"github.com/Andressep/QuoteMaker/internal/infrastructure/db"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			config.New,
			db.New,
		),
		fx.Invoke(),
	)

	app.Run()
}
