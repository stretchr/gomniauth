package spotify

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGitHubImplementrsProvider(t *testing.T) {

	var provider common.Provider
	provider = new(SpotifyProvider)

	assert.NotNil(t, provider)

}

func TestGetUser(t *testing.T) {

	g := New("clientID", "secret", "http://myapp.com/")
	creds := &common.Credentials{Map: objx.MSI()}

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)
	testTripperFactory.On("NewTripper", mock.Anything, g).Return(testTripper, nil)
	testResponse := new(http.Response)
	testResponse.Header = make(http.Header)
	testResponse.Header.Set("Content-Type", "application/json")
	testResponse.StatusCode = 200
	testResponse.Body = ioutil.NopCloser(strings.NewReader(`{
  "display_name":"JMWizzler",
  "email":"email@example.com",
  "external_urls":{
    "spotify":"https://open.spotify.com/user/wizzler"
  },
  "href":"https://api.spotify.com/v1/users/wizzler",
  "id":"wizzler",
  "images":[{
    "height":null,
    "url":"https://fbcdn.example.com/2330_n.jpg",
    "width":null
  }],
  "type":"user",
  "uri":"spotify:user:wizzler"
}`))
	testTripper.On("RoundTrip", mock.Anything).Return(testResponse, nil)

	g.tripperFactory = testTripperFactory

	user, err := g.GetUser(creds)

	if assert.NoError(t, err) && assert.NotNil(t, user) {

		assert.Equal(t, user.Name(), "JMWizzler")
		assert.Equal(t, user.AuthCode(), "") // doesn't come from spotify
		assert.Equal(t, user.Nickname(), "") // doesn't come from spotify
		assert.Equal(t, user.Email(), "email@example.com")
		assert.Equal(t, user.AvatarURL(), "https://fbcdn.example.com/2330_n.jpg")
		assert.Equal(t, user.Data()["href"], "https://api.spotify.com/v1/users/wizzler")

		spotifyCreds := user.ProviderCredentials()[spotifyName]
		if assert.NotNil(t, spotifyCreds) {
			assert.Equal(t, "wizzler", spotifyCreds.Get(common.CredentialsKeyID).Str())
		}

	}

}

func TestNewspotify(t *testing.T) {

	g := New("clientID", "secret", "http://myapp.com/")

	if assert.NotNil(t, g) {

		// check config
		if assert.NotNil(t, g.config) {

			assert.Equal(t, "clientID", g.config.Get(oauth2.OAuth2KeyClientID).Data())
			assert.Equal(t, "secret", g.config.Get(oauth2.OAuth2KeySecret).Data())
			assert.Equal(t, "http://myapp.com/", g.config.Get(oauth2.OAuth2KeyRedirectUrl).Data())
			assert.Equal(t, spotifyDefaultScope, g.config.Get(oauth2.OAuth2KeyScope).Data())

			assert.Equal(t, spotifyAuthURL, g.config.Get(oauth2.OAuth2KeyAuthURL).Data())
			assert.Equal(t, spotifyTokenURL, g.config.Get(oauth2.OAuth2KeyTokenURL).Data())

		}

	}

}

func TestspotifyTripperFactory(t *testing.T) {

	g := New("clientID", "secret", "http://myapp.com/")
	g.tripperFactory = nil

	f := g.TripperFactory()

	if assert.NotNil(t, f) {
		assert.Equal(t, f, g.tripperFactory)
	}

}

func TestspotifyName(t *testing.T) {
	g := New("clientID", "secret", "http://myapp.com/")
	assert.Equal(t, spotifyName, g.Name())
}

func TestspotifyGetBeginAuthURL(t *testing.T) {

	common.SetSecurityKey("ABC123")

	state := &common.State{Map: objx.MSI("after", "http://www.stretchr.com/")}

	g := New("clientID", "secret", "http://myapp.com/")

	url, err := g.GetBeginAuthURL(state, nil)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=clientID")
		assert.Contains(t, url, "redirect_uri=http%3A%2F%2Fmyapp.com%2F")
		assert.Contains(t, url, "scope="+spotifyDefaultScope)
		assert.Contains(t, url, "access_type="+oauth2.OAuth2AccessTypeOnline)
		assert.Contains(t, url, "approval_prompt="+oauth2.OAuth2ApprovalPromptAuto)
	}

	state = &common.State{Map: objx.MSI("after", "http://www.stretchr.com/")}

	g = New("clientID", "secret", "http://myapp.com/")

	url, err = g.GetBeginAuthURL(state, objx.MSI(oauth2.OAuth2KeyScope, "avatar"))

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=clientID")
		assert.Contains(t, url, "redirect_uri=http%3A%2F%2Fmyapp.com%2F")
		assert.Contains(t, url, "scope=avatar+"+spotifyDefaultScope)
		assert.Contains(t, url, "access_type="+oauth2.OAuth2AccessTypeOnline)
		assert.Contains(t, url, "approval_prompt="+oauth2.OAuth2ApprovalPromptAuto)
	}

}

// test that scopes get added with spaces
