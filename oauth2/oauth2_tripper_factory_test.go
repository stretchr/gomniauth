package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOAuth2TripperFactoryNewTripper(t *testing.T) {

	testProvider := new(test.TestProvider)

	creds := new(common.Credentials)
	var tripperFactory common.TripperFactory
	tripperFactory = new(OAuth2TripperFactory)

	assert.NotNil(t, tripperFactory)

	var tripper common.Tripper
	tripper, err := tripperFactory.NewTripper(creds, testProvider)

	if assert.NotNil(t, tripper) && assert.NoError(t, err) {

		assert.Equal(t, creds, tripper.Credentials())
		assert.IsType(t, new(OAuth2Tripper), tripper, "OAuth2TripperFactory should make OAuth2Trippers")

	}

	mock.AssertExpectationsForObjects(t, testProvider.Mock)

}
