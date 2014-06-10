package oauth2

import (
	"testing"

	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
)

func TestOAuth2HandlerBeginAuthURLWithBase(t *testing.T) {

	common.SetSecurityKey("rAALj6QhRjsTo3VKzfWuK21qNZ5bFfqPJ9sYNerSYeKKoMIPAi9vaIusjmqyLE3S")

	base := "https://base.url/auth"

	config := &common.Config{Map: objx.MSI()}
	config.
		Set("client_id", "client_id").
		Set("redirect_uri", "redirect_uri").
		Set("scope", "scope").
		Set("access_type", "access_type").
		Set("approval_prompt", "approval_prompt")

	state := &common.State{Map: objx.MSI("after", "http://www.stretchr.com/")}
	base64State, _ := state.Base64()

	url, err := GetBeginAuthURLWithBase(base, state, config)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=client_id")
		assert.Contains(t, url, "redirect_uri=redirect_uri")
		assert.Contains(t, url, "scope=scope")
		assert.Contains(t, url, "access_type=access_type")
		assert.Contains(t, url, "approval_prompt=approval_prompt")
		assert.Contains(t, url, "state="+base64State)
	}

}

func TestOAuth2HandlerBeginAuthURLWithBaseMultipleScope(t *testing.T) {

	common.SetSecurityKey("rAALj6QhRjsTo3VKzfWuK21qNZ5bFfqPJ9sYNerSYeKKoMIPAi9vaIusjmqyLE3S")

	base := "https://base.url/auth"

	config := &common.Config{Map: objx.MSI()}
	config.
		Set("client_id", "client_id").
		Set("redirect_uri", "redirect_uri").
		Set("scope", "scope1 scope2").
		Set("access_type", "access_type").
		Set("approval_prompt", "approval_prompt")

	state := &common.State{Map: objx.MSI("after", "http://www.stretchr.com/")}
	base64State, _ := state.Base64()

	url, err := GetBeginAuthURLWithBase(base, state, config)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=client_id")
		assert.Contains(t, url, "redirect_uri=redirect_uri")
		assert.Contains(t, url, "scope=scope1+scope2")
		assert.Contains(t, url, "access_type=access_type")
		assert.Contains(t, url, "approval_prompt=approval_prompt")
		assert.Contains(t, url, "state="+base64State)
	}

}

func TestOAuth2HandlerBeginAuthURLWithBaseWithoutState(t *testing.T) {

	common.SetSecurityKey("rAALj6QhRjsTo3VKzfWuK21qNZ5bFfqPJ9sYNerSYeKKoMIPAi9vaIusjmqyLE3S")

	base := "https://base.url/auth"

	config := &common.Config{Map: objx.MSI()}
	config.
		Set("client_id", "client_id").
		Set("redirect_uri", "redirect_uri").
		Set("scope", "scope").
		Set("access_type", "access_type").
		Set("approval_prompt", "approval_prompt")

	url, err := GetBeginAuthURLWithBase(base, nil, config)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=client_id")
		assert.Contains(t, url, "redirect_uri=redirect_uri")
		assert.Contains(t, url, "scope=scope")
		assert.Contains(t, url, "access_type=access_type")
		assert.Contains(t, url, "approval_prompt=approval_prompt")
		assert.NotContains(t, url, "state=")
	}

}
