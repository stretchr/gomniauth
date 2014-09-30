package common

import (
	"net/http"
)

const (
	PrefixForErrors string = "gomniauth: "
)

/*
  MissingParameterError
  -------------------------------------------------------
*/

type MissingParameterError struct {
	ParameterName string
}

func (e *MissingParameterError) Error() string {
	return PrefixForErrors + "Parameter '" + e.ParameterName + "' is required but missing."
}

/*
  AuthServerError
  -------------------------------------------------------
*/

type AuthServerError struct {
	ErrorMessage string
	Response     *http.Response
}

func (e *AuthServerError) Error() string {
	return PrefixForErrors + "Auth server responded with error '" + e.ErrorMessage + "'."
}

/*
  AuthServerError
  -------------------------------------------------------
*/

type MissingProviderError struct {
	ProviderName string
}

func (e *MissingProviderError) Error() string {
	return PrefixForErrors + "No provider with name '" + e.ProviderName + "' was found."
}
