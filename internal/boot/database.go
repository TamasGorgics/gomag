package boot

import (
	"github.com/TamasGorgics/gomag/pkg/service/database"
)

func (a *App) SQLite(dsn string) *database.SQLite {
	return database.NewSQLite(a.Service, dsn)
}

func (a *App) PostgreSQL(dsn string) *database.PostgreSQL {
	return database.NewPostgreSQL(a.Service, dsn)
}
