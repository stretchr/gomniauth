package gomniauth

import (
	"fmt"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

// ProviderPublicData gets the public data for the specified provider.
//
// The options should contain the `loginpathFormat`, which will determine how the
// loginpath value is created.
func ProviderPublicData(provider common.Provider, options map[string]interface{}) (interface{}, error) {

	optionsx := objx.New(options)

	return map[string]interface{}{
		"name":      provider.Name(),
		"display":   provider.DisplayName(),
		"loginpath": fmt.Sprintf(optionsx.Get("loginpathFormat").Str("auth/%s/login"), provider.Name()),
	}, nil

}
