package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

// Provider gets a provider by name from the
// SharedProviderList, or returns a common.MissingProviderError
// if no provider with that name is registered.
func Provider(name string) (common.Provider, error) {
	return SharedProviderList.Provider(name)
}

// NewState creates a new object that can be used to persist
// state across authentication requests.
func NewState(keyAndValuePairs ...interface{}) *common.State {
	return common.NewState(keyAndValuePairs...)
}

// StateFromParam decodes the state parameter hash and turns it
// into a usable State object.
func StateFromParam(paramValue string) (*common.State, error) {

	stateMap, err := objx.FromSignedBase64(paramValue, GetSecurityKey())

	if err != nil {
		return nil, err
	}

	return &common.State{Map: stateMap}, nil
}
