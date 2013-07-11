package gomniauth

import (
	"errors"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/stew/objects"
	"log"
	"net/http"
)

type Session struct {
	manager *Manager
	id      string
}

func NewSession(manager *Manager, id string) *Session {
	return &Session{manager: manager, id: id}
}

func (s *Session) Manager() *Manager {
	return s.manager
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) IsLoggedIn() bool {
	return false
}

func (s *Session) GetAuthURL(provider Provider, state objects.Map) (string, error) {

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

func (s *Session) HandleCallback(provider Provider, id string, request *http.Request) error {

	switch provider.AuthType() {
	case common.AuthTypeOAuth2:

		// get the code from the request
		code := request.FormValue("code")

		// make the config
		var config = &oauth2.Config{
			Map: provider.Config().Copy(),
		}

		transport := &oauth2.Transport{Config: config}
		token, exchangeErr := transport.Exchange(code)

		if exchangeErr != nil {
			return exchangeErr
		}

		log.Printf("Got the token: %s", token)

	}

	return errors.New("gomniauth: HandleCallback: Unsupported common.AuthType: " + provider.AuthType().String())

}
