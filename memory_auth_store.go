package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
)

// MemoryAuthStore is a simple in-memory implementation of AuthStore.
//
// MemoryAuthStore is great for development, or single machine applications,
// but for larger scale situations, you should consider using an AuthStore
// backed by some kind of persistant storage.
type MemoryAuthStore map[string]*common.Auth

// GetAuth gets the common.Auth object for the specified ID,
// or returns nil if none could be found.
func (s MemoryAuthStore) GetAuth(id string) (*common.Auth, error) {
	return s[id], nil
}

// PutAuth adds or updates the common.Auth object for the specified ID.
func (s MemoryAuthStore) PutAuth(id string, auth *common.Auth) error {
	s[id] = auth
	return nil
}

// DeleteAuth deletes the common.Auth object for the specified ID.
func (s MemoryAuthStore) DeleteAuth(id string) error {
	delete(s, id)
	return nil
}
