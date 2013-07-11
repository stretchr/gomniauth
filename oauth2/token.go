package oauth2

import (
	"time"
)

// Token contains an end-user's tokens.  This is the data you
// must store to persist authentication.
type Token struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time // If zero the token has no (known) expiry time.
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
