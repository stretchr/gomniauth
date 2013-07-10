package common

type AuthType uint8

const (
	AuthTypeUnknown AuthType = iota
	AuthTypeOAuth2
)
