package github

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/stew/objects"
	"net/http"
)

const (
	githubDefaultScope    string = "user"
	githubName            string = "github"
	githubAuthURL         string = "https://github.com/login/oauth/authorize"
	githubTokenURL        string = "https://github.com/login/oauth/access_token"
	githubEndpointProfile string = "https://api.github.com/user"
)

// GithubProvider implements the Provider interface and provides Github
// OAuth2 communication capabilities.
type GithubProvider struct {
	config         *common.Config
	tripperFactory common.TripperFactory
}

func New(clientId, clientSecret, redirectUrl string) *GithubProvider {

	p := new(GithubProvider)
	p.config = &common.Config{Map: objects.M(
		oauth2.OAuth2KeyAuthURL, githubAuthURL,
		oauth2.OAuth2KeyTokenURL, githubTokenURL,
		oauth2.OAuth2KeyClientID, clientId,
		oauth2.OAuth2KeySecret, clientSecret,
		oauth2.OAuth2KeyRedirectUrl, redirectUrl,
		oauth2.OAuth2KeyScope, githubDefaultScope,
		oauth2.OAuth2KeyAccessType, oauth2.OAuth2AccessTypeOnline,
		oauth2.OAuth2KeyApprovalPrompt, oauth2.OAuth2ApprovalPromptAuto,
		oauth2.OAuth2KeyResponseType, oauth2.OAuth2KeyCode)}
	return p
}

// TripperFactory gets an OAuth2TripperFactory
func (provider *GithubProvider) TripperFactory() common.TripperFactory {

	if provider.tripperFactory == nil {
		provider.tripperFactory = new(oauth2.OAuth2TripperFactory)
	}

	return provider.tripperFactory
}

// Name is the unique name for this provider.
func (provider *GithubProvider) Name() string {
	return githubName
}

// GetBeginAuthURL gets the URL that the client must visit in order
// to begin the authentication process.
//
// The state argument contains anything you wish to have sent back to your
// callback endpoint.
// The options argument takes any options used to configure the auth request
// sent to the provider. In the case of OAuth2, the options map can contain:
//   1. A "scope" key providing the desired scope(s). It will be merged with the default scope.
func (provider *GithubProvider) GetBeginAuthURL(state *common.State, options objects.Map) (string, error) {
	if options != nil {
		scope := oauth2.MergeScopes(options.GetStringOrEmpty(oauth2.OAuth2KeyScope), githubDefaultScope)
		provider.config.Set(oauth2.OAuth2KeyScope, scope)
	}
	return oauth2.GetBeginAuthURLWithBase(provider.config.GetString(oauth2.OAuth2KeyAuthURL), state, provider.config)
}

// Get makes an authenticated request and returns the data in the
// response as a data map.
func (provider *GithubProvider) Get(creds *common.Credentials, endpoint string) (objects.Map, error) {
	return oauth2.Get(provider, creds, endpoint)
}

// GetUser uses the specified common.Credentials to access the users profile
// from the remote provider, and builds the appropriate User object.
func (provider *GithubProvider) GetUser(creds *common.Credentials) (common.User, error) {

	profileData, err := provider.Get(creds, githubEndpointProfile)

	if err != nil {
		return nil, err
	}

	// build user
	user := NewUser(profileData, creds, provider)

	return user, nil
}

// CompleteAuth takes a map of arguments that are used to
// complete the authorisation process, completes it, and returns
// the appropriate Credentials.
func (provider *GithubProvider) CompleteAuth(data objects.Map) (*common.Credentials, error) {
	return oauth2.CompleteAuth(provider.TripperFactory(), data, provider.config, provider)
}

// GetClient returns an authenticated http.Client that can be used to make requests to
// protected Github resources
func (provider *GithubProvider) GetClient(creds *common.Credentials) (*http.Client, error) {
	return oauth2.GetClient(provider.TripperFactory(), creds, provider)
}
