package gomniauth

type DataService interface {

	// GetUserByAuthToken retrieves a user object from the data store by
	// the provided auth token.
	//
	// If no user exists for the specified authToken, a nil User object will
	// be returned.
	//
	// If an empty authToken is specified, a nil User object is always returned,
	// but this is not an error.
	GetUserByAuthToken(authToken string) (User, error)

	// GetUserByProviderId retrieves a user object from the data store by
	// the provided provider and providerId.
	//
	// This will be called after a successful provider negotiation in order
	// to retrieve an already existing user object from the data store.
	//
	// If a user object is not found, a nil User will be returned.
	// This is an entirely new user and a new User object should be created,
	// then stored by calling PutUser.
	GetUserByProviderId(provider, providerId string) (User, error)

	// PutUser persists the specified User.
	PutUser(user User) error
}
