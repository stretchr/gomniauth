package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetUser(t *testing.T) {

	testProvider := new(test.TestProvider)
	creds := new(common.Credentials)

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)
	testTripperFactory.On("NewTripper", creds, testProvider).Return(testTripper, nil)
	testResponse := new(http.Response)
	testResponse.Header = make(http.Header)
	testResponse.Header.Set("Content-Type", "application/json")
	testResponse.StatusCode = 200
	testResponse.Body = ioutil.NopCloser(strings.NewReader(`{"name":"their-name","id":"uniqueid","login":"loginname","email":"email@address.com","avatar_url":"http://myface.com/","blog":"http://blog.com/"}`))
	testTripper.On("RoundTrip", mock.Anything).Return(testResponse, nil)

	client := &http.Client{Transport: testTripper}
	testProvider.On("GetClient", creds).Return(client, nil)

	data, err := Get(testProvider, creds, "endpoint")

	if assert.NoError(t, err) && assert.NotNil(t, data) {

		assert.Equal(t, data["name"], "their-name")
		assert.Equal(t, data["id"], "uniqueid")
		assert.Equal(t, data["login"], "loginname")
		assert.Equal(t, data["email"], "email@address.com")
		assert.Equal(t, data["avatar_url"], "http://myface.com/")
		assert.Equal(t, data["blog"], "http://blog.com/")

	}

}
