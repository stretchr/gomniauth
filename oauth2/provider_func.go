package oauth2

import (
	"github.com/stretchr/gomniauth/common"
)

// ProviderFunc is a function that configures and returns a specific
// provider given the essential OAuth2 information specified.
type ProviderFunc func(clientId, clientSecret, redirectURL string, scope ...string) common.Provider
