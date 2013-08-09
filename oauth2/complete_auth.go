package oauth2

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
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
func CompleteAuth(tripperFactory common.TripperFactory, data objects.Map, config *common.Config, provider common.Provider) (*common.Credentials, error) {

	// get the code
	codeList := data.Get(OAuth2KeyCode)
	if codeList == nil || len(codeList.([]string)) == 0 {
		return nil, &common.MissingParameterError{OAuth2KeyCode}
	}
	code := codeList.([]string)[0]
	if len(code) == 0 {
		return nil, &common.MissingParameterError{OAuth2KeyCode}
	}

	client, clientErr := GetClient(tripperFactory, common.EmptyCredentials, provider)
	if clientErr != nil {
		return nil, clientErr
	}

	params := objects.M(OAuth2KeyGrantType, OAuth2GrantTypeAuthorizationCode,
		OAuth2KeyRedirectUrl, config.GetStringOrEmpty(OAuth2KeyRedirectUrl),
		OAuth2KeyScope, config.GetStringOrEmpty(OAuth2KeyScope),
		OAuth2KeyCode, code,
		OAuth2KeyClientID, config.GetStringOrEmpty(OAuth2KeyClientID),
		OAuth2KeySecret, config.GetStringOrEmpty(OAuth2KeySecret))

	// post the form
	response, requestErr := client.PostForm(config.GetString(OAuth2KeyTokenURL), params.URLValues())

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
		return nil, &common.AuthServerError{fmt.Sprintf("Server replied with %s.", response.Status)}
	}

	content, _, mimeTypeErr := mime.ParseMediaType(response.Header.Get("Content-Type"))

	if mimeTypeErr != nil {
		return nil, mimeTypeErr
	}

	// prepare the credentials object
	creds := &common.Credentials{objects.M()}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch content {
	case "application/x-www-form-urlencoded", "text/plain":

		vals, err := objects.NewMapFromURLQuery(string(body))
		if err != nil {
			return nil, err
		}

		// did an error occur?
		if len(vals.GetStringOrEmpty("error")) > 0 {
			return nil, &common.AuthServerError{vals.GetStringOrEmpty("error")}
		}

		expiresIn, _ := time.ParseDuration(vals.GetStringOrEmpty(OAuth2KeyExpiresIn) + "s")

		creds.Set(OAuth2KeyAccessToken, vals.GetStringOrEmpty(OAuth2KeyAccessToken))
		creds.Set(OAuth2KeyRefreshToken, vals.GetStringOrEmpty(OAuth2KeyRefreshToken))
		creds.Set(OAuth2KeyExpiresIn, expiresIn)

	default: // use JSON

		var data objects.Map

		jsonErr := json.Unmarshal(body, &data)

		if jsonErr != nil {
			return nil, jsonErr
		}

		// handle the time
		timeDuration := data.Get(OAuth2KeyExpiresIn).(float64)
		data.Set(OAuth2KeyExpiresIn, time.Duration(timeDuration)*time.Second)

		// merge this data into the creds
		creds.MergeHere(data)

	}

	return creds, nil
}
