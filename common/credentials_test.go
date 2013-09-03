package common

import (
	"github.com/stretchr/codecs"
	"github.com/stretchr/core/stretchr/constants"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCredentials_PublicData(t *testing.T) {

	creds := &Credentials{objects.M(constants.ParamAuthCode, "ABC123", CredentialsKeyID, 123)}

	publicData, _ := codecs.PublicData(creds, nil)

	if assert.NotNil(t, publicData) {
		assert.Equal(t, "ABC123", publicData.(objects.Map)[constants.ParamAuthCode])
		assert.Equal(t, "123", publicData.(objects.Map)[CredentialsKeyID], "CredentialsKeyID ("+CredentialsKeyID+") must be turned into a string.")
	}

}
