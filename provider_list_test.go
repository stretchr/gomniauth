package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestInterface(t *testing.T) {

	var list common.ProviderList
	list = new(ProviderList)

	assert.NotNil(t, list)

}

func TestWithProviders(t *testing.T) {

	common.SetSecurityKey("ABC123")

	prov1 := new(test.TestProvider)
	prov2 := new(test.TestProvider)

	list := WithProviders(prov1, prov2)

	if assert.NotNil(t, list) {

		if assert.Equal(t, 2, len(list.providers)) {
			assert.Equal(t, prov1, list.providers[0])
			assert.Equal(t, prov2, list.providers[1])
		}

		// make sure the SharedProviderList was assigned too
		assert.Equal(t, SharedProviderList, list)

	}

}

func TestProviders(t *testing.T) {

	prov1 := new(test.TestProvider)
	prov2 := new(test.TestProvider)

	list := WithProviders(prov1, prov2)

	if assert.Equal(t, 2, len(list.Providers())) {
		assert.Equal(t, prov1, list.Providers()[0])
		assert.Equal(t, prov2, list.Providers()[1])
	}

}

func TestProviderListProvider(t *testing.T) {

	prov1 := new(test.TestProvider)
	prov2 := new(test.TestProvider)

	prov1.On("Name").Return("prov1")
	prov2.On("Name").Return("prov2")

	// build a list
	list := WithProviders(prov1, prov2)

	returnedProv, err := list.Provider("prov1")

	if assert.NoError(t, err) {
		assert.Equal(t, returnedProv, prov1)
	}

	// check nonsense name
	returnedProv, err = list.Provider("no such provider")

	if assert.Nil(t, returnedProv) {
		assert.IsType(t, &common.MissingProviderError{}, err, "MissingProviderError expected")
	}

	mock.AssertExpectationsForObjects(t, prov1.Mock, prov2.Mock)

}

func TestAdd(t *testing.T) {

	prov1 := new(test.TestProvider)
	prov2 := new(test.TestProvider)
	prov3 := new(test.TestProvider)

	// build a list
	list := WithProviders(prov1, prov2)

	if assert.Equal(t, 2, len(list.providers)) {

		// add prov3
		assert.Equal(t, list, list.Add(prov3), "Add should chain")
		assert.Equal(t, 3, len(list.providers), "Add should add the provider")
		assert.Equal(t, prov3, list.providers[2])

	}

}
