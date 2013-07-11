package common

import (
	"errors"
)

type AuthType uint8

const (
	AuthTypeInvalid AuthType = iota

	AuthTypeOAuth2
)

var authTypeStrings = map[string]AuthType{
	"":       AuthTypeInvalid,
	"oauth2": AuthTypeOAuth2,
}

func (a AuthType) String() string {
	for authTypeStr, authType := range authTypeStrings {
		if authType == a {
			return authTypeStr
		}
	}
	return ""
}

func ParseAuthType(s string) (AuthType, error) {
	if authType, ok := authTypeStrings[s]; ok {
		return authType, nil
	}
	return AuthTypeInvalid, errors.New("gomniauth: \"" + s + "\" is not a valid AuthType.")
}
