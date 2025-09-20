package main

import (
	"log"

	"github.com/TamasGorgics/gomag/internal/boot"
)

func main() {
	app := boot.NewApp("credit-backend", boot.NewConfig())

	app.SQLite()
	app.HTTPWorker()

	if err := app.Run(); err != nil {
		log.Fatalf("gomag: %v", err)
	}
}
