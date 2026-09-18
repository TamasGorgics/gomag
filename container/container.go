package container

import "reflect"

type (
	key struct {
		typ  reflect.Type
		name string
	}

	Container struct {
		components map[key]any
	}
)

func New() *Container {
	return &Container{
		components: make(map[key]any),
	}
}

func Register[T any](c *Container, constructor func() T) T {
	return maybeRegister(c, reflect.TypeOf((*T)(nil)).Name(), constructor)
}

func RegisterNamed[T any](c *Container, name string, constructor func() T) T {
	return maybeRegister(c, name, constructor)
}

func maybeRegister[T any](c *Container, name string, constructor func() T) T {
	k := key{typ: reflect.TypeOf((*T)(nil)), name: name}
	if comp, ok := c.components[k]; ok {
		return comp.(T)
	}

	comp := constructor()
	c.components[k] = comp
	return comp
}
