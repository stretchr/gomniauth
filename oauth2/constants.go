package oauth2

const (

	// ApprovalPromptForce indicates that the user will always
	// have to reauthorize access if the AccessType is online.
	ApprovalPromptForce string = "force"

	// ApprovalPromptAuto indicates that the user will not
	// have to reauthorize access.
	ApprovalPromptAuto string = "auto"
)

const (

	// AccessTypeOnline indicates that the access type is online.
	AccessTypeOnline string = "online"

	// AccessTypeOffline indicates that the access type is offline.
	AccessTypeOffline string = "offline"
)
