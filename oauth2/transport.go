package oauth2

import (
	"encoding/json"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"time"
)

// Transport implements http.RoundTripper. When configured with a valid
// Config and Token it can be used to make authenticated HTTP requests.
//
//      t := &oauth.Transport{config}
//      t.Exchange(code)
//      // t now contains a valid Token
//      r, _, err := t.Client().Get("http://example.org/url/requiring/auth")
//
// It will automatically refresh the Token if it can,
// updating the supplied Token in place.
type Transport struct {
	Config *Config
	Token  *Token

	// Transport is the HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	// (It should never be an oauth.Transport.)
	Transport http.RoundTripper
}

// transport gets the http.RoundTripper or uses the http.DefaultTransport
// if none is set.
func (t *Transport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// Client returns an *http.Client that makes OAuth-authenticated requests.
func (t *Transport) Client() *http.Client {
	return &http.Client{Transport: t}
}

// RoundTrip executes a single HTTP transaction using the Transport's
// Token as authorization headers.
//
// This method will attempt to renew the Token if it has expired and may return
// an error related to that Token renewal before attempting the client request.
// If the Token cannot be renewed a non-nil os.Error value will be returned.
// If the Token is invalid callers should expect HTTP-level errors,
// as indicated by the Response's StatusCode.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {

	if t.Token == nil {
		if t.Config == nil {
			return nil, OAuth2Error{"RoundTrip", "no Config supplied"}
		}
		return nil, OAuth2Error{"RoundTrip", "no Token supplied"}
	}

	// Refresh the Token if it has expired.
	if t.Token.HasExpired() {
		if err := t.Refresh(); err != nil {
			return nil, err
		}
	}

	// To set the Authorization header, we must make a copy of the Request
	// so that we don't modify the Request we were given.
	// This is required by the specification of http.RoundTripper.
	req = cloneRequest(req)
	req.Header.Set("Authorization", "Bearer "+t.Token.AccessToken)

	// Make the HTTP request.
	return t.transport().RoundTrip(req)
}

// Refresh renews the Transport's AccessToken using its RefreshToken.
func (t *Transport) Refresh() error {

	if t.Token == nil {
		return OAuth2Error{"Refresh", "no existing Token"}
	}
	if t.Token.RefreshToken == "" {
		return OAuth2Error{"Refresh", "Token expired; no Refresh Token"}
	}
	if t.Config == nil {
		return OAuth2Error{"Refresh", "no Config supplied"}
	}

	err := t.updateToken(t.Token, url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {t.Token.RefreshToken},
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *Transport) updateToken(tok *Token, v url.Values) error {
	v.Set("client_id", t.Config.ClientId())
	v.Set("client_secret", t.Config.ClientSecret())
	r, err := (&http.Client{Transport: t.transport()}).PostForm(t.Config.TokenURL(), v)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return OAuth2Error{"updateToken", r.Status}
	}
	var b struct {
		Access    string        `json:"access_token"`
		Refresh   string        `json:"refresh_token"`
		ExpiresIn time.Duration `json:"expires_in"`
	}

	content, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	switch content {
	case "application/x-www-form-urlencoded", "text/plain":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		vals, err := url.ParseQuery(string(body))
		if err != nil {
			return err
		}

		b.Access = vals.Get("access_token")
		b.Refresh = vals.Get("refresh_token")
		b.ExpiresIn, _ = time.ParseDuration(vals.Get("expires_in") + "s")
	default:
		if err = json.NewDecoder(r.Body).Decode(&b); err != nil {
			return err
		}
		// The JSON parser treats the unitless ExpiresIn like 'ns' instead of 's' as above,
		// so compensate here.
		b.ExpiresIn *= time.Second
	}
	tok.AccessToken = b.Access
	// Don't overwrite `RefreshToken` with an empty value
	if len(b.Refresh) > 0 {
		tok.RefreshToken = b.Refresh
	}
	if b.ExpiresIn == 0 {
		tok.Expiry = time.Time{}
	} else {
		tok.Expiry = time.Now().Add(b.ExpiresIn)
	}
	return nil
}

// Exchange takes a code and gets access Token from the remote server.
func (t *Transport) Exchange(code string) (*Token, error) {

	if t.Config == nil {
		return nil, OAuth2Error{"Exchange", "no Config supplied"}
	}

	// If the transport already has a token, it is
	// passed to `updateToken` to preserve existing refresh token.
	tok := t.Token

	if tok == nil {
		tok = new(Token)
	}
	err := t.updateToken(tok, url.Values{
		"grant_type":   {"authorization_code"},
		"redirect_uri": {t.Config.RedirectURL()},
		"scope":        {t.Config.Scope()},
		"code":         {code},
	})
	if err != nil {
		return nil, err
	}
	t.Token = tok

	return tok, nil
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header)
	for k, s := range r.Header {
		r2.Header[k] = s
	}
	return r2
}
