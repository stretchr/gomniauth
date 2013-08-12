package oauth2

import (
	"github.com/stretchr/gomniauth/common"
)

// OAuth2TripperFactory provides the NewTripper function that creates a Tripper object for the OAuth2
// authentication protocol.
type OAuth2TripperFactory struct{}

// NewTripper creates a Tripper object for the OAuth2 authentication protocol.
func (f *OAuth2TripperFactory) NewTripper(creds *common.Credentials, provider common.Provider) (common.Tripper, error) {
	return NewOAuth2Tripper(creds, provider), nil
}
