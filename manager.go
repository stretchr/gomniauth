package gomniauth

type Manager struct {
	tokenStore TokenStore
	providers  map[string]*Provider
}

func NewManager(tokenStore TokenStore) *Manager {
	m := &Manager{tokenStore: tokenStore}

	// TODO: install default providers

	return m
}

func (m *Manager) WithID(id string) *Session {
	return NewSession(m, id)
}
