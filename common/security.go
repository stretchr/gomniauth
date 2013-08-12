package common

import (
	"github.com/stretchr/signature"
)

var securityKey string = ""

// SetSecurityKey sets the global security key to be used for signing the state variable
// in the auth request. This allows gomniauth to detect if the data in the
// state variable has been changed.
func SetSecurityKey(key string) {
	securityKey = key
}

// GetSecurityKey gets the global security key.
func GetSecurityKey() string {
	if len(securityKey) == 0 {
		panic(PrefixForErrors + "You must set gomniauth.SecurityKey to a secret key.  Why not use this one that we made just for you:\n\n\tgomniauth.SetSecurityKey(\"" + signature.RandomKey(64) + "\")\n")
	}
	return securityKey
}
