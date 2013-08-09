package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAddProviderCredentials(t *testing.T) {

	provider := new(test.TestProvider)
	creds := new(common.Credentials)

	provider.On("Name").Return("provider-name")

	user := &User{objects.M()}
	var userInterface common.User = user

	assert.NotNil(t, userInterface)

	user.AddProviderCredentials(provider, creds)

	if assert.NotNil(t, user.Get(UserKeyProviderCredentials)) {
		assert.Equal(t, creds, user.GetMap(UserKeyProviderCredentials).Get("provider-name"))
	}

	mock.AssertExpectationsForObjects(t, provider.Mock)

}
