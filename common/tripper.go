package common

import (
	"net/http"
)

// Tripper represents an object capable of making authenticated
// round trips.
type Tripper interface {
	http.RoundTripper
	// Credentials gets the authentication credentials that
	// this Tripper will use.
	Credentials() *Credentials

	// Provider gets the Provider that this tripper will make
	// requests to.
	Provider() Provider
}
