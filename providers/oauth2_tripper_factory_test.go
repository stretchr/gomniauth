package providers

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOAuth2TripperFactoryNewTripper(t *testing.T) {

	var tripperFactory gomniauth.TripperFactory
	tripperFactory = new(OAuth2TripperFactory)

	assert.NotNil(t, tripperFactory)

	creds := new(common.Credentials)
	var tripper gomniauth.Tripper
	tripper, err := tripperFactory.NewTripper(creds)

	if assert.NotNil(t, tripper) && assert.NoError(t, err) {

		assert.Equal(t, creds, tripper.Credentials())
		assert.IsType(t, new(OAuth2Tripper), tripper, "OAuth2TripperFactory should make OAuth2Trippers")

	}

}
