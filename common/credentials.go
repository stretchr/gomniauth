package common

import (
	"github.com/stretchr/stew/objects"
)

// Credentials represent data that describes information needed
// to make authenticated requests.
type Credentials struct {
	objects.Map
}

var EmptyCredentials *Credentials = nil
