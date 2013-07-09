package gomniauth

import (
	"code.google.com/p/goauth2/oauth"
)

const (
	Github = "github"
)

var providerBaseConfigs = map[string]oauth.Config{
	Github: oauth.Config{
		AuthURL:     "https://github.com/login/oauth/authorize",
		TokenURL:    "https://github.com/login/oauth/access_token",
		RedirectURL: "github/callback",
	},
}
