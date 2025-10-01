package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/TamasGorgics/gomag/pkg/logx"
	"github.com/TamasGorgics/gomag/pkg/manager"
)

type Service struct {
	name      string
	container *container.Container
	manager   *manager.Manager
	logger    logx.Logger
}

func New(name string) *Service {
	logger := logx.InitDefaultLogger()

	return &Service{
		name:      name,
		container: container.New(),
		manager:   manager.New(logger),
		logger:    logger,
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
	s.logger.Info(ctx, "service: received shutdown signal", "name", s.name)

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopCancel()
	if err := s.manager.Stop(stopCtx); err != nil {
		return err
	}

	return nil
}

func (s *Service) Logger() logx.Logger {
	return s.logger
}
