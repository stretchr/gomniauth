package common

type PreTripper interface {
	PreRoundTrip(tripper Tripper) error
}
