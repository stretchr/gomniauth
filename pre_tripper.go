package gomniauth

type PreTripper interface {
	PreRoundTrip(tripper Tripper) error
}
