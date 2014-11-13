package main

import (
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"io"
	"net/http"
)

const (
	// NOTE: Don't change this, the auth settings on the providers
	// are coded to this path for this example.
	Address string = ":8080"
)

func main() {

	// setup the providers
	gomniauth.SetSecurityKey("yLiCQYG7CAflDavqGH461IO0MHp7TEbpg6TwHBWdJzNwYod1i5ZTbrIF5bEoO3oP") // NOTE: DO NOT COPY THIS - MAKE YOR OWN!
	gomniauth.WithProviders(
		github.New("3d1e6ba69036e0624b61", "7e8938928d802e7582908a5eadaaaf22d64babf1", "http://localhost:8080/auth/github/callback"),
		google.New("1051709296778.apps.googleusercontent.com", "7oZxBGwpCI3UgFMgCq80Kx94", "http://localhost:8080/auth/google/callback"),
		facebook.New("537611606322077", "f9f4d77b3d3f4f5775369f5c9f88f65e", "http://localhost:8080/auth/facebook/callback"),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := `
			<!DOCTYPE html>
			<html>
				<body>
				<h2>Log in with...</h2>
				<ul>
					<li>
					<a href="auth/github/login">GitHub</a>
					</li>
					<li>
					<a href="auth/google/login">Google</a>
					</li>
					<li>
					<a href="auth/facebook/login">Facebook</a>
					</li>
				</ul>
				</body>
			</html>
			`
		io.WriteString(w, template)
	})

	providers := []string{"google", "github", "facebook"}
	for _, provider := range providers {
		http.HandleFunc(fmt.Sprintf("/auth/%s/login", provider), loginHandler(provider))
		http.HandleFunc(fmt.Sprintf("/auth/%s/callback", provider), callbackHandler(provider))
	}

	http.ListenAndServe(Address, nil)

}

func loginHandler(providerName string) http.HandlerFunc {
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {

		state := gomniauth.NewState("after", "success")

		// This code borrowed from goweb example and not fixed.
		// if you want to request additional scopes from the provider,
		// pass them as login?scope=scope1,scope2
		//options := objx.MSI("scope", ctx.QueryValue("scope"))

		authUrl, err := provider.GetBeginAuthURL(state, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// redirect
		http.Redirect(w, r, authUrl, http.StatusFound)

	}
}

func callbackHandler(providerName string) http.HandlerFunc {
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {

		omap, err := objx.FromURLQuery(r.URL.RawQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		creds, err := provider.CompleteAuth(omap)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		/*
			// This code borrowed from goweb example and not fixed.
			// get the state
			state, err := gomniauth.StateFromParam(ctx.QueryValue("state"))

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// redirect to the 'after' URL
			afterUrl := state.GetStringOrDefault("after", "error?e=No after parameter was set in the state")

		*/

		// load the user
		user, userErr := provider.GetUser(creds)

		if userErr != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := fmt.Sprintf("%#v", user)
		io.WriteString(w, data)

		// redirect
		//return goweb.Respond.WithRedirect(ctx, afterUrl)

	}
}
