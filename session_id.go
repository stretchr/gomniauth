package gomniauth

import (
	"github.com/nu7hatch/gouuid"
)

// CreateSessionID creates a unique session ID string.
func CreateSessionID() string {
	u, _ := uuid.NewV4()
	return u.String()
}
