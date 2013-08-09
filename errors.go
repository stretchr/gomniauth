package gomniauth

type MissingParameterError struct {
	ParameterName string
}

func (e *MissingParameterError) Error() string {
	return "Parameter '" + e.ParameterName + "' is required but missing."
}

type AuthServerError struct {
	ErrorMessage string
}

func (e *AuthServerError) Error() string {
	return "Auth server responded with error '" + e.ErrorMessage + "'."
}
