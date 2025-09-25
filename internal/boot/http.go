package boot

import (
	"net/http"

	"github.com/TamasGorgics/gomag/internal/infra/controllers/health"
	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/TamasGorgics/gomag/pkg/middleware"
	"github.com/TamasGorgics/gomag/pkg/service/httpworker"
)

func (a *App) HTTPWorker() *httpworker.HttpWorker {
	return container.RegisterNamed(a.Container(), "http-server", func() *httpworker.HttpWorker {
		mux := http.NewServeMux()
		mux.Handle("GET /health", a.HealthController())
		srv := &http.Server{
			Addr:    ":8080",
			Handler: middleware.RequestID(mux),
		}
		return httpworker.New(a.Service, srv)
	})
}

func (a *App) HealthController() *health.HealthController {
	return container.RegisterNamed(a.Container(), "health-controller", func() *health.HealthController {
		return health.NewHealthController(a.SQLite().DB, a.HealthStorage())
	})
}
