package boot

import (
	"github.com/TamasGorgics/gomag/container"
	"github.com/TamasGorgics/gomag/examples/todo-app/internal/app"
)

func (a *App) TaskService() *app.TaskService {
	return container.RegisterNamed(a.Container(), "task-controller", func() *app.TaskService {
		return app.NewTaskService(*a.TaskRepository())
	})
}
