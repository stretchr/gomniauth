package oauth2

import (
	"net/url"
)

// Config is the configuration of an OAuth consumer.
type Config struct {

	// ClientId is the OAuth client identifier used when communicating with
	// the configured OAuth provider.
	ClientId string

	// ClientSecret is the OAuth client secret used when communicating with
	// the configured OAuth provider.
	ClientSecret string

	// Scope identifies the level of access being requested. Multiple scope
	// values should be provided as a space-delimited string.
	Scope string

	// AuthURL is the URL the user will be directed to in order to grant
	// access.
	AuthURL string

	// TokenURL is the URL used to retrieve OAuth tokens.
	TokenURL string

	// RedirectURL is the URL to which the user will be returned after
	// granting (or denying) access.
	RedirectURL string

	// AccessType determins the type of access. (Optional)
	AccessType string

	// ApprovalPrompt indicates whether the user should be
	// re-prompted for consent. If set to ApprovalPromptAuto (default) the
	// user will be prompted only if they haven't previously
	// granted consent and the code can only be exchanged for an
	// access token.
	//
	// If set to ApprovalPromptForce the user will always be prompted, and the
	// code can be exchanged for a refresh token.
	ApprovalPrompt string
}

// AuthCodeURL returns a URL that the end-user should be redirected to,
// so that they may obtain an authorization code.
func (c *Config) GetAuthCodeURL(state string) string {

	url_, err := url.Parse(c.AuthURL)
	if err != nil {
		panic("oauth2: Config.AuthURL error: " + err.Error())
	}
	q := url.Values{
		"response_type":   {"code"},
		"client_id":       {c.ClientId},
		"redirect_uri":    {c.RedirectURL},
		"scope":           {c.Scope},
		"state":           {state},
		"access_type":     {c.AccessType},
		"approval_prompt": {c.ApprovalPrompt},
	}.Encode()
	if url_.RawQuery == "" {
		url_.RawQuery = q
	} else {
		url_.RawQuery += "&" + q
	}
	return url_.String()
}
