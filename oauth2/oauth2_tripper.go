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

func NewOAuth2Tripper(creds *common.Credentials, provider common.Provider) *OAuth2Tripper {
	return &OAuth2Tripper{http.DefaultTransport, creds, provider}
}

func (t *OAuth2Tripper) RoundTrip(req *http.Request) (*http.Response, error) {

	if t.Credentials() != nil {

		// copy the request
		req = cloneRequest(req)

		// set the header
		headerK, headerV := AuthorizationHeader(t.Credentials())
		req.Header.Set(headerK, headerV)

	}

	return t.underlyingTransport.RoundTrip(req)
}

func (t *OAuth2Tripper) Credentials() *common.Credentials {
	return t.credentials
}

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
