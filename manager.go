package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
)

type Manager struct {
	authStore AuthStore
	providers map[string]common.Provider
}

func NewManager(authStore AuthStore, providers ...common.Provider) *Manager {
	m := &Manager{authStore: authStore}

	m.providers = make(map[string]common.Provider)

	for _, provider := range providers {
		m.providers[provider.Name()] = provider
	}

	return m
}

func (m *Manager) Providers() map[string]common.Provider {
	return m.providers
}

func (m *Manager) WithID(id string) *Session {
	return NewSession(m, id)
}
