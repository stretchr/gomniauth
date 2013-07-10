package gomniauth

type TokenStore interface {
	GetToken(id string) (*Token, error)

	PutToken(id string, token *Token) error
}
