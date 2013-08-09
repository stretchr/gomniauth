package common

import (
	"github.com/stretchr/stew/objects"
)

const (
	CredentialsKeyID string = "id"
)

// Credentials represent data that describes information needed
// to make authenticated requests.
type Credentials struct {
	objects.Map
}

var EmptyCredentials *Credentials = nil
