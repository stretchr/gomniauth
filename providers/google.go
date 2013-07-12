package providers

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/stew/objects"
	"strings"
)

type GoogleProvider struct {
	config objects.Map
}

func (p *GoogleProvider) Name() string {
	return "Google"
}

func (p *GoogleProvider) Config() objects.Map {

	if p.config == nil {
		p.config = objects.NewMap(
			"authURL", "https://accounts.google.com/o/oauth2/auth",
			"tokenURL", "https://accounts.google.com/o/oauth2/token",
		)
	}

	return p.config

}

func (p *GoogleProvider) AuthType() common.AuthType {
	return common.AuthTypeOAuth2
}

var Google oauth2.ProviderFunc = func(clientId, clientSecret, redirectURL string, scope ...string) common.Provider {
	p := new(GoogleProvider)
	p.Config().
		Set("clientId", clientId).
		Set("clientSecret", clientSecret).
		Set("redirectURL", redirectURL).
		Set("scope", strings.Join(scope, ","))
	return p
}
