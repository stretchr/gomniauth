package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorizationHeader(t *testing.T) {

	creds := &common.Credentials{Map: objx.MSI()}
	accessTokenVal := "ACCESSTOKEN"
	creds.Set(OAuth2KeyAccessToken, accessTokenVal)
	k, v := AuthorizationHeader(creds)

	assert.Equal(t, k, "Authorization")
	assert.Equal(t, "Bearer "+accessTokenVal, v)

}
