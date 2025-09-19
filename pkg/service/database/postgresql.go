package database

import (
	"context"
	"fmt"
	"time"

	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/TamasGorgics/gomag/pkg/manager"
	"github.com/TamasGorgics/gomag/pkg/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ manager.Node = (*PostgreSQL)(nil)

type PostgreSQL struct {
	cfg  *pgxpool.Config
	pool *pgxpool.Pool
}

func NewPostgreSQL(service *service.Service, dsn string) *PostgreSQL {
	return container.RegisterNamed(service.Container(), "postgresql", func() *PostgreSQL {
		cfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			panic(err)
		}

		// TODO add config for tuning
		cfg.MaxConns = 10
		cfg.MinConns = 2
		cfg.HealthCheckPeriod = 30 * time.Second
		cfg.MaxConnLifetime = 30 * time.Minute
		cfg.MaxConnIdleTime = 5 * time.Minute
		cfg.ConnConfig.ConnectTimeout = 3 * time.Second

		postgresql := &PostgreSQL{
			cfg: cfg,
		}
		service.Manage(postgresql)

		return postgresql
	})
}

func (p *PostgreSQL) Name() string {
	return "postgresql"
}

func (p *PostgreSQL) Start(ctx context.Context) error {
	pool, err := pgxpool.NewWithConfig(ctx, p.cfg)
	if err != nil {
		return fmt.Errorf("new pool: %w", err)
	}

	p.pool = pool

	return nil
}

func (p *PostgreSQL) Stop(_ context.Context) error {
	if p.pool != nil {
		p.pool.Close()
	}

	return nil
}
