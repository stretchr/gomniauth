package providers

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewOAuth2Tripper(t *testing.T) {

	creds := &common.Credentials{objects.M()}
	var tripper gomniauth.Tripper = NewOAuth2Tripper(creds)

	if assert.NotNil(t, tripper) {
		assert.Equal(t, creds, tripper.Credentials())
		assert.Equal(t, http.DefaultTransport, tripper.(*OAuth2Tripper).underlyingTransport)
	}

}
