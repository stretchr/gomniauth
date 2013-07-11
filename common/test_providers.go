package common

import (
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/mock"
)

type TestProvider struct {
	mock.Mock
}

func (p *TestProvider) Name() string {
	return p.Called().String(0)
}

func (p *TestProvider) Config() objects.Map {
	return p.Called().Get(0).(objects.Map)
}

func (p *TestProvider) AuthType() AuthType {
	return p.Called().Get(0).(AuthType)
}

type TestProvider2 struct {
	mock.Mock
}

func (p *TestProvider2) Name() string {
	return p.Called().String(0)
}

func (p *TestProvider2) Config() objects.Map {
	return p.Called().Get(0).(objects.Map)
}

func (p *TestProvider2) AuthType() AuthType {
	return p.Called().Get(0).(AuthType)
}
