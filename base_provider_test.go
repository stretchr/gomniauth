package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBaseProviderNewTripper(t *testing.T) {

	creds := new(common.Credentials)
	provider := new(BaseProvider)

	testTripperFactory := new(test.TestTripperFactory)
	provider.TripperFactory = testTripperFactory

	testTripper := new(test.TestTripper)
	testTripperFactory.On("NewTripper", creds, provider).Return(testTripper, nil)

	returnedTripper, err := provider.NewTripper(creds)

	if assert.NoError(t, err) {
		assert.Equal(t, returnedTripper, testTripper)
	}

	mock.AssertExpectationsForObjects(t, testTripperFactory.Mock)

}

func TestBaseProviderSetTripperFactory(t *testing.T) {

	testTripperFactory := new(test.TestTripperFactory)
	provider := new(BaseProvider)

	provider.SetTripperFactory(testTripperFactory)

	assert.Equal(t, testTripperFactory, provider.TripperFactory)

}

func TestBaseProviderGetClient(t *testing.T) {

	//g := Github("clientID", "secret", "http://myapp.com/")
	g := new(BaseProvider)

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)

	g.SetTripperFactory(testTripperFactory)
	creds := new(common.Credentials)

	testTripperFactory.On("NewTripper", creds, mock.Anything).Return(testTripper, nil)

	client, clientErr := g.GetClient(creds)

	if assert.NoError(t, clientErr) {
		if assert.NotNil(t, client) {
			assert.Equal(t, client.Transport, testTripper)
		}
	}

	mock.AssertExpectationsForObjects(t, testTripperFactory.Mock, testTripper.Mock)

}
