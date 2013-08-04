package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const securitykey = "securitykey"

func TestSignedState(t *testing.T) {

	state := objects.M("id", "abc123", "targetUrl", "http://www.google.com/")
	signed, err := state.SignedBase64(securitykey)

	if assert.NoError(t, err) {
		assert.Equal(t, "eyJpZCI6ImFiYzEyMyIsInRhcmdldFVybCI6Imh0dHA6Ly93d3cuZ29vZ2xlLmNvbS8ifQ==_e04e352ec9b2d57faf07a1c786e690bddfaa7493", signed)
	}

}

func TestStateFromRequest(t *testing.T) {

	r, _ := http.NewRequest("GET", "http://www.test.com/?state=eyJpZCI6ImFiYzEyMyIsInRhcmdldFVybCI6Imh0dHA6Ly93d3cuZ29vZ2xlLmNvbS8ifQ==_e04e352ec9b2d57faf07a1c786e690bddfaa7493", nil)

	s, err := StateFromRequest(common.AuthTypeOAuth2, r, securitykey)

	if assert.NoError(t, err) {
		if assert.NotNil(t, s) {
			assert.Equal(t, "abc123", s.Get("id"))
			assert.Equal(t, "http://www.google.com/", s.Get("targetUrl"))
		}
	}

}

func TestStateWithID(t *testing.T) {

	state := StateWithID("id")
	if assert.NotNil(t, state) {
		assert.Equal(t, "id", state["id"])
	}

}
