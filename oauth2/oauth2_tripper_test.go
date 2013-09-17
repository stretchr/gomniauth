package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	testifyhttp "github.com/stretchr/testify/http"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestNewOAuth2Tripper(t *testing.T) {

	testProvider := new(test.TestProvider)
	creds := &common.Credentials{Map: objx.MSI()}
	var tripper common.Tripper = NewOAuth2Tripper(creds, testProvider)

	if assert.NotNil(t, tripper) {
		assert.Equal(t, creds, tripper.Credentials())
		assert.Equal(t, http.DefaultTransport, tripper.(*OAuth2Tripper).underlyingTransport)
		assert.Equal(t, testProvider, tripper.Provider())
	}

}

func TestRoundTrip(t *testing.T) {

	underlyingTripper := new(testifyhttp.TestRoundTripper)
	testProvider := new(test.TestProvider)
	creds := &common.Credentials{Map: objx.MSI()}
	creds.Set(OAuth2KeyAccessToken, "This is a real access token :)")

	tripper := new(OAuth2Tripper)
	tripper.underlyingTransport = underlyingTripper
	tripper.credentials = creds
	tripper.provider = testProvider

	request, _ := http.NewRequest("GET", "something", nil)

	underlyingTripper.On("RoundTrip", mock.Anything).Return(new(http.Response), nil)

	response, err := tripper.RoundTrip(request)

	if assert.NoError(t, err) {
		if assert.NotNil(t, response) {

			actualRequest := underlyingTripper.Calls[0].Arguments[0].(*http.Request)

			if assert.NotEqual(t, &actualRequest, &request, "Actual request should be different") {
				headerK, headerV := AuthorizationHeader(creds)
				assert.Equal(t, actualRequest.Header.Get(headerK), headerV)
			}

		}
	}

	mock.AssertExpectationsForObjects(t, testProvider.Mock, underlyingTripper.Mock)

}
