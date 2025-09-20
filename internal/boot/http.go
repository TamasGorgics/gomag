package boot

import (
	"log"
	"net/http"
	"time"

	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/TamasGorgics/gomag/pkg/service/httpworker"
)

func (a *App) HTTPWorker() *httpworker.HttpWorker {
	return container.RegisterNamed(a.Container(), "http-server", func() *httpworker.HttpWorker {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tx, err := a.SQLite().DB.BeginTx(r.Context(), nil)
			if err != nil {
				log.Fatalf("gomag: %v", err)
			}
			defer tx.Rollback()
			live, err := a.HealthStorage().GetHealth(r.Context(), tx)
			if err != nil {
				log.Fatalf("gomag: %v", err)
			}
			if !live {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			for i := range 5 {
				time.Sleep(1 * time.Second)
				log.Printf("http-server: %d", i)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, World!"))
		})
		srv := &http.Server{
			Addr:    ":8080",
			Handler: mux,
		}
		return httpworker.New(a.Service, srv)
	})
}
