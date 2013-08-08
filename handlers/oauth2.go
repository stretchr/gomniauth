package handlers

import (
	"github.com/stretchr/gomniauth/common"
	"net/http"
)

type OAuth2Tripper struct {
	UnderlyingTransport http.RoundTripper
	Credentials         common.Credentials
}

func (t *OAuth2Tripper) RoundTrip(req *http.Request) (*http.Response, error) {

	// use the default transport if none is specified
	if t.UnderlyingTransport == nil {
		t.UnderlyingTransport = http.DefaultTransport
	}

	// make the round trip
	return t.UnderlyingTransport.RoundTrip(req)
}

type OAuth2Handler struct{}

var DefaultOAuth2Handler = &OAuth2Handler{}

func (h *OAuth2Handler) NewRoundTripper() (http.RoundTripper, error) {
	return new(OAuth2Tripper), nil
}
