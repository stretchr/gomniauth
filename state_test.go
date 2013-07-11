package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStateFromRequest(t *testing.T) {

	r, _ := http.NewRequest("GET", "http://www.test.com/?state=eyJpZCI6ImFiYzEyMyIsInRhcmdldFVybCI6Imh0dHA6Ly93d3cuZ29vZ2xlLmNvbS8ifQ%3D%3D", nil)

	s, err := StateFromRequest(common.AuthTypeOAuth2, r)

	if assert.NoError(t, err) {
		if assert.NotNil(t, s) {
			assert.Equal(t, "abc123", s.Get("id"))
			assert.Equal(t, "http://www.google.com/", s.Get("targetUrl"))
		}
	}

}

func TestIDFromState(t *testing.T) {

	s := objects.NewMap("id", "abc123")

	id, err := IDFromState(common.AuthTypeOAuth2, s)

	if assert.NoError(t, err) {
		assert.Equal(t, id, "abc123")
	}

}

func TestTargetURLFromState(t *testing.T) {

	s := objects.NewMap("targetUrl", "http://www.test.com/")

	targetUrl := TargetURLFromState(common.AuthTypeOAuth2, s)

	assert.Equal(t, targetUrl, "http://www.test.com/")

}
