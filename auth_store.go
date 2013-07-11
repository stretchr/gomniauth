package gomniauth

type AuthStore interface {
	GetAuth(id string) (*Auth, error)

	PutAuth(id string, auth *Auth) error
}
