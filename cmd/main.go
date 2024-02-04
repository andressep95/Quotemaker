package main

import "github.com/Andressep/QuoteMaker/internal/setup"

func main() {
	e := setup.InitializeApp()
	// Inicia el servidor
	e.Logger.Fatal(e.Start(":8080"))

}
