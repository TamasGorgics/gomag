package httpworker

import (
	"context"
	"log"
	"net/http"

	"github.com/TamasGorgics/gomag/pkg/manager"
)

var _ manager.Node = (*HttpWorker)(nil)

type HttpWorker struct {
	server *http.Server
}

func NewWorker(server *http.Server) *HttpWorker {
	return &HttpWorker{
		server: server,
	}
}

func (w *HttpWorker) Name() string {
	return "http-worker"
}

func (w *HttpWorker) Start(_ context.Context) error {
	go func() {
		if err := w.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("http-worker: %v", err)
		}
	}()
	return nil
}

func (w *HttpWorker) Stop(ctx context.Context) error {
	return w.server.Shutdown(ctx)
}
