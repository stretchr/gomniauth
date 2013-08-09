package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
)

type OAuth2TripperFactory struct{}

func (f *OAuth2TripperFactory) NewTripper(creds *common.Credentials, provider common.Provider) (common.Tripper, error) {
	return NewOAuth2Tripper(creds, provider), nil
}
