package health

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/TamasGorgics/gomag/internal/infra/adapters/database/health"
	"github.com/TamasGorgics/gomag/pkg/tx"
)

type HealthController struct {
	db            *sql.DB
	healthStorage *health.Storage
}

func NewHealthController(db *sql.DB, healthStorage *health.Storage) *HealthController {
	return &HealthController{db: db, healthStorage: healthStorage}
}

func (h *HealthController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tx.Exec(
		r.Context(),
		h.db,
		sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  true,
		},
		func(ctx context.Context, tx *sql.Tx) error {
			live, err := h.healthStorage.GetHealth(ctx, tx)
			if err != nil {
				return err
			}
			if !live {
				w.WriteHeader(http.StatusServiceUnavailable)
				return nil
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return nil
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
