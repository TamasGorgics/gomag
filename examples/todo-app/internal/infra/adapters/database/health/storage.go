package health

import (
	"context"
	"database/sql"
)

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) GetHealth(ctx context.Context, tx *sql.Tx) (bool, error) {
	row := tx.QueryRowContext(ctx, "SELECT 1")
	var result int
	err := row.Scan(&result)
	if err != nil {
		return false, err
	}
	return result == 1, nil
}
