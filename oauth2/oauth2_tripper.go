package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"net/http"
)

type OAuth2Tripper struct {
	underlyingTransport http.RoundTripper
	credentials         *common.Credentials
	provider            common.Provider
}

// NewOAuth2Tripper creates a new OAuth2Tripper with the given arguments.
func NewOAuth2Tripper(creds *common.Credentials, provider common.Provider) *OAuth2Tripper {
	return &OAuth2Tripper{common.GetRoundTripper(), creds, provider}
}

// RoundTrip is called by the http package when making a request to a server. This
// implementation of RoundTrip inserts the appropriate authorization headers
// into the request in order to enable access of protected resources.
//
// If the auth token has expired, RoundTrip will attempt to renew it.
func (t *OAuth2Tripper) RoundTrip(req *http.Request) (*http.Response, error) {

	//TODO: check token expiration and renew if necessary

	if t.Credentials() != nil {

		// copy the request
		req = cloneRequest(req)

		// set the header
		headerK, headerV := AuthorizationHeader(t.Credentials())
		req.Header.Set(headerK, headerV)

		// set the accept header to ask for JSON
		req.Header.Set("Accept", "application/json")

	}

	return t.underlyingTransport.RoundTrip(req)
}

// Credentials returns the credentials associated with this OAuth2Tripper
func (t *OAuth2Tripper) Credentials() *common.Credentials {
	return t.credentials
}

// Provider returns the provider associated with this OAuth2Tripper
func (t *OAuth2Tripper) Provider() common.Provider {
	return t.provider
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header)
	for k, s := range r.Header {
		r2.Header[k] = s
	}
	return r2
}
