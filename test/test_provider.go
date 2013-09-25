package test

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type TestProvider struct {
	mock.Mock
}

// Name is the unique name for this provider.
func (p *TestProvider) Name() string {
	return p.Called().String(0)
}

// DisplayName is the human readable name for this provider.
func (p *TestProvider) DisplayName() string {
	return p.Called().String(0)
}

// GetBeginAuthURL gets the URL that the client must visit in order
// to begin the authentication process.
func (p *TestProvider) GetBeginAuthURL(state *common.State, options objx.Map) (string, error) {
	args := p.Called(state, options)
	return args.String(0), args.Error(1)
}

// CompleteAuth takes a map of arguments that are used to
// complete the authorisation process, completes it, and returns
// the appropriate common.Credentials.
func (p *TestProvider) CompleteAuth(data objx.Map) (*common.Credentials, error) {
	args := p.Called(data)
	return args.Get(0).(*common.Credentials), args.Error(1)
}

func (p *TestProvider) Get(creds *common.Credentials, endpoint string) (objx.Map, error) {
	args := p.Called(creds, endpoint)
	return args.Get(0).(objx.Map), args.Error(1)
}

// GetUser uses the specified common.Credentials to access the users profile
// from the remote provider, and builds the appropriate User object.
func (p *TestProvider) GetUser(creds *common.Credentials) (common.User, error) {
	args := p.Called(creds)
	return args.Get(0).(common.User), args.Error(1)
}

// GetClient gets an http.Client authenticated with the specified
// common.Credentials.
func (p *TestProvider) GetClient(creds *common.Credentials) (*http.Client, error) {
	args := p.Called(creds)
	return args.Get(0).(*http.Client), args.Error(1)
}

func (p *TestProvider) PublicData(options map[string]interface{}) (publicData interface{}, err error) {
	args := p.Called(options)
	return args.Get(0), args.Error(1)
}
