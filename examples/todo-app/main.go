package main

import (
	"context"

	"github.com/TamasGorgics/gomag/examples/todo-app/internal/boot"
	"github.com/TamasGorgics/gomag/logx"
)

func main() {
	app := boot.NewApp("Todo App", boot.NewConfig())
	logx.Info(context.Background(), "Starting Todo App backend")

	app.SQLite()
	app.HTTPWorker()

	if err := app.Run(); err != nil {
		logx.Fatal(context.Background(), err, "Failed to run Todo App backend")
	}

	logx.Info(context.Background(), "Todo App backend done")
}
