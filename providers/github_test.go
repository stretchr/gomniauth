package providers

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGitHubImplementrsProvider(t *testing.T) {

	var provider common.Provider
	provider = new(GithubProvider)

	assert.NotNil(t, provider)

}

func TestLoadUser(t *testing.T) {

	g := Github("clientID", "secret", "http://myapp.com/")

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)
	testTripperFactory.On("NewTripper", mock.Anything, g).Return(testTripper, nil)
	testResponse := new(http.Response)
	testResponse.Header = make(http.Header)
	testResponse.Header.Set("Content-Type", "application/json")
	testResponse.StatusCode = 200
	testResponse.Body = ioutil.NopCloser(strings.NewReader(`{"id":"uniqueid","login":"loginname","email":"email@address.com","avatar_url":"http://myface.com/","blog":"http://blog.com/"}`))
	testTripper.On("RoundTrip", mock.Anything).Return(testResponse, nil)

	g.tripperFactory = testTripperFactory

}

func TestNewGithub(t *testing.T) {

	g := Github("clientID", "secret", "http://myapp.com/")

	if assert.NotNil(t, g) {

		// check config
		if assert.NotNil(t, g.config) {

			assert.Equal(t, "clientID", g.config.Get(oauth2.OAuth2KeyClientID))
			assert.Equal(t, "secret", g.config.Get(oauth2.OAuth2KeySecret))
			assert.Equal(t, "http://myapp.com/", g.config.Get(oauth2.OAuth2KeyRedirectUrl))
			assert.Equal(t, githubDefaultScope, g.config.Get(oauth2.OAuth2KeyScope))

			assert.Equal(t, githubAuthURL, g.config.Get(oauth2.OAuth2KeyAuthURL))
			assert.Equal(t, githubTokenURL, g.config.Get(oauth2.OAuth2KeyTokenURL))

		}

		// check factory

	}

}

func TestGithubTripperFactory(t *testing.T) {

	g := Github("clientID", "secret", "http://myapp.com/")
	g.tripperFactory = nil

	f := g.TripperFactory()

	if assert.NotNil(t, f) {
		assert.Equal(t, f, g.tripperFactory)
	}

}

func TestGithubName(t *testing.T) {
	g := Github("clientID", "secret", "http://myapp.com/")
	assert.Equal(t, githubName, g.Name())
}

func TestGitHubGetBeginAuthURL(t *testing.T) {

	common.SetSecurityKey("ABC123")

	state := &common.State{objects.M("after", "http://www.stretchr.com/")}

	g := Github("clientID", "secret", "http://myapp.com/")

	url, err := g.GetBeginAuthURL(state)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=clientID")
		assert.Contains(t, url, "redirect_uri=http%3A%2F%2Fmyapp.com%2F")
		assert.Contains(t, url, "scope="+githubDefaultScope)
		assert.Contains(t, url, "access_type="+oauth2.OAuth2AccessTypeOnline)
		assert.Contains(t, url, "approval_prompt="+oauth2.OAuth2ApprovalPromptAuto)
	}

}
