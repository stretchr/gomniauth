package common

const (
	UserKeyEmail               string = "UserKeyEmail"
	UserKeyName                string = "UserKeyName"
	UserKeyNickname            string = "UserKeyNickname"
	UserKeyAvatar              string = "UserKeyAvatar"
	UserKeyID                  string = "UserKeyID"
	UserKeyProviderCredentials string = "UserKeyProviderCredentials"
)

type User interface {
	// Email gets the users email address.
	Email() string

	// Name gets the users full name.
	Name() string

	// Nickname gets the users nickname or username.
	Nickname() string

	// AvatarURL gets the URL of an image representing the user.
	AvatarURL() string

	// ProviderCredentials gets a map of Credentials (by provider name).
	ProviderCredentials() map[string]*Credentials

	// ID gets this user's globally unique ID.
	ID() string

	// GetValue gets any User field by name.
	GetValue(key string) interface{}

	// AddProviderCredentials adds the provider credentials to the
	// ProviderCredentials for this User.
	AddProviderCredentials(Provider, *Credentials) error
}
