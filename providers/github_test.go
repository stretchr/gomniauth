package providers

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGitHubImplementrsProvider(t *testing.T) {

	var provider gomniauth.Provider
	provider = new(GithubProvider)

	assert.NotNil(t, provider)

}

func TestNewGithub(t *testing.T) {

	g := Github("clientID", "secret", "http://myapp.com/")

	if assert.NotNil(t, g) && assert.NotNil(t, g.config) {
		assert.Equal(t, "clientID", g.config.Get(OAuth2KeyClientID))
		assert.Equal(t, "secret", g.config.Get(OAuth2KeySecret))
		assert.Equal(t, "http://myapp.com/", g.config.Get(OAuth2KeyRedirectUrl))
		assert.Equal(t, githubDefaultScope, g.config.Get(OAuth2KeyScope))
	}

}

func TestGithubName(t *testing.T) {
	g := Github("clientID", "secret", "http://myapp.com/")
	assert.Equal(t, githubName, g.Name())
}

func TestGetBeginAuthURL(t *testing.T) {

	gomniauth.SecurityKey = "ABC123"

	state := &common.State{objects.M("after", "http://www.stretchr.com/")}

	g := Github("clientID", "secret", "http://myapp.com/")

	url, err := g.GetBeginAuthURL(state)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=clientID")
		assert.Contains(t, url, "redirect_url=http%3A%2F%2Fmyapp.com%2F")
		assert.Contains(t, url, "scope="+githubDefaultScope)
		assert.Contains(t, url, "access_type="+OAuth2AccessTypeOnline)
		assert.Contains(t, url, "approval_prompt="+OAuth2ApprovalPromptAuto)
	}

}
