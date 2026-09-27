package manager

import (
	"context"
	"time"
)

type (
	Node interface {
		Name() string
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}

	Manager struct {
		nodes  []Node
		logger Logger
	}
)

func New(logger Logger) *Manager {
	return &Manager{
		logger: logger,
	}
}

func (m *Manager) AddNode(node Node) {
	m.nodes = append(m.nodes, node)
}

func (m *Manager) Start(ctx context.Context) error {
	m.logger.Info(ctx, "manager: starting nodes", "count", len(m.nodes))
	for _, node := range m.nodes {
		err := func() error {
			m.logger.Info(ctx, "manager: starting node", "node", node.Name())
			return node.Start(ctx)
		}()

		if err != nil {
			m.logger.Error(ctx, err, "manager: failed to start node", "node", node.Name())
			return err
		}
	}
	m.logger.Info(ctx, "manager: all nodes started successfully")

	return nil
}

func (m *Manager) Stop(ctx context.Context) error {
	m.logger.Info(ctx, "manager: stopping nodes", "count", len(m.nodes))

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer stopCancel()

	for _, node := range m.nodes {
		err := func() error {
			m.logger.Info(ctx, "manager: stopping node", "node", node.Name())
			return node.Stop(stopCtx)
		}()

		if err != nil {
			m.logger.Error(ctx, err, "manager: failed to stop node", "node", node.Name())
			return err
		}
	}
	m.logger.Info(ctx, "manager: all nodes stopped successfully")

	return nil
}
