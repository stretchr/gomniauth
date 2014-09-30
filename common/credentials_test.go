package common

import (
	"testing"

	"github.com/stretchr/codecs"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
)

func TestCredentials_PublicData(t *testing.T) {

	creds := &Credentials{objx.MSI("auth", "ABC123", CredentialsKeyID, 123)}

	publicData, _ := codecs.PublicData(creds, nil)

	if assert.NotNil(t, publicData) {
		assert.Equal(t, "ABC123", publicData.(objx.Map)["auth"])
		assert.Equal(t, "123", publicData.(objx.Map)[CredentialsKeyID], "CredentialsKeyID ("+CredentialsKeyID+") must be turned into a string.")
	}

}
