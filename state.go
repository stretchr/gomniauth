package gomniauth

import (
	"errors"
	"fmt"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"net/http"
)

func StateFromRequest(authType common.AuthType, r *http.Request) (objects.Map, error) {

	switch authType {
	case common.AuthTypeOAuth2:

		state := r.FormValue("state")
		return objects.NewMapFromBase64(state), nil

	}

	return nil, errors.New("gomniauth: StateFromRequest: Unsupported common.AuthType: " + authType.String())

}

func fieldFromState(authType common.AuthType, state objects.Map, field string) (string, error) {
	obj := state.Get(field)
	if obj == nil {
		return "", errors.New("gomniauth: Cannot find field \"" + field + "\" in state: " + fmt.Sprintf("%s", state))
	}
	return obj.(string), nil
}

func IDFromState(authType common.AuthType, state objects.Map) (string, error) {
	return fieldFromState(authType, state, "id")
}
func TargetURLFromState(authType common.AuthType, state objects.Map) (string, error) {
	return fieldFromState(authType, state, "targetUrl")
}
