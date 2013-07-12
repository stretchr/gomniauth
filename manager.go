package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"strings"
)

// Manager provides access to omniauth functionality.
//
// Manager keeps track of Providers that are supported by
// the application, and allows Sessions to access the AuthStore.
//
// Your application should have only one manager, and use the WithID method
// to get a Session for each request.
type Manager struct {
	authStore AuthStore
	providers map[string]common.Provider
}

// NewManager creates a new Manager configured with the specified AuthStore
// and any providers specified.
func NewManager(authStore AuthStore, providers ...common.Provider) *Manager {
	m := &Manager{authStore: authStore}

	m.providers = make(map[string]common.Provider)

	for _, provider := range providers {
		m.AddProvider(provider)
	}

	return m
}

// Providers gets the map of currently installed Providers.
func (m *Manager) Providers() map[string]common.Provider {
	return m.providers
}

// AddProvider adds the specified provider.
func (m *Manager) AddProvider(provider common.Provider) *Manager {
	m.providers[strings.ToLower(provider.Name())] = provider
	return m
}

// Provider gets a provider by name.
func (m *Manager) Provider(name string) common.Provider {
	return m.Providers()[strings.ToLower(name)]
}

// WithID creates a new Session with the specified ID.
func (m *Manager) WithID(id string) *Session {
	return NewSession(m, id)
}
