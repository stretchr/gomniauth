package providers

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"net/http"
)

const (
	githubDefaultScope string = "user"
	githubName         string = "github"
)

type GithubProvider struct {
	OAuth2Provider
}

func Github(clientId, clientSecret, redirectUrl string) *GithubProvider {

	p := new(GithubProvider)
	p.config = &common.Config{objects.M(
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
	return g.GetBeginAuthURLWithBase("https://github.com/login/oauth/authorize", state, g.config)
}

// CompleteAuth takes a map of arguments that are used to
// complete the authorisation process, completes it, and returns
// the appropriate common.Credentials.
func (g *GithubProvider) CompleteAuth(data objects.Map) (*common.Credentials, error) {
	return nil, nil
}

// LoadUser uses the specified common.Credentials to access the users profile
// from the remote provider, and builds the appropriate User object.
func (g *GithubProvider) LoadUser(creds *common.Credentials) (common.User, error) {
	return nil, nil
}

// GetClient gets an http.Client authenticated with the specified
// common.Credentials.
func (g *GithubProvider) GetClient(creds *common.Credentials) (*http.Client, error) {
	return nil, nil
}
