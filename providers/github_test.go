package providers

import (
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBeginAuthURL(t *testing.T) {

	g := new(Github)
	url, err := g.BeginAuthURL(objects.M("state", objects.M("proj", "test", "after", "success")))

	if assert.NoError(t, err) {
		assert.Equal(t, "", url)
	}

}
