package providers

import (
	"github.com/stretchr/gomniauth/common"
	"net/http"
)

type OAuth2Tripper struct {
	underlyingTransport http.RoundTripper
	credentials         *common.Credentials
}

func NewOAuth2Tripper(creds *common.Credentials) *OAuth2Tripper {
	return &OAuth2Tripper{http.DefaultTransport, creds}
}

func (t *OAuth2Tripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// make the round trip
	return t.underlyingTransport.RoundTrip(req)
}

func (t *OAuth2Tripper) Credentials() *common.Credentials {
	return t.credentials
}
