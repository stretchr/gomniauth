package lineOAuth

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/objx"
	"net/http"
)

const (
	lineDefaultScope	string = "profile"
	lineName			string = "line"
	lineDisplayName		string = "LINE"
	lineAuthURL			string = "https://access.line.me/dialog/oauth/weblogin"
	lineTokenURL		string = "https://api.line.me/v1/oauth/accessToken"
	lineEndpointProfile	string = "https://api.line.me/v1/profile"
)

// LineProvider implements the Provider interface and provides LINE
// OAuth2 communication capabilities.
type LineProvider struct {
	config			*common.Config
	tripperFactory	common.TripperFactory
}

func New(clientId, clientSecret, redirectUrl string) *LineProvider {
	p := new(LineProvider)
	p.config = &common.Config{Map: objx.MSI(
			oauth2.OAuth2KeyAuthURL, lineAuthURL,
			oauth2.OAuth2KeyTokenURL, lineTokenURL,
			oauth2.OAuth2KeyClientID, clientId,
			oauth2.OAuth2KeySecret, clientSecret,
			oauth2.OAuth2KeyRedirectUrl, redirectUrl,
			oauth2.OAuth2KeyScope, lineDefaultScope,
			oauth2.OAuth2KeyAccessType, oauth2.OAuth2AccessTypeOnline,
			oauth2.OAuth2KeyApprovalPrompt, oauth2.OAuth2ApprovalPromptAuto,
			oauth2.OAuth2KeyResponseType, oauth2.OAuth2KeyCode)}
	return p
}

// TripperFactory gets an OAuth2TripperFactory
func (provider *LineProvider) TripperFactory() common.TripperFactory {
	if provider.tripperFactory == nil {
		provider.tripperFactory = new(oauth2.OAuth2TripperFactory)
	}

	return provider.tripperFactory
}

// PublicData gets a public readable view of this provider.
func (provider *LineProvider) PublicData(options map[string]interface{}) (interface{}, error) {
	return gomniauth.ProviderPublicData(provider, options)
}

// Name is the unique name for the provider.
func (provider *LineProvider) Name() string {
	return lineName
}

// Name is the unique name for this provider.
func (provider *LineProvider) DisplayName() string {
	return lineDisplayName
}

// GetBeginAuthURL gets the URL that the client must visit in order
// to begin the authentication process.
//
// The state argument contains anything you wish to have sent back to your
// callback endpoint.
// The options argument takes any options used to configure the auth request
// sent to the provider. In the case of OAuth2, the options map can contain:
//   1. A "scope" key providing the desired scope(s). It will be merged with the default scope.
func (provider *LineProvider) GetBeginAuthURL(state *common.State, options objx.Map) (string, error) {
	if options != nil {
		scope := oauth2.MergeScopes(options.Get(oauth2.OAuth2KeyScope).Str(), lineDefaultScope)
		provider.config.Set(oauth2.OAuth2KeyScope, scope)
	}
	return oauth2.GetBeginAuthURLWithBase(provider.config.Get(oauth2.OAuth2KeyAuthURL).Str(), state, provider.config)
}

// Get makes an authenticated request and returns the data in the
// response as a data map.
func (provider *LineProvider) Get(creds *common.Credentials, endpoint string) (objx.Map, error) {
	return oauth2.Get(provider, creds, endpoint)
}

// GetUser uses the specified common.Credintials to access the users profile
// from the remote provider, and builds the appropriate User object.
func (provider *LineProvider) GetUser(creds *common.Credentials) (common.User, error) {

	profileData, err := provider.Get(creds, lineEndpointProfile)

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
func (provider *LineProvider) CompleteAuth(data objx.Map) (*common.Credentials, error) {
	return oauth2.CompleteAuth(provider.TripperFactory(), data, provider.config, provider)
}

// GetClient returns an authenticated http.Client that can be used to make requests to
// protected line resources
func (provider *LineProvider) GetClient(creds *common.Credentials) (*http.Client, error) {
	return oauth2.GetClient(provider.TripperFactory(), creds, provider)
}
