package boot

import (
	"net/http"

	"github.com/TamasGorgics/gomag/container"
	"github.com/TamasGorgics/gomag/examples/todo-app/internal/infra/controllers"
	"github.com/TamasGorgics/gomag/middleware"
	"github.com/TamasGorgics/gomag/service/httpworker"
)

func (a *App) HTTPWorker() *httpworker.HttpWorker {
	return container.RegisterNamed(a.Container(), "http-server", func() *httpworker.HttpWorker {
		mux := http.NewServeMux()
		mux.Handle("GET /health", a.HealthController())
		mux.Handle("/tasks", a.TaskController())
		mux.Handle("/tasks/{id}", a.TaskController())
		srv := &http.Server{
			Addr:    ":8080",
			Handler: middleware.RequestID(middleware.Logging(a.Service.Logger(), mux)),
		}
		return httpworker.New(a.Service, srv)
	})
}

func (a *App) HealthController() *controllers.HealthController {
	return container.RegisterNamed(a.Container(), "health-controller", func() *controllers.HealthController {
		return controllers.NewHealthController(a.SQLite().DB, a.HealthStorage())
	})
}

func (a *App) TaskController() *controllers.TaskController {
	return container.RegisterNamed(a.Container(), "task-controller", func() *controllers.TaskController {
		return controllers.NewTaskController(a.TaskService(), a.SQLite().DB)
	})
}
