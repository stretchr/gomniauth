package providers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGithub(t *testing.T) {

	g := Github("clientId", "clientSecret", "redirectURL", "scope1", "scope2")

	assert.Equal(t, g.Config().GetString("authURL"), "https://github.com/login/oauth/authorize")
	assert.Equal(t, g.Config().GetString("tokenURL"), "https://github.com/login/oauth/access_token")
	assert.Equal(t, g.Config().GetString("clientId"), "clientId")
	assert.Equal(t, g.Config().GetString("clientSecret"), "clientSecret")
	assert.Equal(t, g.Config().GetString("redirectURL"), "redirectURL")
	assert.Equal(t, g.Config().GetString("scope"), "scope1,scope2")

}
