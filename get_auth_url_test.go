package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestManager_GetAuthURL(t *testing.T) {

	id := "abc123"
	targetUrl := "http://www.google.com/"
	prov := new(common.TestProvider)

	state := objects.NewMap("id", id, "targetUrl", targetUrl)

	provider := new(common.TestProvider)
	provider.On("Config").Return(objects.NewMap("clientId", "CLIENTID",
		"redirectURL", "http://www.test.com/",
		"accessType", "online",
		"approvalPrompt", "force"))

	provider.On("AuthType").Return(common.AuthTypeOAuth2)

	key := "security-key"
	url, err := GetAuthURL(provider, state, key)

	if assert.NoError(t, err) {
		assert.Equal(t, "?access_type=online&approval_prompt=force&client_id=CLIENTID&redirect_uri=http%3A%2F%2Fwww.test.com%2F&response_type=code&scope=&state=eyJpZCI6ImFiYzEyMyIsInRhcmdldFVybCI6Imh0dHA6Ly93d3cuZ29vZ2xlLmNvbS8ifQ%3D%3D_72b967bc068ee9e48f0d1d0779924de446377be1", url)
	}

	mock.AssertExpectationsForObjects(t, provider.Mock, prov.Mock)

}
