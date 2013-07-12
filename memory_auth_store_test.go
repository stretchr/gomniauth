package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryAuthStore(t *testing.T) {

	store := make(MemoryAuthStore)
	var authStore AuthStore = store

	auth := &common.Auth{}

	// put it
	if assert.NoError(t, authStore.PutAuth("id", auth)) {
		if assert.Equal(t, 1, len(store)) {
			assert.Equal(t, store["id"], auth)
		}
	}

	// get it
	gotAuth, err := authStore.GetAuth("id")
	if assert.NoError(t, err) {
		assert.Equal(t, auth, gotAuth)
	}

	// delete it
	if assert.NoError(t, authStore.DeleteAuth("id")) {
		assert.Equal(t, 0, len(store))
	}

	// get it
	gotAuth, err = authStore.GetAuth("id")
	if assert.NoError(t, err) {
		assert.Nil(t, gotAuth)
	}

}
