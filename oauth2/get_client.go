package oauth2

import (
	"github.com/stretchr/gomniauth/common"
	"net/http"
)

// GetClient gets an http.Client authenticated with the specified
// common.Credentials.
func GetClient(tripperFactory common.TripperFactory, creds *common.Credentials, provider common.Provider) (*http.Client, error) {

	tripper, tripperErr := tripperFactory.NewTripper(creds, provider)

	if tripperErr != nil {
		return nil, tripperErr
	}

	return &http.Client{Transport: tripper}, nil
}
