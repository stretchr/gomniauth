package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestOAuth2HandlerBeginAuthURLWithBase(t *testing.T) {

	SecurityKey = "ABC123"

	h := &OAuth2Provider{}
	base := "https://base.url/auth"

	config := &common.Config{objects.M()}
	config.
		Set("client_id", "client_id").
		Set("redirect_url", "redirect_url").
		Set("scope", "scope").
		Set("access_type", "access_type").
		Set("approval_prompt", "approval_prompt")

	state := &common.State{objects.M("after", "http://www.stretchr.com/")}
	base64State, _ := state.Base64()

	url, err := h.GetBeginAuthURLWithBase(base, state, config)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=client_id")
		assert.Contains(t, url, "redirect_url=redirect_url")
		assert.Contains(t, url, "scope=scope")
		assert.Contains(t, url, "access_type=access_type")
		assert.Contains(t, url, "approval_prompt=approval_prompt")
		assert.Contains(t, url, "state="+base64State)
	}

}

func TestOAuth2Provider_CompleteAuth_URLEncodedResponse(t *testing.T) {

	g := &OAuth2Provider{} // ("clientID", "secret", "http://myapp.com/")

	g.Config = &common.Config{
		objects.M(
			OAuth2KeyRedirectUrl, OAuth2KeyRedirectUrl,
			OAuth2KeyScope, OAuth2KeyScope,
			OAuth2KeyClientID, OAuth2KeyClientID,
			OAuth2KeySecret, OAuth2KeySecret,
			OAuth2KeyAuthURL, OAuth2KeyAuthURL,
			OAuth2KeyTokenURL, OAuth2KeyTokenURL)}

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)

	g.SetTripperFactory(testTripperFactory)
	creds := new(common.Credentials)

	testResponse := new(http.Response)
	testResponse.Header = make(http.Header)
	testResponse.Header.Set("Content-Type", "text/plain")
	testResponse.StatusCode = 200
	testResponse.Body = ioutil.NopCloser(strings.NewReader("expires_in=20&access_token=ACCESSTOKEN&refresh_token=REFRESHTOKEN"))

	testTripperFactory.On("NewTripper", common.EmptyCredentials, mock.Anything).Return(testTripper, nil)
	testTripper.On("RoundTrip", mock.Anything).Return(testResponse, nil)

	data := objects.M(OAuth2KeyCode, "123")
	creds, err := g.CompleteAuth(data)

	if assert.NoError(t, err) {
		if assert.NotNil(t, creds, "Creds should be returned") {

			assert.Equal(t, creds.GetStringOrEmpty(OAuth2KeyAccessToken), "ACCESSTOKEN")
			assert.Equal(t, creds.GetStringOrEmpty(OAuth2KeyRefreshToken), "REFRESHTOKEN")
			assert.Equal(t, creds.Get(OAuth2KeyExpiresIn).(time.Duration), 20000000000)

		}
	}

	mock.AssertExpectationsForObjects(t, testTripperFactory.Mock, testTripper.Mock)

}

func TestOAuth2Provider_CompleteAuth_JSON(t *testing.T) {

	g := &OAuth2Provider{} // ("clientID", "secret", "http://myapp.com/")

	g.Config = &common.Config{
		objects.M(
			OAuth2KeyRedirectUrl, OAuth2KeyRedirectUrl,
			OAuth2KeyScope, OAuth2KeyScope,
			OAuth2KeyClientID, OAuth2KeyClientID,
			OAuth2KeySecret, OAuth2KeySecret,
			OAuth2KeyAuthURL, OAuth2KeyAuthURL,
			OAuth2KeyTokenURL, OAuth2KeyTokenURL)}

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)

	g.SetTripperFactory(testTripperFactory)
	creds := new(common.Credentials)

	testResponse := new(http.Response)
	testResponse.Header = make(http.Header)
	testResponse.Header.Set("Content-Type", "application/json")
	testResponse.StatusCode = 200
	testResponse.Body = ioutil.NopCloser(strings.NewReader(`{"expires_in":20,"access_token":"ACCESSTOKEN","refresh_token":"REFRESHTOKEN"}`))

	testTripperFactory.On("NewTripper", common.EmptyCredentials, mock.Anything).Return(testTripper, nil)
	testTripper.On("RoundTrip", mock.Anything).Return(testResponse, nil)

	data := objects.M(OAuth2KeyCode, "123")
	creds, err := g.CompleteAuth(data)

	if assert.NoError(t, err) {
		if assert.NotNil(t, creds, "Creds should be returned") {

			assert.Equal(t, creds.GetStringOrEmpty(OAuth2KeyAccessToken), "ACCESSTOKEN")
			assert.Equal(t, creds.GetStringOrEmpty(OAuth2KeyRefreshToken), "REFRESHTOKEN")
			assert.Equal(t, creds.Get(OAuth2KeyExpiresIn).(time.Duration), 20000000000)

		}
	}

	mock.AssertExpectationsForObjects(t, testTripperFactory.Mock, testTripper.Mock)

}
