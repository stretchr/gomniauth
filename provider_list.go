package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
)

// SharedProviderList keeps track of the last created ProviderList
// useful for using shortcut methods directly on gomniauth package
// rather than having to refer to the list.
var SharedProviderList common.ProviderList

// ProviderList represents a simple common.ProviderList that holds
// an array of providers, and allows access to them.
type ProviderList struct {
	providers []common.Provider
}

// WithProviders generates a new ProviderList which should be
// used to interact with Gomniauth services.
func WithProviders(providers ...common.Provider) *ProviderList {
	list := &ProviderList{providers}
	SharedProviderList = list
	return list
}

// Provider gets a provider by name, or returns a common.MissingProviderError
// if no provider with that name is registered.
func (l *ProviderList) Provider(name string) (common.Provider, error) {

	// panic on nil
	if l == nil {
		panic(common.PrefixForErrors + "No providers have been initialised.  Make sure you have called gomniauth.WithProviders(...).")
	}

	for _, provider := range l.providers {
		if provider.Name() == name {
			return provider, nil
		}
	}

	return nil, &common.MissingProviderError{name}
}

// Providers gets all registered Provider objects.
func (l *ProviderList) Providers() []common.Provider {
	return l.providers
}

// Provider gets a provider by name from the
// SharedProviderList, or returns a common.MissingProviderError
// if no provider with that name is registered.
func Provider(name string) (common.Provider, error) {
	return SharedProviderList.Provider(name)
}
