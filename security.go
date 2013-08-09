package gomniauth

import (
	"github.com/stretchr/gomniauth/common"
)

func SetSecurityKey(key string) {
	common.SetSecurityKey(key)
}

func GetSecurityKey() string {
	return common.GetSecurityKey()
}
