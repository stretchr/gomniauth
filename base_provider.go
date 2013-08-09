package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"net/http"
)

type BaseProvider struct {
	Config         *common.Config
	TripperFactory common.TripperFactory
}

func (p *BaseProvider) NewTripper(creds *common.Credentials) (common.Tripper, error) {
	return p.TripperFactory.NewTripper(creds, p)
}

func (p *BaseProvider) SetTripperFactory(factory common.TripperFactory) {
	p.TripperFactory = factory
}

// GetClient gets an http.Client authenticated with the specified
// common.Credentials.
func (p *BaseProvider) GetClient(creds *common.Credentials) (*http.Client, error) {

	tripper, tripperErr := p.NewTripper(creds)

	if tripperErr != nil {
		return nil, tripperErr
	}

	return &http.Client{Transport: tripper}, nil
}

// The functions below are here only to satisfy the interface and will be overridden
// when this type is composed into
func (p *BaseProvider) Name() string {
	return ""
}
func (p *BaseProvider) CompleteAuth(data objects.Map) (*common.Credentials, error) {
	return nil, nil
}
func (p *BaseProvider) GetBeginAuthURL(state *common.State) (string, error) {
	return "", nil
}
func (p *BaseProvider) LoadUser(creds *common.Credentials) (common.User, error) {
	return nil, nil
}
