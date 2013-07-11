package providers

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/stew/objects"
	"strings"
)

type GithubProvider struct {
	config objects.Map
}

func (p *GithubProvider) Name() string {
	return "github"
}

func (p *GithubProvider) Config() objects.Map {

	if p.config == nil {
		p.config = objects.NewMap(
			"authURL", "https://github.com/login/oauth/authorize",
			"tokenURL", "https://github.com/login/oauth/access_token",
		)
	}

	return p.config

}

func (p *GithubProvider) AuthType() common.AuthType {
	return common.AuthTypeOAuth2
}

var Github oauth2.ProviderFunc = func(clientId, clientSecret, redirectURL string, scope ...string) common.Provider {
	p := new(GithubProvider)
	p.Config().
		Set("clientId", clientId).
		Set("clientSecret", clientSecret).
		Set("redirectURL", redirectURL).
		Set("scope", strings.Join(scope, ","))
	return p
}
