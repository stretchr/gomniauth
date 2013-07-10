package oauth2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var config = &Config{
	ClientId:       "id",
	ClientSecret:   "secret",
	Scope:          "https://test.com/auth/buzz",
	AuthURL:        "https://test.com/oauth2/auth",
	TokenURL:       "https://test.com/oauth2/token",
	RedirectURL:    "http://you.example.org/handler",
	AccessType:     AccessTypeOnline,
	ApprovalPrompt: ApprovalPromptAuto,
}

func TestNewConfig(t *testing.T) {

	assert.NotNil(t, config)

}

func TestGetAuthCodeURL(t *testing.T) {

	state := "This is the state"
	authCodeUrl := config.GetAuthCodeURL(state)

	assert.Equal(t, authCodeUrl, "https://test.com/oauth2/auth?access_type=online&approval_prompt=auto&client_id=id&redirect_uri=http%3A%2F%2Fyou.example.org%2Fhandler&response_type=code&scope=https%3A%2F%2Ftest.com%2Fauth%2Fbuzz&state=This+is+the+state")

}
