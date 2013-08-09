package test

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/testify/mock"
)

type TestTripperFactory struct {
	mock.Mock
}

func (t *TestTripperFactory) NewTripper(creds *common.Credentials, provider common.Provider) (common.Tripper, error) {
	args := t.Called(creds, provider)
	return args.Get(0).(common.Tripper), args.Error(1)
}
