package boot

import "github.com/TamasGorgics/gomag/pkg/service"

type App struct {
	*service.Service
	config Config
}

func NewApp(name string, config Config) *App {
	return &App{
		Service: service.New(name),
		config:  config,
	}
}
