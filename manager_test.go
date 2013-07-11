package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewManager(t *testing.T) {

	authStore := new(TestAuthStore)
	prov1 := new(common.TestProvider)
	prov2 := new(common.TestProvider2)

	prov1.On("Name").Return("Prov1")
	prov2.On("Name").Return("Prov2")

	m := NewManager(authStore, prov1, prov2)

	if assert.NotNil(t, m) {
		assert.Equal(t, authStore, m.authStore)
		if assert.Equal(t, 2, len(m.providers)) {
			assert.Equal(t, prov1, m.providers["Prov1"])
			assert.Equal(t, prov2, m.providers["Prov2"])
		}
	}

}

func TestManager_Providers(t *testing.T) {

	authStore := new(TestAuthStore)
	prov1 := new(common.TestProvider)
	prov2 := new(common.TestProvider2)

	prov1.On("Name").Return("Prov1")
	prov2.On("Name").Return("Prov2")

	m := NewManager(authStore, prov1, prov2)

	if assert.NotNil(t, m) {
		assert.Equal(t, authStore, m.authStore)
		if assert.Equal(t, 2, len(m.Providers())) {
			assert.Equal(t, prov1, m.Providers()["Prov1"])
			assert.Equal(t, prov2, m.Providers()["Prov2"])
		}
	}

}

func TestManager_WithID(t *testing.T) {

	id := "abc123"

	man := new(Manager)
	session := man.WithID(id)

	assert.Equal(t, session.Manager(), man)
	assert.Equal(t, session.ID(), id)

}
