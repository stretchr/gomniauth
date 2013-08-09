package common

const (
	errorMessagePrefix string = "gomniauth: "
)

/*
  MissingParameterError
  -------------------------------------------------------
*/

type MissingParameterError struct {
	ParameterName string
}

func (e *MissingParameterError) Error() string {
	return errorMessagePrefix + "Parameter '" + e.ParameterName + "' is required but missing."
}

/*
  AuthServerError
  -------------------------------------------------------
*/

type AuthServerError struct {
	ErrorMessage string
}

func (e *AuthServerError) Error() string {
	return errorMessagePrefix + "Auth server responded with error '" + e.ErrorMessage + "'."
}

/*
  AuthServerError
  -------------------------------------------------------
*/

type MissingProviderError struct {
	ProviderName string
}

func (e *MissingProviderError) Error() string {
	return errorMessagePrefix + "No provider with name '" + e.ProviderName + "' was found."
}
