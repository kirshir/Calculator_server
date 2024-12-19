package main

import (
	"github.com/kirshir/Calculator_server/internal/application"
)

func main() {
	app := application.New()
	app.Run()
}
