package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/testify/mock"
)

type TestAuthStore struct {
	mock.Mock
}

func (s *TestAuthStore) GetAuth(id string) (*common.Auth, error) {
	args := s.Called(id)
	return args.Get(0).(*common.Auth), args.Error(1)
}
func (s *TestAuthStore) PutAuth(id string, auth *common.Auth) error {
	return s.Called(id, auth).Error(0)
}
func (s *TestAuthStore) DeleteAuth(id string) error {
	return s.Called(id).Error(0)
}
