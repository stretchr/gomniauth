package gomniauth

var SecurityKey string = ""

func GetSecurityKey() string {
	if len(SecurityKey) == 0 {
		panic("gomniauth: You must set gomniauth.SecurityKey to something secure.")
	}
	return SecurityKey
}
