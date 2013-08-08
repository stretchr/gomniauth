package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultOAuth2Handler(t *testing.T) {

	assert.NotNil(t, DefaultOAuth2Handler)
	assert.IsType(t, &OAuth2Handler{}, DefaultOAuth2Handler)

}
