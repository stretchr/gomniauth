package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStateFromRequest(t *testing) {

	r, _ := http.NewRequest("GET", "http://www.test.com/?state=obj", "")

	s, err := StateFromRequest(common.AuthTypeOAuth2, r)

	if assert.NoError(t, err) {
		if assert.NotNil(t, s) {
			assert.Equal(t, "Tyler", s.Get("name"))
		}
	}

}

func TestIDFromState(t *testing.T) {

	s := objects.NewMap("id", "abc123")

	id, err := IDFromState(s)

	if assert.NoError(t, err) {
		assert.Equal(t, id, "abc123")
	}

}

func TestTargetURLFromState(t *testing.T) {

	s := objects.NewMap("targetUrl", "http://www.test.com/")

	targetUrl, err := TargetURLFromState(s)

	if assert.NoError(t, err) {
		assert.Equal(t, targetUrl, "http://www.test.com/")
	}

}
