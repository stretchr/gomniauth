package github

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
)

const (
	githubKeyID        string = "id"
	githubKeyName      string = "name"
	githubKeyNickname  string = "login"
	githubKeyEmail     string = "email"
	githubKeyAvatarUrl string = "avatar_url"
)

type User struct {
	data objects.Map
}

func NewUser(data objects.Map, creds *common.Credentials, provider common.Provider) *User {
	user := &User{data}

	creds.Set(common.CredentialsKeyID, data[githubKeyID])

	// set provider credentials
	user.data[common.UserKeyProviderCredentials] = map[string]*common.Credentials{
		provider.Name(): creds,
	}

	return user
}

// Email gets the users email address.
func (u *User) Email() string {
	return u.Data().GetStringOrEmpty(githubKeyEmail)
}

// Name gets the users full name.
func (u *User) Name() string {
	return u.Data().GetStringOrEmpty(githubKeyName)

}

// Nickname gets the users nickname or username.
func (u *User) Nickname() string {
	return u.Data().GetStringOrEmpty(githubKeyNickname)

}

// AvatarURL gets the URL of an image representing the user.
func (u *User) AvatarURL() string {
	return u.Data().GetStringOrEmpty(githubKeyAvatarUrl)
}

// ProviderCredentials gets a map of Credentials (by provider name).
func (u *User) ProviderCredentials() map[string]*common.Credentials {
	return u.Data().Get(common.UserKeyProviderCredentials).(map[string]*common.Credentials)
}

// ID gets this user's globally unique ID.
func (u *User) ID() string {
	return u.Data().GetStringOrEmpty(common.UserKeyID)
}

// GetValue gets any User field by name.
func (u *User) Data() objects.Map {
	return u.data
}
