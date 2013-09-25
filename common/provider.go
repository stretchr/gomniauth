package common

import (
	"github.com/stretchr/codecs"
	"github.com/stretchr/objx"
	"net/http"
)

// Provider represents an authentication provider.
type Provider interface {
	codecs.Facade

	// Name is the unique name for this provider.
	Name() string

	// DisplayName is a human readable name for the provider.
	DisplayName() string

	// GetBeginAuthURL gets the URL that the client must visit in order
	// to begin the authentication process.
	//
	// The state argument contains anything you wish to have sent back to your
	// callback endpoint.
	// The options argument takes any options used to configure the auth request
	// sent to the provider.
	GetBeginAuthURL(state *State, options objx.Map) (string, error)

	// CompleteAuth takes a map of arguments that are used to
	// complete the authorisation process, completes it, and returns
	// the appropriate Credentials.
	CompleteAuth(data objx.Map) (*Credentials, error)

	// GetUser uses the specified Credentials to access the users profile
	// from the remote provider, and builds the appropriate User object.
	GetUser(creds *Credentials) (User, error)

	// Get makes an authenticated request and returns the data in the
	// response as a data map.
	Get(creds *Credentials, endpoint string) (objx.Map, error)

	// GetClient gets an http.Client authenticated with the specified
	// Credentials.
	GetClient(creds *Credentials) (*http.Client, error)
}
