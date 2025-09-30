package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/TamasGorgics/gomag/pkg/logx"
	"github.com/TamasGorgics/gomag/pkg/manager"
	"github.com/TamasGorgics/gomag/pkg/service"
)

var _ manager.Node = (*PostgreSQL)(nil)

type PostgreSQL struct {
	dsn string
	db  *sql.DB
}

func NewPostgreSQL(service *service.Service, dsn string) *PostgreSQL {
	return container.RegisterNamed(service.Container(), "postgresql", func() *PostgreSQL {
		postgresql := &PostgreSQL{
			dsn: dsn,
		}
		service.Manage(postgresql)

		return postgresql
	})
}

func (p *PostgreSQL) Name() string {
	return "postgresql"
}

func (p *PostgreSQL) Start(ctx context.Context) error {
	db, err := sql.Open("pgx", p.dsn)
	if err != nil {
		return fmt.Errorf("new pool: %w", err)
	}

	p.db = db

	// TODO add config for tuning
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		logx.Fatal(ctx, err, "Unable to connect to database")
	}

	logx.Info(ctx, "Successfully connected to the database!")

	return nil
}

func (p *PostgreSQL) Stop(_ context.Context) error {
	if p.db != nil {
		p.db.Close()
	}

	return nil
}
