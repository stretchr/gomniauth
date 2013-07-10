package gomniauth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewManager(t *testing.T) {

	tokenStore := new(TestTokenStore)
	m := NewManager(tokenStore)

	if assert.NotNil(t, m) {
		assert.Equal(t, tokenStore, m.tokenStore)
	}

}

func TestManager_WithID(t *testing.T) {

	id := "abc123"

	man := new(Manager)
	session := man.WithID(id)

	assert.Equal(t, session.Manager(), man)
	assert.Equal(t, session.ID(), id)

}
