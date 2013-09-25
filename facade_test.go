package gomniauth

import (
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProviderPublicData(t *testing.T) {

	provider := new(test.TestProvider)

	provider.On("Name").Return("TestName")
	provider.On("DisplayName").Return("TestDisplayName")

	publicData, _ := ProviderPublicData(provider, objx.MSI("loginpathFormat", "~auth/%s/login"))
	publicDataMap := publicData.(map[string]interface{})

	assert.Equal(t, publicDataMap["name"], "TestName")
	assert.Equal(t, publicDataMap["display"], "TestDisplayName")
	assert.Equal(t, publicDataMap["loginpath"], "~auth/TestName/login")

}
