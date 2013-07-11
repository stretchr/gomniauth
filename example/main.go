package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/*

  This is a full web-stack example showing how to use gomniauth.

*/

var Address string = ":8080"

const (
	CookieNameSession string = "session"
)

func main() {

	authStore := new(ExampleAuthStore)
	authManager := gomniauth.NewManager(authStore,
		providers.Google("id", "secret", "http://localhost:8080/auth/google/login"),
		providers.Github("3d1e6ba69036e0624b61", "7e8938928d802e7582908a5eadaaaf22d64babf1", "http://localhost:8080/auth/github/callback", "user"))

	// GET /protected
	//
	// Attempts to get a protected resource.
	goweb.Map("GET", "/protected", func(c context.Context) error {

		sessionCookie, sessionCookieErr := c.HttpRequest().Cookie(CookieNameSession)

		// do they have a session ID?
		if sessionCookie == nil || sessionCookieErr != nil {
			// no - they will need to login
			goweb.Respond.WithRedirect(c, "login")
			return nil
		}

		// get the authSession
		authSession := authManager.WithID(sessionCookie.Value)

		// are they logged in?
		if !authSession.IsLoggedIn() {
			// no - they will need to login
			goweb.Respond.WithRedirect(c, "login")
			return nil
		}

		// the user is logged in, they now have access to the
		// resources.

		return goweb.Respond.With(c, http.StatusOK, []byte("You are successfully logged in."))

	})

	// GET /login
	//
	// Presents the user with a list of providers they can use to login.
	goweb.Map("GET", "/login", func(c context.Context) error {

		// create them a session ID
		sessionId := gomniauth.CreateSessionID()

		// get the authSession
		authSession := authManager.WithID(sessionId)

		c.HttpResponseWriter().Write([]byte(`
			<html>
				<head>
					<title>Login</title>
				</head>
				<body>
					<h2>Login</h2>
					<p>Select your method of login:</p>
					<ul>
		`))

		for providerName, provider := range authManager.Providers() {

			authUrl, _ := authSession.GetAuthURL(provider, sessionId)

			c.HttpResponseWriter().Write([]byte(`
						<li>
			`))
			c.HttpResponseWriter().Write([]byte(fmt.Sprintf("<a href=\"%s\">%s</a>", authUrl, providerName)))
			c.HttpResponseWriter().Write([]byte(`
						</li>
			`))
		}

		c.HttpResponseWriter().Write([]byte(`
					</ul>
				</body>
			</html>
		`))

		return nil

	})

	// GET /auth/oauth2/{provider}/login
	goweb.Map("/auth/{provider}/callback", func(c context.Context) error {

		provider, providerOk := authManager.Providers()[c.PathValue("provider")]

		if !providerOk {
			return errors.New("Unsupported provider")
		}

		state := gomniauth.StateFromRequest(provider.AuthType(), c.HttpRequest())
		returnedSessionId, sessionErr := gomniauth.IDFromState(provider.AuthType(), state)

		if sessionErr != nil {
			return sessionErr
		}

		// get the authSession
		authSession := authManager.WithID(returnedSessionId)
		callbackErr := authSession.HandleCallback(provider, returnedSessionId, c.HttpRequest())

		if callbackErr != nil {

			// save their session ID in the cookie
			c.HttpRequest().AddCookie(&http.Cookie{Name: CookieNameSession, Value: returnedSessionId})

		}

		return nil
	})

	/*

	   START OF WEB SERVER CODE

	*/

	log.Print("Goweb 2")
	log.Print("by Mat Ryer and Tyler Bunnell")
	log.Print(" ")
	log.Print("Starting Goweb powered server...")

	// make a http server using the goweb.DefaultHttpHandler()
	s := &http.Server{
		Addr:           Address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", Address)

	log.Printf("  visit: %s", Address)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", listenErr)
	}

	log.Println("")
	log.Println("Try some of these routes:")
	log.Printf("%s", goweb.DefaultHttpHandler())

	go func() {
		for _ = range c {

			// sig is a ^C, handle it

			// stop the HTTP server
			log.Print("Stopping the server...")
			listener.Close()

			/*
			   Tidy up and tear down
			*/
			log.Print("Tearing down...")

			// TODO: tidy code up here

			log.Fatal("Finished - bye bye.  ;-)")

		}
	}()

	// begin the server
	log.Fatalf("Error in Serve: %s", s.Serve(listener))

	/*

	   END OF WEB SERVER CODE

	*/

}

// ExampleAuthStore is an AuthStore that just keeps the Auth objects
// in memory.
type ExampleAuthStore struct {
	auths map[string]*gomniauth.Auth
}

func (s *ExampleAuthStore) GetAuth(id string) (*gomniauth.Auth, error) {
	return s.auths[id], nil
}
func (s *ExampleAuthStore) PutAuth(id string, auth *gomniauth.Auth) error {
	s.auths[id] = auth
	return nil
}
