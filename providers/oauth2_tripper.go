package providers

import (
	"github.com/stretchr/gomniauth/common"
	"net/http"
)

type OAuth2Tripper struct {
	UnderlyingTransport http.RoundTripper
	Credentials         *common.Credentials
}

func NewOAuth2Tripper(creds *common.Credentials) *OAuth2Tripper {
	return &OAuth2Tripper{http.DefaultTransport, creds}
}

func (t *OAuth2Tripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// make the round trip
	return t.UnderlyingTransport.RoundTrip(req)
}
