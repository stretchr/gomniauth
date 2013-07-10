package gomniauth

import (
	"github.com/stretchr/testify/mock"
)

type TestTokenStore struct {
	mock.Mock
}

func (s *TestTokenStore) GetToken(id string) (*Token, error) {
	args := s.Called(id)
	return args.Get(0).(*Token), args.Error(1)
}
func (s *TestTokenStore) PutToken(id string, token *Token) error {
	return s.Called(id, token).Error(0)
}
