package gomniauth

import (
	"net/http"
)

type Session struct {
	manager *Manager
	id      string
}

func NewSession(manager *Manager, id string) *Session {
	return &Session{manager: manager, id: id}
}

func (s *Session) Manager() *Manager {
	return s.manager
}

func (s *Session) ID() string {
	return s.id
}

// func (s *Session) HttpClient() *http.Client {

// }
