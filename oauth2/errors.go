package oauth2

type OAuth2Error struct {
	prefix string
	msg    string
}

func (oe OAuth2Error) Error() string {
	return "oauth2: " + oe.prefix + ": " + oe.msg
}
