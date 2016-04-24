package flux

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestNewDispatcher(t *testing.T) {
	a := &st{}
	b := &st{}
	c := NewDispatcher(a, b)
	assert.Equal(t, 2, len(c.stores))
	assert.Equal(t, a, c.stores[0])
	assert.Equal(t, b, c.stores[1])
}

func TestDispatcher_Register(t *testing.T) {
	a := &st{}
	b := NewDispatcher()
	b.Register(a)
	assert.Equal(t, 1, len(b.stores))
	assert.Equal(t, a, b.stores[0])
}

func TestDispatcher_Dispatch(t *testing.T) {
	a := &st1{}
	b := NewDispatcher(a)
	b.Dispatch("a")
	assert.Equal(t, "a", a.handled)
}

type st1 struct {
	*Store
	handled ActionInterface
}

func (s *st1) Handle(payload *Payload) (finished bool) {
	s.handled = payload.Action
	return true
}
