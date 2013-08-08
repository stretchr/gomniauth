package providers

import (
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
	OAuth2Provider
}

func Github(clientId, clientSecret, redirectUrl string) *GithubProvider {

	p := new(GithubProvider)
	p.config = &common.Config{objects.M(
		OAuth2KeyAuthURL, githubAuthURL,
		OAuth2KeyTokenURL, githubTokenURL,
		OAuth2KeyClientID, clientId,
		OAuth2KeySecret, clientSecret,
		OAuth2KeyRedirectUrl, redirectUrl,
		OAuth2KeyScope, githubDefaultScope,
		OAuth2KeyAccessType, OAuth2AccessTypeOnline,
		OAuth2KeyApprovalPrompt, OAuth2ApprovalPromptAuto)}
	return p
}

// Name is the unique name for this provider.
func (g *GithubProvider) Name() string {
	return githubName
}

// GetBeginAuthURL gets the URL that the client must visit in order
// to begin the authentication process.
func (g *GithubProvider) GetBeginAuthURL(state *common.State) (string, error) {
	return g.GetBeginAuthURLWithBase(g.config.GetString(OAuth2KeyAuthURL), state, g.config)
}

// LoadUser uses the specified common.Credentials to access the users profile
// from the remote provider, and builds the appropriate User object.
func (g *GithubProvider) LoadUser(creds *common.Credentials) (common.User, error) {
	return nil, nil
}
