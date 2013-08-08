package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
	"net/http"
)

type Tripper interface {
	http.RoundTripper
	// Credentials gets the authentication credentials that
	// this Tripper will use.
	Credentials() *common.Credentials
}
