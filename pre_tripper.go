package gomniauth

type PreTripper interface {
	PreRoundTrip(creds *common.Credentials) error
}
