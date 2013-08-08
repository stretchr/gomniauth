package providers

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
)

const (
	OAuth2KeyClientID       string = "client_id"
	OAuth2KeySecret         string = "secret"
	OAuth2KeyRedirectUrl    string = "redirect_url"
	OAuth2KeyScope          string = "scope"
	OAuth2KeyAccessType     string = "access_type"
	OAuth2KeyApprovalPrompt string = "approval_prompt"
)

const (

	// ApprovalPromptForce indicates that the user will always
	// have to reauthorize access if the AccessType is online.
	OAuth2ApprovalPromptForce string = "force"

	// ApprovalPromptAuto indicates that the user will not
	// have to reauthorize access.
	OAuth2ApprovalPromptAuto string = "auto"
)

const (

	// AccessTypeOnline indicates that the access type is online.
	OAuth2AccessTypeOnline string = "online"

	// AccessTypeOffline indicates that the access type is offline.
	OAuth2AccessTypeOffline string = "offline"
)

type OAuth2Provider struct {
	Provider
}

func (p *OAuth2Provider) GetBeginAuthURLWithBase(base string, state *common.State, config *common.Config) (string, error) {

	if config == nil {
		panic("OAuth2Handler: Must have valid Config specified.")
	}

	// copy the config
	params := objects.M(
		OAuth2KeyClientID, config.GetStringOrEmpty(OAuth2KeyClientID),
		OAuth2KeyRedirectUrl, config.GetStringOrEmpty(OAuth2KeyRedirectUrl),
		OAuth2KeyScope, config.GetStringOrEmpty(OAuth2KeyScope),
		OAuth2KeyAccessType, config.GetStringOrEmpty(OAuth2KeyAccessType),
		OAuth2KeyApprovalPrompt, config.GetStringOrEmpty(OAuth2KeyApprovalPrompt))

	// set the state
	stateValue, stateErr := state.SignedBase64(gomniauth.GetSecurityKey())

	if stateErr != nil {
		return "", stateErr
	}

	params.Set("state", stateValue)

	// generate the query part
	query, queryErr := params.URLQuery()

	if queryErr != nil {
		return "", queryErr
	}

	// put the strings together
	return base + "?" + query, nil
}
