package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"time"
)

// Token contains an end-user's tokens.  This is the data you
// must store to persist authentication.
type Token struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time // If zero the token has no (known) expiry time.
}

func NewTokenFromAuth(auth *common.Auth) *Token {

	t := &Token{
		AccessToken:  auth.GetString("accessToken"),
		RefreshToken: auth.GetString("refreshToken"),
	}

	if timeObj, ok := auth.Get("expiry").(time.Time); ok {
		t.Expiry = timeObj
	}

	return t
}

// Auth creates a common.Auth object from this token.
func (t *Token) Auth() *common.Auth {
	return &common.Auth{objects.M("accessToken", t.AccessToken,
		"refreshToken", t.RefreshToken,
		"expiry", t.Expiry)}
}

// HasExpired gets whether the token has expired or not.
func (t *Token) HasExpired() bool {
	if t.Expiry.IsZero() {
		return false
	}
	return t.Expiry.Before(getCurrentTime())
}

// getCurrentTime gets the current time.
var getCurrentTime func() time.Time = func() time.Time {
	return time.Now()
}
