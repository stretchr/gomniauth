package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
)

// TripperFactory describes an object responsible for making
// authenticated Trippers.
type TripperFactory interface {
	NewTripper(*common.Credentials) (Tripper, error)
}
