package oauth2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOAuth2ParseScope(t *testing.T) {

	scope := "user,email"
	parsed := ParseScope(scope)
	assert.Equal(t, parsed, "user email")

	scope = "user email"
	parsed = ParseScope(scope)
	assert.Equal(t, parsed, "user email")

	scope = "user"
	parsed = ParseScope(scope)
	assert.Equal(t, parsed, "user")

	scope = "  user  "
	parsed = ParseScope(scope)
	assert.Equal(t, parsed, "user")

	scope = " user   email "
	parsed = ParseScope(scope)
	assert.Equal(t, parsed, "user email")

	scope = "user,   email"
	parsed = ParseScope(scope)
	assert.Equal(t, parsed, "user email")

}

func TestOAuth2MergeScopes(t *testing.T) {

	scope := MergeScopes("user", "email")
	assert.Equal(t, scope, "user email")

	scope = MergeScopes("user,email", "avatar")
	assert.Equal(t, scope, "user email avatar")

	scope = MergeScopes("user email", "avatar")
	assert.Equal(t, scope, "user email avatar")

	scope = MergeScopes("user,  email ", "avatar   ")
	assert.Equal(t, scope, "user email avatar")

	scope = MergeScopes("user", " email ", "avatar   ")
	assert.Equal(t, scope, "user email avatar")

	scope = MergeScopes("", "email")
	assert.Equal(t, scope, "email")

	scope = MergeScopes()
}
