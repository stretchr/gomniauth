package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTokenFromAuth(t *testing.T) {

	auth := &common.Auth{objects.NewMap()}
	auth.Set("accessToken", "at")
	auth.Set("refreshToken", "rt")
	auth.Set("expiry", time.Now())

	var token *Token = NewTokenFromAuth(auth)

	assert.Equal(t, "at", token.AccessToken)
	assert.Equal(t, "rt", token.RefreshToken)
	assert.Equal(t, auth.Get("expiry").(time.Time), token.Expiry)

}

func TestToken_Auth(t *testing.T) {

	token := &Token{
		AccessToken:  "",
		RefreshToken: "",
		Expiry:       time.Now(),
	}

	auth := token.Auth()

	assert.Equal(t, auth.GetString("accessToken"), token.AccessToken)
	assert.Equal(t, auth.GetString("refreshToken"), token.RefreshToken)
	assert.Equal(t, auth.Get("expiry").(time.Time), token.Expiry)

}

func TestIsExpired(t *testing.T) {

	// control time
	now := time.Now()
	getCurrentTime = func() time.Time {
		return now
	}

	tok := &Token{
		AccessToken:  "",
		RefreshToken: "",
		Expiry:       now.Add(0 - time.Second),
	}

	assert.True(t, tok.HasExpired())

	tok = &Token{
		AccessToken:  "",
		RefreshToken: "",
		Expiry:       now.Add(time.Second),
	}

	assert.False(t, tok.HasExpired())

	tok = &Token{
		AccessToken:  "",
		RefreshToken: "",
	}

	assert.False(t, tok.HasExpired())

}
