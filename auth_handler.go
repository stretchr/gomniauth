package gomniauth

import (
	"net/http"
)

type AuthHandler interface {
	NewRoundTripper() (http.RoundTripper, error)
}
