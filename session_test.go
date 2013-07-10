package gomniauth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSession(t *testing.T) {

	manager := new(Manager)
	id := "abc123"
	s := NewSession(manager, id)

	if assert.NotNil(t, s) {
		assert.Equal(t, manager, s.Manager())
		assert.Equal(t, id, s.ID())
	}

}
