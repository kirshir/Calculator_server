package main

import (
	"log"

	"github.com/kirshir/Calculator_server/internal/application"
)

func main() {
	app := application.New()
	// app.Run()
	err := app.RunServer()
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
