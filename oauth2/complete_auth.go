package oauth2

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
	"io/ioutil"
	"mime"
	"net/http"
	"time"
)

// CompleteAuth takes a map of arguments that are used to
// complete the authorisation process, completes it, and returns
// the appropriate common.Credentials.
//
// The data must contain an OAuth2KeyCode obtained from the auth
// server.
func CompleteAuth(tripperFactory common.TripperFactory, data objx.Map, config *common.Config, provider common.Provider) (*common.Credentials, error) {

	// get the code
	codeList := data.Get(OAuth2KeyCode).Data()

	code, ok := codeList.(string)
	if !ok {

		if codeList == nil || len(codeList.([]string)) == 0 {
			return nil, &common.MissingParameterError{ParameterName: OAuth2KeyCode}
		}
		code = codeList.([]string)[0]
		if len(code) == 0 {
			return nil, &common.MissingParameterError{ParameterName: OAuth2KeyCode}
		}
	}

	client, clientErr := GetClient(tripperFactory, common.EmptyCredentials, provider)
	if clientErr != nil {
		return nil, clientErr
	}

	params := objx.MSI(OAuth2KeyGrantType, OAuth2GrantTypeAuthorizationCode,
		OAuth2KeyRedirectUrl, config.Get(OAuth2KeyRedirectUrl).Str(),
		OAuth2KeyScope, config.Get(OAuth2KeyScope).Str(),
		OAuth2KeyCode, code,
		OAuth2KeyClientID, config.Get(OAuth2KeyClientID).Str(),
		OAuth2KeySecret, config.Get(OAuth2KeySecret).Str())

	// post the form
	response, requestErr := client.PostForm(config.Get(OAuth2KeyTokenURL).Str(), params.URLValues())

	if requestErr != nil {
		return nil, requestErr
	}

	// make sure we close the body
	defer func() {
		if response.Body != nil {
			response.Body.Close()
		}
	}()

	// make sure we have an OK response
	if response.StatusCode != http.StatusOK {
		return nil, &common.AuthServerError{
			ErrorMessage: fmt.Sprintf("Server replied with %s.", response.Status),
			Response: response,
		}
	}

	content, _, mimeTypeErr := mime.ParseMediaType(response.Header.Get("Content-Type"))

	if mimeTypeErr != nil {
		return nil, mimeTypeErr
	}

	// prepare the credentials object
	creds := &common.Credentials{Map: objx.MSI()}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch content {
	case "application/x-www-form-urlencoded", "text/plain":

		vals, err := objx.FromURLQuery(string(body))
		if err != nil {
			return nil, err
		}

		// did an error occur?
		if len(vals.Get("error").Str()) > 0 {
			return nil, &common.AuthServerError{
				ErrorMessage: vals.Get("error").Str(),
				Response: response,
			}
		}

		expiresIn, _ := time.ParseDuration(vals.Get(OAuth2KeyExpiresIn).Str() + "s")

		creds.Set(OAuth2KeyAccessToken, vals.Get(OAuth2KeyAccessToken).Str())
		creds.Set(OAuth2KeyRefreshToken, vals.Get(OAuth2KeyRefreshToken).Str())
		creds.Set(OAuth2KeyExpiresIn, expiresIn)

	default: // use JSON

		var data objx.Map

		jsonErr := json.Unmarshal(body, &data)

		if jsonErr != nil {
			return nil, jsonErr
		}

		// handle the time
		timeDuration := data.Get(OAuth2KeyExpiresIn).Float64()
		data.Set(OAuth2KeyExpiresIn, time.Duration(timeDuration)*time.Second)

		// merge this data into the creds
		creds.MergeHere(data)

	}

	return creds, nil
}
