package boot

import "github.com/TamasGorgics/gomag/pkg/service"

type App struct {
	*service.Service
}

func NewApp(name string) *App {
	return &App{
		Service: service.New(name),
	}
}
