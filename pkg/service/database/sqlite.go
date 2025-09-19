package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/TamasGorgics/gomag/pkg/manager"
	_ "github.com/mattn/go-sqlite3"
)

var _ manager.Node = (*SQLite)(nil)

type SQLite struct {
	DB *sql.DB
}

func NewSQLite(dsn string) *SQLite {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	// Connection pool tuning for SQLite.
	// With WAL enabled (below), many readers + single writer works well.
	db.SetMaxOpenConns(10)                 // allow concurrent readers
	db.SetMaxIdleConns(10)                 // keep some open
	db.SetConnMaxIdleTime(5 * time.Minute) // recycle idle conns
	db.SetConnMaxLifetime(30 * time.Minute)

	// Apply pragmatic defaults. Do this on a short context.
	cfgCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	pragma := []string{
		`PRAGMA journal_mode=WAL;`,   // better concurrency
		`PRAGMA synchronous=NORMAL;`, // durability/perf tradeoff
		`PRAGMA foreign_keys=ON;`,    // enforce FKs
		`PRAGMA busy_timeout=5000;`,  // wait up to 5s if DB is locked
		`PRAGMA cache_size=-20000;`,  // ~20MB cache (negative => KiB)
	}
	for _, q := range pragma {
		if _, err := db.ExecContext(cfgCtx, q); err != nil {
			_ = db.Close()
			panic(fmt.Errorf("apply %q: %w", q, err))
		}
	}

	return &SQLite{
		DB: db,
	}
}

func (s *SQLite) Name() string {
	return "sqlite"
}

func (s *SQLite) Start(_ context.Context) error {
	return nil
}

func (s *SQLite) Stop(_ context.Context) error {
	s.DB.Close()
	return nil
}
