package oauth2

import (
	"github.com/stretchr/stew/objects"
	"net/url"
)

// Config represents the an OAuth2 configuration.
//
// Make a new config like this:
//
//    var config = &Config{
//    	Map: objects.NewMap(
//    		"clientId", "id",
//    		"clientSecret", "secret",
//    		"scope", "https://test.com/auth/buzz",
//    		"authURL", "https://test.com/oauth2/auth",
//    		"tokenURL", "https://test.com/oauth2/token",
//    		"redirectURL", "http://you.example.org/handler",
//    		"accessType", AccessTypeOnline,
//    		"approvalPrompt", ApprovalPromptAuto),
//    }
type Config struct {
	objects.Map
}

// ClientId is the OAuth client identifier used when communicating with
// the configured OAuth provider.
func (c *Config) ClientId() string {
	return c.GetStringOrDefault("clientId", "")
}

// ClientSecret is the OAuth client secret used when communicating with
// the configured OAuth provider.
func (c *Config) ClientSecret() string {
	return c.GetStringOrDefault("clientSecret", "")
}

// Scope identifies the level of access being requested. Multiple scope
// values should be provided as a space-delimited string.
func (c *Config) Scope() string {
	return c.GetStringOrDefault("scope", "")
}

// AuthURL is the URL the user will be directed to in order to grant
// access.
func (c *Config) AuthURL() string {
	return c.GetStringOrDefault("authURL", "")
}

// TokenURL is the URL used to retrieve OAuth tokens.
func (c *Config) TokenURL() string {
	return c.GetStringOrDefault("tokenURL", "")
}

// RedirectURL is the URL to which the user will be returned after
// granting (or denying) access.
func (c *Config) RedirectURL() string {
	return c.GetStringOrDefault("redirectURL", "")
}

// AccessType determins the type of access. (Optional)
func (c *Config) AccessType() string {
	return c.GetStringOrDefault("accessType", "")
}

// ApprovalPrompt indicates whether the user should be
// re-prompted for consent. If set to ApprovalPromptAuto (default) the
// user will be prompted only if they haven't previously
// granted consent and the code can only be exchanged for an
// access token.
//
// If set to ApprovalPromptForce the user will always be prompted, and the
// code can be exchanged for a refresh token.
func (c *Config) ApprovalPrompt() string {
	return c.GetStringOrDefault("approvalPrompt", "")
}

// AuthCodeURL returns a URL that the end-user should be redirected to,
// so that they may obtain an authorization code.
func (c *Config) AuthCodeURL(state string) string {

	url_, err := url.Parse(c.AuthURL())
	if err != nil {
		panic("oauth2: Config.AuthURL error: " + err.Error())
	}
	q := url.Values{
		"response_type":   {"code"},
		"client_id":       {c.ClientId()},
		"redirect_uri":    {c.RedirectURL()},
		"scope":           {c.Scope()},
		"state":           {state},
		"access_type":     {c.AccessType()},
		"approval_prompt": {c.ApprovalPrompt()},
	}.Encode()
	if url_.RawQuery == "" {
		url_.RawQuery = q
	} else {
		url_.RawQuery += "&" + q
	}
	return url_.String()
}
