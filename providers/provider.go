package providers

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"net/http"
)

type Provider struct {
	config         *common.Config
	tripperFactory gomniauth.TripperFactory
}

func (p *Provider) NewTripper(creds *common.Credentials) (gomniauth.Tripper, error) {
	return p.tripperFactory.NewTripper(creds)
}

func (p *Provider) SetTripperFactory(factory gomniauth.TripperFactory) {
	p.tripperFactory = factory
}

// GetClient gets an http.Client authenticated with the specified
// common.Credentials.
func (p *Provider) GetClient(creds *common.Credentials) (*http.Client, error) {

	tripper, tripperErr := p.NewTripper(creds)

	if tripperErr != nil {
		return nil, tripperErr
	}

	return &http.Client{Transport: tripper}, nil
}
