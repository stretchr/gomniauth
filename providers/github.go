package providers

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/handlers"
	"github.com/stretchr/stew/objects"
	"net/url"
)

type Github struct {
	config      common.Config
	authHandler *handlers.OAuth2Handler
}

func (g *Github) AuthHandler() gomniauth.AuthHandler {
	return g.authHandler
}

func (g *Github) BeginAuthURL(params objects.Map) (string, error) {

	// get the state (signed)
	state, stateErr := params.GetMap("state").SignedBase64(gomniauth.GetSecurityKey())

	if stateErr != nil {
		return "", stateErr
	}

	paramList := url.Values{
		"response_type":   {"code"},
		"client_id":       {g.config.GetString("client_id")},
		"redirect_uri":    {g.config.GetString("redirect_url")},
		"scope":           {g.config.GetString("scope")},
		"state":           {state},
		"access_type":     {g.config.GetString("access_type")},
		"approval_prompt": {g.config.GetString("approval_prompt")},
	}.Encode()

	paramStr, err := g.authHandler.GetParamsString(params)

	if err != nil {
		return "", err
	}

	return "https://github.com/login/oauth/authorize?" + paramStr, nil
}
