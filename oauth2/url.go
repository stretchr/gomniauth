package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

// GetBeginAuthURLWithBase returns the OAuth2 authorization URL from the given arguments.
//
// The state object will be encoded to base64 and signed to ensure integrity.
func GetBeginAuthURLWithBase(base string, state *common.State, config *common.Config) (string, error) {

	if config == nil {
		panic("OAuth2Handler: Must have valid Config specified.")
	}

	// copy the config
	params := objx.MSI(
		OAuth2KeyClientID, config.Get(OAuth2KeyClientID).Str(),
		OAuth2KeyRedirectUrl, config.Get(OAuth2KeyRedirectUrl).Str(),
		OAuth2KeyScope, config.Get(OAuth2KeyScope).Str(),
		OAuth2KeyAccessType, config.Get(OAuth2KeyAccessType).Str(),
		OAuth2KeyApprovalPrompt, config.Get(OAuth2KeyApprovalPrompt).Str(),
		OAuth2KeyResponseType, config.Get(OAuth2KeyResponseType).Str())

	if state != nil {

		// set the state
		stateValue, stateErr := state.SignedBase64(common.GetSecurityKey())

		if stateErr != nil {
			return "", stateErr
		}

		params.Set("state", stateValue)

	}

	// generate the query part
	query, queryErr := params.URLQuery()

	if queryErr != nil {
		return "", queryErr
	}

	// put the strings together
	return base + "?" + query, nil
}
