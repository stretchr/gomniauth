package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOAuth2TripperFactoryNewTripper(t *testing.T) {

	g := new(OAuth2Provider)

	creds := new(common.Credentials)
	var tripperFactory common.TripperFactory
	tripperFactory = new(OAuth2TripperFactory)

	assert.NotNil(t, tripperFactory)

	var tripper common.Tripper
	tripper, err := tripperFactory.NewTripper(creds, g)

	if assert.NotNil(t, tripper) && assert.NoError(t, err) {

		assert.Equal(t, creds, tripper.Credentials())
		assert.IsType(t, new(OAuth2Tripper), tripper, "OAuth2TripperFactory should make OAuth2Trippers")

	}

}
