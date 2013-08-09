package oauth2

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

func TestOAuth2Provider_CompleteAuth_URLEncodedResponse(t *testing.T) {

	config := &common.Config{
		objects.M(
			OAuth2KeyRedirectUrl, OAuth2KeyRedirectUrl,
			OAuth2KeyScope, OAuth2KeyScope,
			OAuth2KeyClientID, OAuth2KeyClientID,
			OAuth2KeySecret, OAuth2KeySecret,
			OAuth2KeyAuthURL, OAuth2KeyAuthURL,
			OAuth2KeyTokenURL, OAuth2KeyTokenURL)}

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)
	testProvider := new(test.TestProvider)

	creds := new(common.Credentials)

	testResponse := new(http.Response)
	testResponse.Header = make(http.Header)
	testResponse.Header.Set("Content-Type", "text/plain")
	testResponse.StatusCode = 200
	testResponse.Body = ioutil.NopCloser(strings.NewReader("expires_in=20&access_token=ACCESSTOKEN&refresh_token=REFRESHTOKEN"))

	testTripperFactory.On("NewTripper", common.EmptyCredentials, mock.Anything).Return(testTripper, nil)
	testTripper.On("RoundTrip", mock.Anything).Return(testResponse, nil)

	data := objects.M(OAuth2KeyCode, "123")
	creds, err := CompleteAuth(testTripperFactory, data, config, testProvider)

	if assert.NoError(t, err) {
		if assert.NotNil(t, creds, "Creds should be returned") {

			assert.Equal(t, creds.GetStringOrEmpty(OAuth2KeyAccessToken), "ACCESSTOKEN")
			assert.Equal(t, creds.GetStringOrEmpty(OAuth2KeyRefreshToken), "REFRESHTOKEN")
			assert.Equal(t, creds.Get(OAuth2KeyExpiresIn).(time.Duration), 20000000000)

		}
	}

	mock.AssertExpectationsForObjects(t, testTripperFactory.Mock, testTripper.Mock, testProvider.Mock)

}

func TestOAuth2Provider_Non200Response(t *testing.T) {

	config := &common.Config{
		objects.M(
			OAuth2KeyRedirectUrl, OAuth2KeyRedirectUrl,
			OAuth2KeyScope, OAuth2KeyScope,
			OAuth2KeyClientID, OAuth2KeyClientID,
			OAuth2KeySecret, OAuth2KeySecret,
			OAuth2KeyAuthURL, OAuth2KeyAuthURL,
			OAuth2KeyTokenURL, OAuth2KeyTokenURL)}

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)
	testProvider := new(test.TestProvider)

	testResponse := new(http.Response)
	testResponse.Header = make(http.Header)
	testResponse.Header.Set("Content-Type", "text/plain")
	testResponse.StatusCode = 401
	testResponse.Body = ioutil.NopCloser(strings.NewReader("No mate"))

	testTripperFactory.On("NewTripper", common.EmptyCredentials, mock.Anything).Return(testTripper, nil)
	testTripper.On("RoundTrip", mock.Anything).Return(testResponse, nil)

	data := objects.M(OAuth2KeyCode, "123")
	_, err := CompleteAuth(testTripperFactory, data, config, testProvider)

	if assert.Error(t, err) {
		assert.IsType(t, &common.AuthServerError{}, err)
	}

	mock.AssertExpectationsForObjects(t, testTripperFactory.Mock, testTripper.Mock, testProvider.Mock)

}

func TestOAuth2Provider_CompleteAuth_JSON(t *testing.T) {

	config := &common.Config{
		objects.M(
			OAuth2KeyRedirectUrl, OAuth2KeyRedirectUrl,
			OAuth2KeyScope, OAuth2KeyScope,
			OAuth2KeyClientID, OAuth2KeyClientID,
			OAuth2KeySecret, OAuth2KeySecret,
			OAuth2KeyAuthURL, OAuth2KeyAuthURL,
			OAuth2KeyTokenURL, OAuth2KeyTokenURL)}

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)
	testProvider := new(test.TestProvider)

	creds := new(common.Credentials)

	testResponse := new(http.Response)
	testResponse.Header = make(http.Header)
	testResponse.Header.Set("Content-Type", "application/json")
	testResponse.StatusCode = 200
	testResponse.Body = ioutil.NopCloser(strings.NewReader(`{"expires_in":20,"access_token":"ACCESSTOKEN","refresh_token":"REFRESHTOKEN"}`))

	testTripperFactory.On("NewTripper", common.EmptyCredentials, mock.Anything).Return(testTripper, nil)
	testTripper.On("RoundTrip", mock.Anything).Return(testResponse, nil)

	data := objects.M(OAuth2KeyCode, "123")
	creds, err := CompleteAuth(testTripperFactory, data, config, testProvider)

	if assert.NoError(t, err) {
		if assert.NotNil(t, creds, "Creds should be returned") {

			assert.Equal(t, creds.GetStringOrEmpty(OAuth2KeyAccessToken), "ACCESSTOKEN")
			assert.Equal(t, creds.GetStringOrEmpty(OAuth2KeyRefreshToken), "REFRESHTOKEN")
			assert.Equal(t, creds.Get(OAuth2KeyExpiresIn).(time.Duration), 20000000000)

		}
	}

	mock.AssertExpectationsForObjects(t, testTripperFactory.Mock, testTripper.Mock, testProvider.Mock)

}
