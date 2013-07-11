package common

import (
	"github.com/stretchr/stew/objects"
)

type Provider interface {

	// Name gets the name of this provider.  Name must be URL friendly.
	Name() string

	// Config gets the default configuration for this Provider.
	Config() objects.Map

	// common.AuthType gets the AuthType that this provider uses.
	AuthType() AuthType
}
