package gomniauth

import (
	"errors"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"net/http"
)

var (
	// StateKeyID is the key that will be used to store
	// the ID inside the state.
	StateKeyID string = "id"
)

func StateWithID(id string) objects.Map {
	return objects.NewMap(StateKeyID, id)
}

func StateFromRequest(authType common.AuthType, r *http.Request, stateSecurityKey string) (objects.Map, error) {

	switch authType {
	case common.AuthTypeOAuth2:

		state := r.FormValue("state")
		return objects.NewMapFromSignedBase64String(state, stateSecurityKey)

	}

	return nil, errors.New("gomniauth: StateFromRequest: Unsupported common.AuthType: " + authType.String())

}
