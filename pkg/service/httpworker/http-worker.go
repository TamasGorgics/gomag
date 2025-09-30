package httpworker

import (
	"context"
	"net/http"

	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/TamasGorgics/gomag/pkg/logx"
	"github.com/TamasGorgics/gomag/pkg/manager"
	"github.com/TamasGorgics/gomag/pkg/service"
)

var _ manager.Node = (*HttpWorker)(nil)

type HttpWorker struct {
	server *http.Server
}

func New(service *service.Service, server *http.Server) *HttpWorker {
	return container.RegisterNamed(service.Container(), "http-worker", func() *HttpWorker {
		w := &HttpWorker{
			server: server,
		}
		service.Manage(w)
		return w
	})
}

func (w *HttpWorker) Name() string {
	return "http-worker"
}

func (w *HttpWorker) Start(ctx context.Context) error {
	go func() {
		if err := w.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logx.Error(ctx, err, "http-worker: failed to listen and serve")
		}
	}()
	return nil
}

func (w *HttpWorker) Stop(ctx context.Context) error {
	return w.server.Shutdown(ctx)
}
