package common

import (
	"github.com/stretchr/signature"
)

var securityKey string = ""

func SetSecurityKey(key string) {
	securityKey = key
}

func GetSecurityKey() string {
	if len(securityKey) == 0 {
		panic(PrefixForErrors + "You must set gomniauth.SecurityKey to a secret key.  Why not use this one that we made just for you:\n\n\tgomniauth.SetSecurityKey(\"" + signature.RandomKey(64) + "\")\n")
	}
	return securityKey
}
