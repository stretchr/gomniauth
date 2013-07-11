package providers

import (
	"github.com/stretchr/gomniauth/common"
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

var Github = func(clientId, clientSecret, redirectURL string, scope ...string) *GithubProvider {
	p := new(GithubProvider)
	p.Config().
		Set("clientId", clientId).
		Set("clientSecret", clientSecret).
		Set("redirectURL", redirectURL).
		Set("scope", strings.Join(scope, ","))
	return p
}
