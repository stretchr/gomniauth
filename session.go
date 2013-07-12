package gomniauth

import (
	"errors"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/stew/objects"
	"net/http"
)

// Session provides omniauth functionality for the given ID.
//
// Your application should use one session per request, and build the
// session by calling the WithID method on the Manager.
type Session struct {
	manager *Manager
	id      string
}

// NewSession creates a new session with the given Manager and id.
//
// Users would normally not use this function, and instead, call
// WithID on the Manager object.
func NewSession(manager *Manager, id string) *Session {
	return &Session{manager: manager, id: id}
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

// GetAuthURL uses the specified provider and state objects to build the
// URL which the user must be redirected to in order to get authenticated.
func (s *Session) GetAuthURL(provider common.Provider, state objects.Map) (string, error) {

	switch provider.AuthType() {
	case common.AuthTypeOAuth2:

		// make the config
		var config = &oauth2.Config{
			Map: provider.Config().Copy(),
		}

		encodedState, encodedStateErr := state.Base64()

		if encodedStateErr != nil {
			return "", encodedStateErr
		}

		return config.AuthCodeURL(encodedState), nil

	}

	return "", errors.New("gomniauth: GetAuthURL: Unsupported common.AuthType: " + provider.AuthType().String())

}

// HandleCallback handles the callback (from the third-party authenticator) and completes
// the process of authenticating the user.
func (s *Session) HandleCallback(provider common.Provider, id string, request *http.Request) error {

	switch provider.AuthType() {
	case common.AuthTypeOAuth2:

		// get the code from the request
		code := request.FormValue("code")

		// make the config
		var config = &oauth2.Config{
			Map: provider.Config().Copy(),
		}

		transport := &oauth2.Transport{Config: config}

		// #cache: Set the token on transport IF we have one in the AuthStore

		_, exchangeErr := transport.Exchange(code)

		if exchangeErr != nil {
			return exchangeErr
		}

		// #cache: TODO: get the token from the transport.Exchange and pass it to
		// the auth store.

	}

	return errors.New("gomniauth: HandleCallback: Unsupported common.AuthType: " + provider.AuthType().String())

}
