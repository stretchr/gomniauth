package github

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUserInterface(t *testing.T) {

	var user common.User = new(User)

	assert.NotNil(t, user)

}

func TestNewUser(t *testing.T) {

	testProvider := new(test.TestProvider)
	testProvider.On("Name").Return("providerName")

	data := objects.M(
		githubKeyID, "123435467890",
		githubKeyName, "Mathew",
		githubKeyEmail, "my@email.com",
		githubKeyNickname, "Mat",
		githubKeyAvatarUrl, "http://myface.com/")
	creds := &common.Credentials{objects.M(oauth2.OAuth2KeyAccessToken, "ABC123")}

	user := NewUser(data, creds, testProvider)

	if assert.NotNil(t, user) {

		assert.Equal(t, data, user.Data())

		assert.Equal(t, "Mathew", user.Name())
		assert.Equal(t, "my@email.com", user.Email())
		assert.Equal(t, "Mat", user.Nickname())
		assert.Equal(t, "http://myface.com/", user.AvatarURL())

		// check provider credentials
		creds := user.ProviderCredentials()[testProvider.Name()]
		if assert.NotNil(t, creds) {
			assert.Equal(t, "ABC123", creds.GetStringOrEmpty(oauth2.OAuth2KeyAccessToken))
			assert.Equal(t, "123435467890", creds.GetStringOrEmpty(common.CredentialsKeyID))
		}

	}

	mock.AssertExpectationsForObjects(t, testProvider.Mock)

}

func TestIDForProvider(t *testing.T) {

	user := new(User)
	user.data = objects.M(
		common.UserKeyProviderCredentials,
		map[string]*common.Credentials{
			"github": &common.Credentials{objects.M(common.UserKeyID, "githubid")},
			"google": &common.Credentials{objects.M(common.UserKeyID, "googleid")}})

	assert.Equal(t, "githubid", user.IDForProvider("github"))
	assert.Equal(t, "googleid", user.IDForProvider("google"))

}
