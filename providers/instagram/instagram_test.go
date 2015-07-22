package instagram

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestInstagramImplementrsProvider(t *testing.T) {

	var provider common.Provider
	provider = new(InstagramProvider)

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
  "meta":  {
    "code": 200
  },
  "data":  {
    "username": "maggit",
    "bio": "Programer who loves golang",
    "website": "http://website.com",
    "profile_picture": "http://myinsta.com",
    "full_name": "Raquel H",
    "counts":  {
      "media": 1232323,
      "followed_by": 123,
      "follows": 123
    },
    "id": "12345678"
  }
}`))
	testTripper.On("RoundTrip", mock.Anything).Return(testResponse, nil)

	g.tripperFactory = testTripperFactory

	user, err := g.GetUser(creds)

	if assert.NoError(t, err) && assert.NotNil(t, user) {

		assert.Equal(t, user.Name(), "Raquel H")
		assert.Equal(t, user.AuthCode(), "") // doesn't come from instagram
		assert.Equal(t, user.Nickname(), "maggit")
		assert.Equal(t, user.AvatarURL(), "http://myinsta.com")
		assert.Equal(t, user.Data()["website"], "http://website.com")

		instagramCreds := user.ProviderCredentials()[instagramName]
		if assert.NotNil(t, instagramCreds) {
			assert.Equal(t, "uniqueid", instagramCreds.Get(common.CredentialsKeyID).Str())
		}

	}

}

func TestNewInstagram(t *testing.T) {

	g := New("clientID", "secret", "http://myapp.com/")

	if assert.NotNil(t, g) {

		// check config
		if assert.NotNil(t, g.config) {

			assert.Equal(t, "clientID", g.config.Get(oauth2.OAuth2KeyClientID).Data())
			assert.Equal(t, "secret", g.config.Get(oauth2.OAuth2KeySecret).Data())
			assert.Equal(t, "http://myapp.com/", g.config.Get(oauth2.OAuth2KeyRedirectUrl).Data())
			assert.Equal(t, instagramDefaultScope, g.config.Get(oauth2.OAuth2KeyScope).Data())

			assert.Equal(t, instagramAuthURL, g.config.Get(oauth2.OAuth2KeyAuthURL).Data())
			assert.Equal(t, instagramTokenURL, g.config.Get(oauth2.OAuth2KeyTokenURL).Data())

		}

	}

}

func TestInstagramTripperFactory(t *testing.T) {

	g := New("clientID", "secret", "http://myapp.com/")
	g.tripperFactory = nil

	f := g.TripperFactory()

	if assert.NotNil(t, f) {
		assert.Equal(t, f, g.tripperFactory)
	}

}

func TestInstagramName(t *testing.T) {
	g := New("clientID", "secret", "http://myapp.com/")
	assert.Equal(t, instagramName, g.Name())
}

func TestInstagramGetBeginAuthURL(t *testing.T) {

	common.SetSecurityKey("ABC123")

	state := &common.State{Map: objx.MSI("after", "http://www.stretchr.com/")}

	g := New("clientID", "secret", "http://myapp.com/")

	url, err := g.GetBeginAuthURL(state, nil)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=clientID")
		assert.Contains(t, url, "redirect_uri=http%3A%2F%2Fmyapp.com%2F")
		assert.Contains(t, url, "scope="+instagramDefaultScope)
		assert.Contains(t, url, "access_type="+oauth2.OAuth2AccessTypeOnline)
		assert.Contains(t, url, "approval_prompt="+oauth2.OAuth2ApprovalPromptAuto)
	}

	state = &common.State{Map: objx.MSI("after", "http://www.stretchr.com/")}

	g = New("clientID", "secret", "http://myapp.com/")

	url, err = g.GetBeginAuthURL(state, objx.MSI(oauth2.OAuth2KeyScope, "avatar"))

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=clientID")
		assert.Contains(t, url, "redirect_uri=http%3A%2F%2Fmyapp.com%2F")
		assert.Contains(t, url, "scope=avatar+"+instagramDefaultScope)
		assert.Contains(t, url, "access_type="+oauth2.OAuth2AccessTypeOnline)
		assert.Contains(t, url, "approval_prompt="+oauth2.OAuth2ApprovalPromptAuto)
	}

}
