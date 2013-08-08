package test

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/testify/mock"
)

type TestTripperFactory struct {
	mock.Mock
}

func (t *TestTripperFactory) NewTripper(creds *common.Credentials, provider gomniauth.Provider) (gomniauth.Tripper, error) {
	args := t.Called(creds, provider)
	return args.Get(0).(gomniauth.Tripper), args.Error(1)
}
