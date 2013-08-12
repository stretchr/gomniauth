package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
)

// GetBeginAuthURLWithBase returns the OAuth2 authorization URL from the given arguments.
//
// The state object will be encoded to base64 and signed to ensure integrity.
func GetBeginAuthURLWithBase(base string, state *common.State, config *common.Config) (string, error) {

	if config == nil {
		panic("OAuth2Handler: Must have valid Config specified.")
	}

	// copy the config
	params := objects.M(
		OAuth2KeyClientID, config.GetStringOrEmpty(OAuth2KeyClientID),
		OAuth2KeyRedirectUrl, config.GetStringOrEmpty(OAuth2KeyRedirectUrl),
		OAuth2KeyScope, config.GetStringOrEmpty(OAuth2KeyScope),
		OAuth2KeyAccessType, config.GetStringOrEmpty(OAuth2KeyAccessType),
		OAuth2KeyApprovalPrompt, config.GetStringOrEmpty(OAuth2KeyApprovalPrompt),
		OAuth2KeyResponseType, config.GetStringOrEmpty(OAuth2KeyResponseType))

	// set the state
	stateValue, stateErr := state.SignedBase64(common.GetSecurityKey())

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
