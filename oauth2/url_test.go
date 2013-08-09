package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOAuth2HandlerBeginAuthURLWithBase(t *testing.T) {

	common.SetSecurityKey("rAALj6QhRjsTo3VKzfWuK21qNZ5bFfqPJ9sYNerSYeKKoMIPAi9vaIusjmqyLE3S")

	base := "https://base.url/auth"

	config := &common.Config{objects.M()}
	config.
		Set("client_id", "client_id").
		Set("redirect_url", "redirect_url").
		Set("scope", "scope").
		Set("access_type", "access_type").
		Set("approval_prompt", "approval_prompt")

	state := &common.State{objects.M("after", "http://www.stretchr.com/")}
	base64State, _ := state.Base64()

	url, err := GetBeginAuthURLWithBase(base, state, config)

	if assert.NoError(t, err) {
		assert.Contains(t, url, "client_id=client_id")
		assert.Contains(t, url, "redirect_url=redirect_url")
		assert.Contains(t, url, "scope=scope")
		assert.Contains(t, url, "access_type=access_type")
		assert.Contains(t, url, "approval_prompt=approval_prompt")
		assert.Contains(t, url, "state="+base64State)
	}

}
