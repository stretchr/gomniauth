package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewOAuth2Tripper(t *testing.T) {

	testProvider := new(test.TestProvider)
	creds := &common.Credentials{objects.M()}
	var tripper common.Tripper = NewOAuth2Tripper(creds, testProvider)

	if assert.NotNil(t, tripper) {
		assert.Equal(t, creds, tripper.Credentials())
		assert.Equal(t, http.DefaultTransport, tripper.(*OAuth2Tripper).underlyingTransport)
		assert.Equal(t, testProvider, tripper.Provider())
	}

}
