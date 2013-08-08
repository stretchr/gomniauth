package providers

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
)

type OAuth2TripperFactory struct{}

func (f *OAuth2TripperFactory) NewTripper(creds *common.Credentials) (gomniauth.Tripper, error) {
	return NewOAuth2Tripper(creds), nil
}
