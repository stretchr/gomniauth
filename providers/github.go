package providers

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
)

const (
	githubDefaultScope string = "user"
	githubName         string = "github"
	githubAuthURL      string = "https://github.com/login/oauth/authorize"
	githubTokenURL     string = "https://github.com/login/oauth/access_token"
)

type GithubProvider struct {
	gomniauth.OAuth2Provider
}

func Github(clientId, clientSecret, redirectUrl string) *GithubProvider {

	p := new(GithubProvider)
	p.Config = &common.Config{objects.M(
		gomniauth.OAuth2KeyAuthURL, githubAuthURL,
		gomniauth.OAuth2KeyTokenURL, githubTokenURL,
		gomniauth.OAuth2KeyClientID, clientId,
		gomniauth.OAuth2KeySecret, clientSecret,
		gomniauth.OAuth2KeyRedirectUrl, redirectUrl,
		gomniauth.OAuth2KeyScope, githubDefaultScope,
		gomniauth.OAuth2KeyAccessType, gomniauth.OAuth2AccessTypeOnline,
		gomniauth.OAuth2KeyApprovalPrompt, gomniauth.OAuth2ApprovalPromptAuto)}
	return p
}

// Name is the unique name for this provider.
func (g *GithubProvider) Name() string {
	return githubName
}

// GetBeginAuthURL gets the URL that the client must visit in order
// to begin the authentication process.
func (g *GithubProvider) GetBeginAuthURL(state *common.State) (string, error) {
	return g.GetBeginAuthURLWithBase(g.Config.GetString(gomniauth.OAuth2KeyAuthURL), state, g.Config)
}

// LoadUser uses the specified common.Credentials to access the users profile
// from the remote provider, and builds the appropriate User object.
func (g *GithubProvider) LoadUser(creds *common.Credentials) (common.User, error) {
	return nil, nil
}
