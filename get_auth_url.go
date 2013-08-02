package gomniauth

import (
	"errors"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/stew/objects"
)

// GetAuthURL uses the specified provider and state objects to build the
// URL which the user must be redirected to in order to get authenticated.
func GetAuthURL(provider common.Provider, state objects.Map, stateSecurityKey string) (string, error) {

	switch provider.AuthType() {
	case common.AuthTypeOAuth2:

		// make the config
		var config = &oauth2.Config{
			Map: provider.Config().Copy(),
		}

		encodedState, encodedStateErr := state.SignedBase64(stateSecurityKey)

		if encodedStateErr != nil {
			return "", encodedStateErr
		}

		return config.AuthCodeURL(encodedState), nil

	}

	return "", errors.New("gomniauth: GetAuthURL: Unsupported common.AuthType: " + provider.AuthType().String())

}
