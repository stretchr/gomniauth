package oauth2

import (
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

var config = &Config{
	Map: objects.M(
		"clientId", "id",
		"clientSecret", "secret",
		"scope", "https://test.com/auth/buzz",
		"authURL", "https://test.com/oauth2/auth",
		"tokenURL", "https://test.com/oauth2/token",
		"redirectURL", "http://you.example.org/handler",
		"accessType", AccessTypeOnline,
		"approvalPrompt", ApprovalPromptAuto),
}

func TestNewConfig(t *testing.T) {

	if assert.NotNil(t, config) {

		assert.Equal(t, config.ClientId(), "id")
		assert.Equal(t, config.ClientSecret(), "secret")
		assert.Equal(t, config.Scope(), "https://test.com/auth/buzz")
		assert.Equal(t, config.AuthURL(), "https://test.com/oauth2/auth")
		assert.Equal(t, config.TokenURL(), "https://test.com/oauth2/token")
		assert.Equal(t, config.RedirectURL(), "http://you.example.org/handler")
		assert.Equal(t, config.AccessType(), AccessTypeOnline)
		assert.Equal(t, config.ApprovalPrompt(), ApprovalPromptAuto)

	}

}

func TestGetAuthCodeURL(t *testing.T) {

	state := "This is the state"
	authCodeUrl := config.AuthCodeURL(state)

	assert.Equal(t, authCodeUrl, "https://test.com/oauth2/auth?access_type=online&approval_prompt=auto&client_id=id&redirect_uri=http%3A%2F%2Fyou.example.org%2Fhandler&response_type=code&scope=https%3A%2F%2Ftest.com%2Fauth%2Fbuzz&state=This+is+the+state")

}
