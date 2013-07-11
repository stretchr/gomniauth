package gomniauth

import (
	"github.com/stretchr/testify/mock"
)

type TestAuthStore struct {
	mock.Mock
}

func (s *TestAuthStore) GetAuth(id string) (*Auth, error) {
	args := s.Called(id)
	return args.Get(0).(*Auth), args.Error(1)
}
func (s *TestAuthStore) PutAuth(id string, auth *Auth) error {
	return s.Called(id, auth).Error(0)
}
