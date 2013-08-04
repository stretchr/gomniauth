package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSession(t *testing.T) {

	manager := new(Manager)
	id := "abc123"
	prov := new(common.TestProvider)
	s := NewSession(manager, id, prov)

	if assert.NotNil(t, s) {
		assert.Equal(t, manager, s.Manager())
		assert.Equal(t, id, s.ID())
		assert.Equal(t, prov, s.Provider())
	}

}

func TestNilSessionIsOK(t *testing.T) {

	var s *Session = nil

	assert.NoError(t, s.HandleCallback(nil))

}
