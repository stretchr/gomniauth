package gomniauth

import (
	"errors"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"net/http"
)

// Session provides omniauth functionality for the given ID.
//
// Your application should use one session per request, and build the
// session by calling the WithID method on the Manager.
type Session struct {
	manager            *Manager
	id                 string
	provider           common.Provider
	GetOauth2Transport func(*oauth2.Config) (*oauth2.Transport, error)
}

// NewSession creates a new session with the given Manager and id.
//
// Users would normally not use this function, and instead, call
// WithID on the Manager object.
func NewSession(manager *Manager, id string, provider common.Provider) *Session {
	s := &Session{manager, id, provider, func(config *oauth2.Config) (*oauth2.Transport, error) {
		return &oauth2.Transport{Config: config}, nil
	}}

	return s
}

// Manager gets the Manager assocaited with this session.
func (s *Session) Manager() *Manager {
	return s.manager
}

// ID gets the identifing string that this session will work with.
//
// Normally, this is a session ID, or some other way of identifying
// each user.
func (s *Session) ID() string {
	return s.id
}

// Provider gets the common.Provider that is being used to authenticate
// this session.
func (s *Session) Provider() common.Provider {
	return s.provider
}

// IsAuthenticated gets whether the session is authenticated or not.
func (s *Session) IsAuthenticated() (bool, error) {

	auth, err := s.Manager().AuthStore().GetAuth(s.id)

	if err != nil {
		return false, err
	}

	if auth == nil {
		// no auth
		return false, nil
	}

	// see if the token has expired
	switch s.provider.AuthType() {
	case common.AuthTypeOAuth2:

		token := oauth2.NewTokenFromAuth(auth)
		if token.HasExpired() {
			return false, nil
		}

	}

	return true, nil

}

func (s *Session) AuthenticatedClient() (*http.Client, error) {

	switch s.provider.AuthType() {
	case common.AuthTypeOAuth2:

		// Set the token on transport IF we have one in the AuthStore
		auth, authStoreErr := s.Manager().AuthStore().GetAuth(s.id)

		if authStoreErr != nil {
			return nil, authStoreErr
		}

		// set the token
		if auth == nil {
			return nil, errors.New("gomniauth: AuthenticatedClient: The AuthStore returned a nil Auth, so an AuthenticatedClient cannot be provided.")
		}

		// make a transport
		// make the config
		var config = &oauth2.Config{
			Map: s.provider.Config().Copy(),
		}
		transport := &oauth2.Transport{Config: config}
		transport.Token = oauth2.NewTokenFromAuth(auth)

		return transport.Client(), nil

	}

	return nil, errors.New("gomniauth: AuthenticatedClient: Unsupported common.AuthType: " + s.provider.AuthType().String())

}

// HandleCallback handles the callback (from the third-party authenticator) and completes
// the process of authenticating the user.
func (s *Session) HandleCallback(request *http.Request) error {

	switch s.provider.AuthType() {
	case common.AuthTypeOAuth2:

		// get the code from the request
		code := request.FormValue("code")

		if len(code) == 0 {
			return errors.New("gomniauth: HandleCallback: No code was found.")
		}

		// make the config
		var config = &oauth2.Config{
			Map: s.provider.Config().Copy(),
		}

		transport := &oauth2.Transport{Config: config}

		// get the auth store
		authStore := s.Manager().AuthStore()

		var auth *common.Auth = nil
		if authStore != nil {

			var authStoreErr error

			// Set the token on transport IF we have one in the AuthStore
			auth, authStoreErr = s.Manager().AuthStore().GetAuth(s.id)

			if authStoreErr != nil {
				return authStoreErr
			}

		}

		// set the token
		if auth != nil {
			transport.Token = oauth2.NewTokenFromAuth(auth)
		}

		// perform the exchange
		auth2Token, exchangeErr := transport.Exchange(code)

		if exchangeErr != nil {
			return exchangeErr
		}

		if authStore != nil {

			// Get the token from the transport.Exchange and pass it to
			// the auth store.
			s.Manager().AuthStore().PutAuth(s.id, auth2Token.Auth())

		}

	default:
		return errors.New("gomniauth: HandleCallback: Unsupported common.AuthType: " + s.provider.AuthType().String())
	}

	return nil

}
