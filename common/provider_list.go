package common

// ProviderList represents objects capable of managing
// a collection of providers.
type ProviderList interface {

	// Providers gets all registered Provider objects.
	Providers() []Provider

	// Provider gets a provider by name, or returns a MissingProviderError
	// if no provider with that name is registered.
	Provider(name string) (Provider, error)
}
