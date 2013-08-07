package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
)

type User objects.Map

func (u User) Email() string {
	return ""
}
func (u User) FullName() string {
	return ""
}
func (u User) ProviderCredentials() map[string]common.Credentials {
	return nil
}
func (u User) ID() string {
	return ""
}
func (u User) AvatarURL() string {
	return ""
}
func (u User) AuthToken() string {
	return ""
}
