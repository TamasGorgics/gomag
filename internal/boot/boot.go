package boot

import (
	"github.com/TamasGorgics/gomag/pkg/logx"
	"github.com/TamasGorgics/gomag/pkg/service"
)

type App struct {
	*service.Service
	config Config
}

func NewApp(name string, config Config) *App {
	return &App{
		Service: service.New(name, service.WithLogger(logx.InitDefaultLogger())),
		config:  config,
	}
}
