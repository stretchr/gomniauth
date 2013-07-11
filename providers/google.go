package providers

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"strings"
)

type GoogleProvider struct{}

func (p *GoogleProvider) Name() string {
	return "google"
}

func (p *GoogleProvider) Config() objects.Map {
	return objects.NewMap(
		"scope", "https://www.googleapis.com/auth/userinfo.profile",
		"authURL", "https://accounts.google.com/o/oauth2/auth",
		"tokenURL", "https://accounts.google.com/o/oauth2/token",
	)
}

func (p *GoogleProvider) AuthType() common.AuthType {
	return common.AuthTypeOAuth2
}

// TODO: test this
var Google = func(clientId, clientSecret, redirectURL string, scope ...string) *GoogleProvider {
	p := new(GoogleProvider)
	p.Config().
		Set("clientId", clientId).
		Set("clientSecret", clientSecret).
		Set("redirectURL", redirectURL).
		Set("scope", strings.Join(scope, ","))
	return p
}
