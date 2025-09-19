package database

import (
	"context"
	"fmt"
	"time"

	"github.com/TamasGorgics/gomag/pkg/manager"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ manager.Node = (*PostgreSQL)(nil)

type PostgreSQL struct {
	cfg  *pgxpool.Config
	pool *pgxpool.Pool
}

func NewPostgreSQL(dsn string) *PostgreSQL {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}

	// Pool tuning – tailor to your workload
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.HealthCheckPeriod = 30 * time.Second
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.ConnConfig.ConnectTimeout = 3 * time.Second // per-connection timeout

	return &PostgreSQL{
		cfg: cfg,
	}
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
