package boot

import (
	"github.com/TamasGorgics/gomag/internal/infra/adapters/database/health"
	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/TamasGorgics/gomag/pkg/service/database"
)

func (a *App) SQLite() *database.SQLite {
	return database.NewSQLite(a.Service, a.config.SQLiteDSN())
}

func (a *App) PostgreSQL() *database.PostgreSQL {
	return database.NewPostgreSQL(a.Service, a.config.PostgreSQLDSN())
}

func (a *App) HealthStorage() *health.Storage {
	return container.RegisterNamed(a.Container(), "health-storage", health.NewStorage)
}
