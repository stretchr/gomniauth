package gomniauth

type AuthService struct {
	providers   map[string]Provider
	dataService DataService
}

func (s *AuthService) DataService() DataService {
	return s.dataService
}

func (s *AuthService) Provider(name string) Provider {
	return nil
}
