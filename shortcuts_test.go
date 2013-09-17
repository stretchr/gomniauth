package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewState(t *testing.T) {

	s := NewState("one", 1, "two", 2)
	if assert.NotNil(t, s) {
		assert.Equal(t, 1, s.Get("one").Data())
		assert.Equal(t, 2, s.Get("two").Data())
	}

}

func TestStateFromParam(t *testing.T) {

	state := NewState("name", "Mat", "age", 30)
	hash, _ := state.SignedBase64(GetSecurityKey())

	s, err := StateFromParam(hash)
	if assert.NotNil(t, s) && assert.NoError(t, err) {
		assert.Equal(t, "Mat", s.Get("name").Data())
		assert.Equal(t, 30, s.Get("age").Data())
	}

}

func TestProvider(t *testing.T) {

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
	returnedProv, err = Provider("no such provider")

	if assert.Nil(t, returnedProv) {
		assert.IsType(t, &common.MissingProviderError{}, err, "MissingProviderError expected")
	}

	mock.AssertExpectationsForObjects(t, prov1.Mock, prov2.Mock)

}
