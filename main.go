package main

import (
	"context"

	"github.com/TamasGorgics/gomag/internal/boot"
	"github.com/TamasGorgics/gomag/pkg/logx"
)

func main() {
	app := boot.NewApp("credit-backend", boot.NewConfig())
	logx.Info(context.Background(), "Starting credit backend")

	app.SQLite()
	app.HTTPWorker()

	if err := app.Run(); err != nil {
		logx.Fatal(context.Background(), err, "Failed to run credit backend")
	}

	logx.Info(context.Background(), "Credit backend done")
}
