package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
)

// AuthStore represents an object capable of storing or caching
// common.Auth objects.
type AuthStore interface {
	// GetAuth gets the common.Auth object for the specified ID,
	// or returns nil if none could be found.
	GetAuth(id string) (*common.Auth, error)

	// PutAuth adds or updates the common.Auth object for the specified ID.
	PutAuth(id string, auth *common.Auth) error

	// DeleteAuth deletes the common.Auth object for the specified ID.
	DeleteAuth(id string) error
}
