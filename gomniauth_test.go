package gomniauth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	baseURL = "http://stretchr.com/~auth"
)

func StretchrGomniauth() *Gomniauth {
	return MakeGomniauth(baseURL)
}

func TestGomniauthMakeGomniauth(t *testing.T) {

	g := StretchrGomniauth()

	assert.Equal(t, g.baseURL, baseURL)

}

func TestGomniauthAddProvider(t *testing.T) {

	g := StretchrGomniauth()

	g.AddProvider(Github, "123", "456")

	if assert.Equal(t, len(g.providers), 1) {
		assert.Equal(t, g.providers[Github].Config.ClientId, "123")
		assert.Equal(t, g.providers[Github].Config.ClientSecret, "456")
		assert.Equal(t, g.providers[Github].Config.RedirectURL, "http://stretchr.com/~auth/github/callback")
	}

}
