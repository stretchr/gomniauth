package gomniauth

type Manager struct {
	authStore AuthStore
	providers map[string]Provider
}

func NewManager(authStore AuthStore, providers ...Provider) *Manager {
	m := &Manager{authStore: authStore}

	m.providers = make(map[string]Provider)

	for _, provider := range providers {
		m.providers[provider.Name()] = provider
	}

	return m
}

func (m *Manager) Providers() map[string]Provider {
	return m.providers
}

func (m *Manager) WithID(id string) *Session {
	return NewSession(m, id)
}
