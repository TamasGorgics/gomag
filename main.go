package main

import (
	"log"
	"time"

	"net/http"

	"github.com/TamasGorgics/gomag/internal/boot"
	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/TamasGorgics/gomag/pkg/service/httpworker"
)

func main() {
	s := boot.NewApp("credit-backend")

	sqlite := s.SQLite("file:./test.db?mode=memory&cache=shared")

	hs := func() *http.Server {
		return container.RegisterNamed(s.Container(), "http-server", func() *http.Server {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				res, err := sqlite.DB.ExecContext(r.Context(), "SELECT 1")
				if err != nil {
					log.Fatalf("gomag: %v", err)
				}
				ra, _ := res.RowsAffected()
				log.Printf("gomag: %v", ra)
				for i := range 5 {
					time.Sleep(1 * time.Second)
					log.Printf("http-server: %d", i)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Hello, World!"))
			})
			return &http.Server{
				Addr:    ":8080",
				Handler: mux,
			}
		})
	}()

	s.Manage(httpworker.NewWorker(hs))

	if err := s.Run(); err != nil {
		log.Fatalf("gomag: %v", err)
	}
}
