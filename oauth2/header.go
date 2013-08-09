package oauth2

import (
	"github.com/stretchr/gomniauth/common"
)

func AuthorizationHeader(creds *common.Credentials) (key, value string) {
	return "Authorization", "Bearer " + creds.GetStringOrDefault(OAuth2KeyAccessToken, "Invalid")
}
