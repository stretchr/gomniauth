package oauth2

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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
