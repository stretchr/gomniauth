package providers

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGitHubImplementrsProvider(t *testing.T) {

	var provider common.Provider
	provider = new(GithubProvider)

	assert.NotNil(t, provider)

}

func TestNewGithub(t *testing.T) {

	g := Github("clientID", "secret", "http://myapp.com/")

	if assert.NotNil(t, g) {

		// check config
		if assert.NotNil(t, g.Config) {

			assert.Equal(t, "clientID", g.Config.Get(gomniauth.OAuth2KeyClientID))
			assert.Equal(t, "secret", g.Config.Get(gomniauth.OAuth2KeySecret))
			assert.Equal(t, "http://myapp.com/", g.Config.Get(gomniauth.OAuth2KeyRedirectUrl))
			assert.Equal(t, githubDefaultScope, g.Config.Get(gomniauth.OAuth2KeyScope))

			assert.Equal(t, githubAuthURL, g.Config.Get(gomniauth.OAuth2KeyAuthURL))
			assert.Equal(t, githubTokenURL, g.Config.Get(gomniauth.OAuth2KeyTokenURL))

		}

		// check factory

	}

}

func TestGithubName(t *testing.T) {
	g := Github("clientID", "secret", "http://myapp.com/")
	assert.Equal(t, githubName, g.Name())
}

func TestGitHubGetBeginAuthURL(t *testing.T) {

	gomniauth.SecurityKey = "ABC123"

	state := &common.State{objects.M("after", "http://www.stretchr.com/")}

	g := Github("clientID", "secret", "http://myapp.com/")

	url, err := g.GetBeginAuthURL(state)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=clientID")
		assert.Contains(t, url, "redirect_url=http%3A%2F%2Fmyapp.com%2F")
		assert.Contains(t, url, "scope="+githubDefaultScope)
		assert.Contains(t, url, "access_type="+gomniauth.OAuth2AccessTypeOnline)
		assert.Contains(t, url, "approval_prompt="+gomniauth.OAuth2ApprovalPromptAuto)
	}

}
