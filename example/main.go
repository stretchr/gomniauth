package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/stew/objects"
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

var Address string = ":80"

const (
	CookieNameSession string = "session"
)

func main() {

	RedirectToLoginPage := func(c context.Context) error {
		log.Println("Redirecting to login page...")
		return goweb.Respond.WithRedirect(c, "login?target="+c.HttpRequest().URL.Path)
	}

	authStore := new(ExampleAuthStore)
	authManager := gomniauth.NewManager(authStore,
		providers.Google("id", "secret", "http://www.localhost.com/auth/google/callback"),
		providers.Github("3d1e6ba69036e0624b61", "7e8938928d802e7582908a5eadaaaf22d64babf1", "http://www.localhost.com/auth/github/callback", "user"))

	// TODO: make the callback dynamic

	goweb.Map("GET", "/", func(c context.Context) error {

		return goweb.Respond.With(c, http.StatusOK, []byte(`

			<html>
				<head>
					<title>Welcome</title>
				</head>
				<body>
					<h1>Omniauth Example Web app</h1>
					<p>
					</p>
					<ul>
						<li>
							To get started, just try accessing this <a href="/protected">protected resource</a>.
						</li>
					</ul>
				</body>
			</html>

		`))

	})

	// GET /protected
	//
	// Attempts to get a protected resource.
	goweb.Map("GET", "/protected", func(c context.Context) error {

		log.Print("User tried to access protected resource")

		sessionCookie, sessionCookieErr := c.HttpRequest().Cookie(CookieNameSession)

		log.Printf("%s", c.HttpRequest().Cookies())
		log.Printf("Cookie: %s\nCookie Error: %s", sessionCookie, sessionCookieErr)

		// do they have a session ID?
		if sessionCookie == nil || sessionCookieErr != nil {
			// no - they will need to login
			return RedirectToLoginPage(c)
		}

		return goweb.Respond.With(c, http.StatusOK, []byte(`

			<html>
				<head>
					<title>Welcome</title>
				</head>
				<body>
					<h1>Protected Resource</h1>
					<p>
					</p>
					<ul>
						<li>
							You have successfully accessed the protected resource.
						</li>
						<li>
							Now you can <a href="/logout">log out</a>
						</li>
					</ul>
				</body>
			</html>

		`))
	})

	goweb.Map("GET", "/logout", func(c context.Context) error {

		log.Printf("Logging out: %s", c.HttpRequest().Cookies())

		// delete all cookies
		for _, cookie := range c.HttpRequest().Cookies() {
			cookie.MaxAge = -1
			cookie.Value = ""
			cookie.Path = "/"
			cookie.Domain = "www.localhost.com"
			http.SetCookie(c.HttpResponseWriter(), cookie)
		}

		return goweb.Respond.WithRedirect(c, "/")

	})

	// GET /login
	//
	// Presents the user with a list of providers they can use to login.
	goweb.Map("GET", "/login", func(c context.Context) error {

		// create them a session ID
		sessionId := gomniauth.CreateSessionID()

		// get the target URL (URL to redirect to after they've logged in)
		targetUrl := c.QueryValue("target")

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

			// TODO: Make StateWith(id, target) or similar
			authUrl, _ := authSession.GetAuthURL(provider, objects.NewMap("id", sessionId, "targetUrl", targetUrl))

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

	// GET /auth/oauth2/{provider}/callback
	goweb.Map("/auth/{provider}/callback", func(c context.Context) error {

		provider, providerOk := authManager.Providers()[c.PathValue("provider")]

		if !providerOk {
			return errors.New("Unsupported provider")
		}

		state, stateErr := gomniauth.StateFromRequest(provider.AuthType(), c.HttpRequest())

		if stateErr != nil {
			return stateErr
		}

		returnedSessionId, sessionErr := gomniauth.IDFromState(provider.AuthType(), state)

		if sessionErr != nil {
			return sessionErr
		}

		// get the authSession
		authSession := authManager.WithID(returnedSessionId)
		callbackErr := authSession.HandleCallback(provider, returnedSessionId, c.HttpRequest())

		if callbackErr != nil {

			// save their session ID in the cookie
			cookie := &http.Cookie{Name: CookieNameSession,
				Value:  returnedSessionId,
				Path:   "/",
				Domain: "www.localhost.com",
			}
			http.SetCookie(c.HttpResponseWriter(), cookie)

			log.Printf("User has been logged in with session ID: %s", returnedSessionId)
			log.Printf("Set cookie: %s", cookie)

			targetUrl := gomniauth.TargetURLFromState(provider.AuthType(), state)

			if len(targetUrl) == 0 {
				targetUrl = "/"
			}

			goweb.Respond.WithRedirect(c, targetUrl)

			/*
				c.HttpResponseWriter().Write([]byte(`
					<html>
						<head>
				`))
				c.HttpResponseWriter().Write([]byte("<meta-equiv name=\"refresh\" content=\"0;URL='" + targetUrl + "\">"))
				c.HttpResponseWriter().Write([]byte(`
						</head>
						<body>
				`))

				c.HttpResponseWriter().Write([]byte("<a href='" + targetUrl + "'>Click here</a> to continue."))

				c.HttpResponseWriter().Write([]byte(`
						</body>
					</html>
				`))
			*/

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
