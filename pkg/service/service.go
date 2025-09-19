package service

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/TamasGorgics/gomag/pkg/manager"
)

type Service struct {
	name      string
	container *container.Container
	manager   *manager.Manager
}

func New(name string) *Service {
	return &Service{
		name:      name,
		container: container.New(),
		manager:   manager.New(),
	}
}

func (s *Service) Name() string {
	return s.name
}

func (s *Service) Container() *container.Container {
	return s.container
}

func (s *Service) Manage(w manager.Node) {
	s.manager.AddNode(w)
}

func (s *Service) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := s.manager.Start(ctx); err != nil {
		return err
	}

	<-ctx.Done()
	log.Printf("service: %s received shutdown signal", s.name)

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopCancel()
	if err := s.manager.Stop(stopCtx); err != nil {
		return err
	}

	return nil
}
