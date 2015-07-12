package spotify

import (
	"testing"

	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/oauth2"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserInterface(t *testing.T) {

	var user common.User = new(User)

	assert.NotNil(t, user)

}

func TestNewUser(t *testing.T) {
	testProvider := new(test.TestProvider)
	testProvider.On("Name").Return("providerName")

	data := objx.MSI(
		spotifyKeyID, "123435467890",
		spotifyKeyName, "Mathew",
		spotifyKeyEmail, "my@email.com",
		spotifyKeyNickname, "Mat",
	)

	creds := &common.Credentials{Map: objx.MSI(oauth2.OAuth2KeyAccessToken, "ABC123")}

	user := NewUser(data, creds, testProvider)

	if assert.NotNil(t, user) {

		assert.Equal(t, data, user.Data())

		assert.Equal(t, "Mathew", user.Name())
		assert.Equal(t, "my@email.com", user.Email())
		assert.Equal(t, "Mat", user.Nickname())

		// check provider credentials
		creds := user.ProviderCredentials()[testProvider.Name()]
		if assert.NotNil(t, creds) {
			assert.Equal(t, "ABC123", creds.Get(oauth2.OAuth2KeyAccessToken).Str())
			assert.Equal(t, "123435467890", creds.Get(common.CredentialsKeyID).Str())
		}

	}

	mock.AssertExpectationsForObjects(t, testProvider.Mock)
}

func TestIDForProvider(t *testing.T) {
	user := new(User)
	user.data = objx.MSI(
		common.UserKeyProviderCredentials,
		map[string]*common.Credentials{
			"github": &common.Credentials{Map: objx.MSI(common.CredentialsKeyID, "githubid")},
			"google": &common.Credentials{Map: objx.MSI(common.CredentialsKeyID, "googleid")}})

	assert.Equal(t, "githubid", user.IDForProvider("github"))
	assert.Equal(t, "googleid", user.IDForProvider("google"))
}
