package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
)

const (
	UserKeyEmail               string = "UserKeyEmail"
	UserKeyName                string = "UserKeyName"
	UserKeyNickname            string = "UserKeyNickname"
	UserKeyAvatar              string = "UserKeyAvatar"
	UserKeyID                  string = "UserKeyID"
	UserKeyProviderCredentials string = "UserKeyProviderCredentials"
)

// User defines a type that will be populated with as much
// data as possible from a provider once authentication is completed.
//
// Not all fields are guaranteed to be populated.
type User struct {
	objects.Map
}

// Email gets the users email address.
func (u User) Email() string {
	return ""
}

// Name gets the users full name.
func (u User) Name() string {
	return ""
}

// Nickname gets the users nickname or username.
func (u User) Nickname() string {
	return ""
}

// AvatarURL gets the URL of an image representing the user.
func (u User) AvatarURL() string {
	return ""
}

// ProviderCredentials gets a map of Credentials (by provider name).
func (u User) ProviderCredentials() map[string]*common.Credentials {
	return nil
}

func (u User) AddProviderCredentials(provider common.Provider, creds *common.Credentials) error {

	providerCreds := u.Get(UserKeyProviderCredentials)

	if providerCreds == nil {
		providerCreds = objects.M()
		u.Set(UserKeyProviderCredentials, providerCreds)
	}

	providerCreds.(objects.Map).Set(provider.Name(), creds)

	return nil
}

// ID gets this user's globally unique ID.
func (u User) ID() string {
	return ""
}

/*
// AuthToken gets the token used to identify a client with
// permission to act on behalf of this user.
//
// You would use AuthToken to load a User object from the
// database.
//
// e.g. from ?auth=213897234
func (u User) AuthToken() string {
	return ""
}
*/
