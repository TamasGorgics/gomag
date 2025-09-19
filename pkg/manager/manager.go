package manager

import (
	"context"
	"log"
)

type (
	Node interface {
		Name() string
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}

	Manager struct {
		nodes []Node
	}
)

func New() *Manager {
	return &Manager{}
}

func (m *Manager) AddNode(node Node) {
	m.nodes = append(m.nodes, node)
}

func (m *Manager) Start(ctx context.Context) error {
	log.Printf("manager: starting %d nodes", len(m.nodes))
	for _, node := range m.nodes {
		err := func() error {
			log.Printf("manager: starting node %s", node.Name())
			return node.Start(ctx)
		}()

		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) Stop(ctx context.Context) error {
	log.Printf("manager: stopping %d nodes", len(m.nodes))
	for _, node := range m.nodes {
		err := func() error {
			log.Printf("manager: stopping node %s", node.Name())
			return node.Stop(ctx)
		}()

		if err != nil {
			return err
		}
	}
	return nil
}
