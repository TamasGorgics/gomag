package boot

import (
	"github.com/TamasGorgics/gomag/pkg/service/database"
)

func (a *App) SQLite() *database.SQLite {
	return database.NewSQLite(a.Service, a.config.SQLiteDSN())
}

func (a *App) PostgreSQL() *database.PostgreSQL {
	return database.NewPostgreSQL(a.Service, a.config.PostgreSQLDSN())
}
