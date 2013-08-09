package gomniauth

import (
	"encoding/json"

	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/stew/objects"
	"io/ioutil"
	"mime"
	"time"
)

const (
	OAuth2KeyClientID       string = "client_id"
	OAuth2KeySecret         string = "client_secret"
	OAuth2KeyRedirectUrl    string = "redirect_url"
	OAuth2KeyScope          string = "scope"
	OAuth2KeyAccessType     string = "access_type"
	OAuth2KeyApprovalPrompt string = "approval_prompt"
	OAuth2KeyAuthURL        string = "auth_url"
	OAuth2KeyTokenURL       string = "token_url"
	OAuth2KeyCode           string = "code"
	OAuth2KeyGrantType      string = "grant_type"
	OAuth2KeyExpiresIn      string = "expires_in"
	OAuth2KeyAccessToken    string = "access_token"
	OAuth2KeyRefreshToken   string = "refresh_token"
)

const (
	OAuth2GrantTypeAuthorizationCode = "authorization_code"
)

const (

	// ApprovalPromptForce indicates that the user will always
	// have to reauthorize access if the AccessType is online.
	OAuth2ApprovalPromptForce string = "force"

	// ApprovalPromptAuto indicates that the user will not
	// have to reauthorize access.
	OAuth2ApprovalPromptAuto string = "auto"
)

const (

	// AccessTypeOnline indicates that the access type is online.
	OAuth2AccessTypeOnline string = "online"

	// AccessTypeOffline indicates that the access type is offline.
	OAuth2AccessTypeOffline string = "offline"
)

type OAuth2Provider struct {
	Provider
}

func (p *OAuth2Provider) GetBeginAuthURLWithBase(base string, state *common.State, config *common.Config) (string, error) {

	if config == nil {
		panic("OAuth2Handler: Must have valid Config specified.")
	}

	// copy the config
	params := objects.M(
		OAuth2KeyClientID, config.GetStringOrEmpty(OAuth2KeyClientID),
		OAuth2KeyRedirectUrl, config.GetStringOrEmpty(OAuth2KeyRedirectUrl),
		OAuth2KeyScope, config.GetStringOrEmpty(OAuth2KeyScope),
		OAuth2KeyAccessType, config.GetStringOrEmpty(OAuth2KeyAccessType),
		OAuth2KeyApprovalPrompt, config.GetStringOrEmpty(OAuth2KeyApprovalPrompt))

	// set the state
	stateValue, stateErr := state.SignedBase64(GetSecurityKey())

	if stateErr != nil {
		return "", stateErr
	}

	params.Set("state", stateValue)

	// generate the query part
	query, queryErr := params.URLQuery()

	if queryErr != nil {
		return "", queryErr
	}

	// put the strings together
	return base + "?" + query, nil
}

// CompleteAuth takes a map of arguments that are used to
// complete the authorisation process, completes it, and returns
// the appropriate common.Credentials.
//
// The data must contain an OAuth2KeyCode obtained from the auth
// server.
func (g *OAuth2Provider) CompleteAuth(data objects.Map) (*common.Credentials, error) {

	// get the code
	code := data.GetStringOrEmpty("code")
	if len(code) == 0 {
		return nil, &common.MissingParameterError{"code"}
	}

	client, clientErr := g.GetClient(common.EmptyCredentials)
	if clientErr != nil {
		return nil, clientErr
	}

	params := objects.M(OAuth2KeyGrantType, OAuth2GrantTypeAuthorizationCode,
		OAuth2KeyRedirectUrl, g.Config.GetStringOrEmpty(OAuth2KeyRedirectUrl),
		OAuth2KeyScope, g.Config.GetStringOrEmpty(OAuth2KeyScope),
		OAuth2KeyCode, code,
		OAuth2KeyClientID, g.Config.GetStringOrEmpty(OAuth2KeyClientID),
		OAuth2KeySecret, g.Config.GetStringOrEmpty(OAuth2KeySecret))

	// post the form
	response, requestErr := client.PostForm(g.Config.GetString(OAuth2KeyAuthURL), params.URLValues())

	if requestErr != nil {
		return nil, requestErr
	}

	// make sure we close the body
	defer func() {
		if response.Body != nil {
			response.Body.Close()
		}
	}()

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

		expiresIn, expiresErr := time.ParseDuration(vals.GetStringOrEmpty(OAuth2KeyExpiresIn) + "s")

		if expiresErr != nil {
			return nil, expiresErr
		}

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
