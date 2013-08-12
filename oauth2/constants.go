package oauth2

const (
	OAuth2KeyClientID       string = "client_id"
	OAuth2KeySecret         string = "client_secret"
	OAuth2KeyRedirectUrl    string = "redirect_uri"
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
	OAuth2KeyResponseType   string = "response_type"
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
