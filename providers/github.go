package providers

import (
	"github.com/stretchr/codecs/services"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/stew/objects"
	"io/ioutil"
	"net/http"
)

const (
	githubDefaultScope    string = "user"
	githubName            string = "github"
	githubAuthURL         string = "https://github.com/login/oauth/authorize"
	githubTokenURL        string = "https://github.com/login/oauth/access_token"
	githubEndpointProfile string = "https://api.github.com/user"
)

type GithubProvider struct {
	config         *common.Config
	tripperFactory common.TripperFactory
}

func Github(clientId, clientSecret, redirectUrl string) *GithubProvider {

	p := new(GithubProvider)
	p.config = &common.Config{objects.M(
		oauth2.OAuth2KeyAuthURL, githubAuthURL,
		oauth2.OAuth2KeyTokenURL, githubTokenURL,
		oauth2.OAuth2KeyClientID, clientId,
		oauth2.OAuth2KeySecret, clientSecret,
		oauth2.OAuth2KeyRedirectUrl, redirectUrl,
		oauth2.OAuth2KeyScope, githubDefaultScope,
		oauth2.OAuth2KeyAccessType, oauth2.OAuth2AccessTypeOnline,
		oauth2.OAuth2KeyApprovalPrompt, oauth2.OAuth2ApprovalPromptAuto)}
	return p
}

func (g *GithubProvider) TripperFactory() common.TripperFactory {

	if g.tripperFactory == nil {
		g.tripperFactory = new(oauth2.OAuth2TripperFactory)
	}

	return g.tripperFactory
}

// Name is the unique name for this provider.
func (g *GithubProvider) Name() string {
	return githubName
}

// GetBeginAuthURL gets the URL that the client must visit in order
// to begin the authentication process.
func (g *GithubProvider) GetBeginAuthURL(state *common.State) (string, error) {
	return oauth2.GetBeginAuthURLWithBase(g.config.GetString(oauth2.OAuth2KeyAuthURL), state, g.config)
}

// LoadUser uses the specified common.Credentials to access the users profile
// from the remote provider, and builds the appropriate User object.
func (g *GithubProvider) LoadUser(creds *common.Credentials) (common.User, error) {

	client, clientErr := g.GetClient(creds)

	if clientErr != nil {
		return nil, clientErr
	}

	response, responseErr := client.Get(githubEndpointProfile)

	if responseErr != nil {
		return nil, responseErr
	}

	body, bodyErr := ioutil.ReadAll(response.Body)

	if bodyErr != nil {
		return nil, bodyErr
	}

	defer response.Body.Close()

	codecs := services.NewWebCodecService()
	codec, getCodecErr := codecs.GetCodec(response.Header.Get("Content-Type"))

	if getCodecErr != nil {
		return nil, getCodecErr
	}

	var data objects.Map
	unmarshalErr := codec.Unmarshal(body, &data)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	// build user
	//user := gomniauth.User(data)

	return nil, nil
}

// CompleteAuth takes a map of arguments that are used to
// complete the authorisation process, completes it, and returns
// the appropriate Credentials.
func (g *GithubProvider) CompleteAuth(data objects.Map) (*common.Credentials, error) {
	return oauth2.CompleteAuth(g.TripperFactory(), data, g.config, g)
}

func (g *GithubProvider) GetClient(creds *common.Credentials) (*http.Client, error) {
	return oauth2.GetClient(g.TripperFactory(), creds, g)
}
