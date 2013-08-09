package common

import (
	"github.com/stretchr/stew/objects"
	"net/http"
)

// Provider represents an authentication provider.
type Provider interface {

	// Name is the unique name for this provider.
	Name() string

	// GetBeginAuthURL gets the URL that the client must visit in order
	// to begin the authentication process.
	GetBeginAuthURL(state *State) (string, error)

	// CompleteAuth takes a map of arguments that are used to
	// complete the authorisation process, completes it, and returns
	// the appropriate Credentials.
	CompleteAuth(data objects.Map) (*Credentials, error)

	// GetUser uses the specified Credentials to access the users profile
	// from the remote provider, and builds the appropriate User object.
	GetUser(creds *Credentials) (User, error)

	// Get makes an authenticated request and returns the data in the
	// response as a data map.
	Get(creds *Credentials, endpoint string) (objects.Map, error)

	// GetClient gets an http.Client authenticated with the specified
	// Credentials.
	GetClient(creds *Credentials) (*http.Client, error)
}
