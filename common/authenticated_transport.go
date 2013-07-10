package common

import (
	"net/http"
)

type AuthenticatedTransport interface {
	http.RoundTripper

	Client() *http.Client
}
