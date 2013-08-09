package common

import (
	"github.com/stretchr/stew/objects"
)

// State represents a map of state arguments that can be used to
// persist values across the authentication process.
type State struct {
	objects.Map
}

// NewState creates a new object that can be used to persist
// state across authentication requests.
func NewState(keyAndValuePairs ...interface{}) *State {
	return &State{objects.M(keyAndValuePairs...)}
}
