package tx

import (
	"context"
	"database/sql"
)

func Exec(
	ctx context.Context,
	db *sql.DB,
	opts sql.TxOptions,
	fn func(ctx context.Context, tx *sql.Tx) error,
) error {
	tx, err := db.BeginTx(ctx, &opts)
	if err != nil {
		return err
	}

	shouldCommit := false
	defer func() {
		if shouldCommit {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return err
	}

	shouldCommit = true

	return nil
}
