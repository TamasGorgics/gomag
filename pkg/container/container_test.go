package container_test

import (
	"testing"

	"github.com/TamasGorgics/gomag/pkg/container"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestService struct {
	container *container.Container
}

func NewTestService() *TestService {
	return &TestService{container: container.New()}
}

type UserServiceComponent struct{}

func (ts *TestService) GetComponentA() *UserServiceComponent {
	return container.Register(ts.container, func() *UserServiceComponent {
		return &UserServiceComponent{}
	})
}

type HTTPServerComponent struct {
	userService *UserServiceComponent
}

func (ts *TestService) GetComponentB() *HTTPServerComponent {
	return container.Register(ts.container, func() *HTTPServerComponent {
		return &HTTPServerComponent{
			userService: ts.GetComponentA(),
		}
	})
}

func TestRegister(t *testing.T) {
	service := NewTestService()

	userService := service.GetComponentA()
	require.NotNil(t, userService)

	httpServer := service.GetComponentB()
	require.NotNil(t, httpServer)
	require.NotNil(t, httpServer.userService)
}

type CounterComponent struct {
	count int
}

func (ts *TestService) GetCounterA() *CounterComponent {
	return container.RegisterNamed(ts.container, "counterA", func() *CounterComponent {
		return &CounterComponent{count: 0}
	})
}

func (ts *TestService) GetCounterB() *CounterComponent {
	return container.RegisterNamed(ts.container, "counterB", func() *CounterComponent {
		return &CounterComponent{count: 1}
	})
}

func TestRegisterNamed(t *testing.T) {
	service := NewTestService()

	counterA := service.GetCounterA()
	require.NotNil(t, counterA)
	require.Equal(t, 0, counterA.count)

	counterB := service.GetCounterB()
	require.NotNil(t, counterB)
	require.Equal(t, 1, counterB.count)

	assert.NotSame(t, counterA, counterB)
	assert.Same(t, counterA, service.GetCounterA())
	assert.Same(t, counterB, service.GetCounterB())
}
