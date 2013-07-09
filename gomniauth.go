package gomniauth

import (
	"code.google.com/p/goauth2/oauth"
	stewstrings "github.com/stretchr/stew/strings"
	"net/http"
	"strings"
)

type Gomniauth struct {
	// baseURL is the url preceeding the dynamic generated URL
	// Example: "http://stretchr.com/~auth" becomes "http://stretchr.com/~auth/github/callback"
	baseURL   string
	providers map[string]*oauth.Transport
}

func MakeGomniauth(baseURL string) *Gomniauth {
	return &Gomniauth{strings.Trim(baseURL, "/"),
		make(map[string]*oauth.Transport),
	}
}

func (g *Gomniauth) AddProvider(provider string, id, secret string, scopes ...string) {

	config := providerBaseConfigs[provider]

	config.ClientId = id
	config.ClientSecret = secret
	config.Scope = stewstrings.JoinStrings(",", scopes...)
	config.RedirectURL = stewstrings.MergeStrings(g.baseURL, "/", config.RedirectURL)

	g.providers[provider] = &oauth.Transport{Config: &config}

}

func (g *Gomniauth) RedirectURL(provider, state string) string {
	return g.providers[provider].Config.AuthCodeURL(state)
}

func (g *Gomniauth) Exchange(provider, code string) {
	g.providers[provider].Exchange(code)
}

func (g *Gomniauth) Get(provider, url string) (*http.Response, error) {
	return g.providers[provider].Client().Get(url)
}
