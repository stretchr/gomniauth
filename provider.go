package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"net/http"
)

type Provider interface {

	// Name is the unique name for this provider.
	Name() string

	// AuthHandler gets the AuthHandler that handles this provider.
	AuthHandler() AuthHandler

	// BeginAuthURL gets the URL that the client must visit in order
	// to begin the authentication process.
	BeginAuthURL(params objects.Map) (string, error)

	// CompleteAuth takes a map of arguments that are used to
	// complete the authorisation process, completes it, and returns
	// the appropriate common.Credentials.
	CompleteAuth(data objects.Map) (common.Credentials, error)

	// LoadUser uses the specified common.Credentials to access the users profile
	// from the remote provider, and builds the appropriate User object.
	LoadUser(creds common.Credentials) (User, error)

	// GetClient gets an http.Client authenticated with the specified
	// common.Credentials.
	GetClient(creds common.Credentials) (*http.Client, error)
}
