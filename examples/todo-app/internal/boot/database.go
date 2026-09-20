package boot

import (
	"github.com/TamasGorgics/gomag/container"
	"github.com/TamasGorgics/gomag/examples/todo-app/internal/infra/adapters/database"
	"github.com/TamasGorgics/gomag/examples/todo-app/internal/infra/adapters/database/health"
	svcDB "github.com/TamasGorgics/gomag/service/database"
)

func (a *App) SQLite() *svcDB.SQLite {
	return svcDB.NewSQLite(a.Service, a.config.SQLiteDSN())
}

func (a *App) PostgreSQL() *svcDB.PostgreSQL {
	return svcDB.NewPostgreSQL(a.Service, a.config.PostgreSQLDSN())
}

func (a *App) HealthStorage() *health.Storage {
	return container.RegisterNamed(a.Container(), "health-storage", health.NewStorage)
}

func (a *App) TaskRepository() *database.TaskRepository {
	return container.RegisterNamed(a.Container(), "task-storage", func() *database.TaskRepository {
		return database.NewTaskRepository(a.SQLite().DB)
	})
}
