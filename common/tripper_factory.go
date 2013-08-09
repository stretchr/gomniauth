package common

// TripperFactory describes an object responsible for making
// authenticated Trippers.
type TripperFactory interface {
	NewTripper(*Credentials, Provider) (Tripper, error)
}
