package oauth2

import (
	"github.com/stretchr/codecs/services"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
	"io/ioutil"
)

// Get executes an authenticated HTTP GET against the given provider and returns an
// objx.Map of the response.
//
// The response type is automatically detected and used to unmarshal the response.
func Get(provider common.Provider, creds *common.Credentials, endpoint string) (objx.Map, error) {

	client, clientErr := provider.GetClient(creds)

	if clientErr != nil {
		return nil, clientErr
	}

	response, responseErr := client.Get(endpoint)

	if responseErr != nil {
		return nil, responseErr
	}

	body, bodyErr := ioutil.ReadAll(response.Body)

	if bodyErr != nil {
		return nil, bodyErr
	}

	defer response.Body.Close()

	codecs := services.NewWebCodecService()
	codec, getCodecErr := codecs.GetCodec(response.Header.Get("Content-Type"))

	if getCodecErr != nil {
		return nil, getCodecErr
	}

	var data objx.Map
	unmarshalErr := codec.Unmarshal(body, &data)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return data, nil

}
