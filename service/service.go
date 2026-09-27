package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/TamasGorgics/gomag/container"
	"github.com/TamasGorgics/gomag/logx"
	"github.com/TamasGorgics/gomag/manager"
)

type Service struct {
	name      string
	container *container.Container
	manager   *manager.Manager
	logger    logx.Logger
}

func New(name string, options ...ServiceOptions) *Service {
	s := &Service{
		name:      name,
		container: container.New(),
	}

	for _, option := range options {
		option(s)
	}

	if s.logger == nil {
		WithLogger(logx.InitLocalLogger())(s)
	}

	s.manager = manager.New(s.logger)

	return s
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

	if err := s.manager.Stop(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Service) Logger() logx.Logger {
	return s.logger
}
