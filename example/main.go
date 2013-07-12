package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/providers"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/stew/objects"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

/*

  This is a full web-stack example showing how to use gomniauth.

	To run this app, just do:

	    clear; go build main.go; sudo ./main; rf main;

	NOTE: The domain is set to www.localhost.com, so you should route this
	      properly by editing your /etc/hosts file.


	Credit:

		The auth buttons were made by this guy:

			http://nicolasgallagher.com/lab/css3-social-signin-buttons/

*/

var Address string = ":80"

const (
	CookieNameSession string = "session"
)

// RedirectToLoginPage responds WithRedirect to the login page,
// passing the current page as the 'target' parameter.
var RedirectToLoginPage = func(c context.Context) error {
	redirectUrl := "login?target=" + c.HttpRequest().URL.Path
	log.Println("Example app: Redirecting to login page: %s", redirectUrl)
	return goweb.Respond.WithRedirect(c, redirectUrl)
}

func main() {

	/*
		Step 1. Create an AuthStore

			- An AuthStore is responsible for caching your auth tokens so you
			  don't have to send the user off to reauthenticate every time you
			  want to access their remote data.

	*/
	authStore := &ExampleAuthStore{make(map[string]*common.Auth)}

	/*
		Step 2. Create and configure an AuthManager

		  - Give it the authStore you created earlier
		  - Give it all the providers you wish to support, remember you'll
		    have to configure each provider for your application by specifying
		    Client ID, secrets, callback URLs and scopes etc. depending on the
		    auth type.

	*/
	authManager := gomniauth.NewManager(authStore,
		providers.Google("815669121291.apps.googleusercontent.com", "QrjJ2WevjIp1CbJxU18449RS", "http://www.localhost.com/auth/google/callback", "profile"),
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

		log.Print("Example app: User tried to access protected resource")

		sessionCookie, sessionCookieErr := c.HttpRequest().Cookie(CookieNameSession)

		log.Printf("Example app: Cookies: %s", c.HttpRequest().Cookies())

		// do they have a session ID?
		if sessionCookie == nil || sessionCookieErr != nil {
			// no - they will need to login
			return RedirectToLoginPage(c)
		}

		segs := strings.Split(sessionCookie.Value, ":")
		providerName := segs[0]
		sessionId := segs[1]

		session := authManager.NewSession(sessionId, authManager.Provider(providerName))
		if authed, _ := session.IsAuthenticated(); !authed {
			// no - they will need to login
			return RedirectToLoginPage(c)
		}

		output := `

			<html>
				<head>
					<title>Welcome</title>
				</head>
				<body>
					<h1>Protected Resource</h1>
					<h3>
						You have successfully accessed the protected resource.
					</h3>
					<p>
						Here's a bit about you:
						<table border="1" bordercolor="#000000" style="background-color:#EAEAEA" width="100%" cellpadding="3" cellspacing="2">
							<tr>
								<td>Key</td>
								<td>Value</td>
							</tr>
							<tr><td></td><td></td></tr>
							$$$
						</table>
					</p>
					<p></p>
					<ul>
						<li>
							Now you can <a href="/logout">log out</a>
						</li>
						<li>
							Or you can <a href="/deauth">expire the internal auth token</a>
						</li>
					</ul>
				</body>
			</html>

		`

		client, err := session.AuthenticatedClient()

		if err != nil {
			//TODO: make this not bad
			return goweb.Respond.WithStatus(c, http.StatusInternalServerError)
		}

		url := ""

		switch session.Provider().Name() {
		case "Github":
			url = "https://api.github.com/user"
		case "Google":
			//url = "https://www.googleapis.com/plus/v1/people/me"
			url = "https://www.googleapis.com/oauth2/v3/userinfo"
		}

		resp, err := client.Get(url)
		if err != nil {
			//TODO: make this not bad
			return goweb.Respond.WithStatus(c, http.StatusInternalServerError)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		var userData map[string]interface{}
		json.Unmarshal(body, &userData)

		replacement := ""
		for k, v := range userData {
			replacement += fmt.Sprintf("<tr><td>%s</td><td>%v</td></tr>", k, v)
		}

		output = strings.Replace(output, "$$$", replacement, 1)

		return goweb.Respond.With(c, http.StatusOK, []byte(output))
	})

	goweb.Map("GET", "/logout", func(c context.Context) error {

		log.Printf("Example app: Logging out")

		// delete all cookies
		for _, cookie := range c.HttpRequest().Cookies() {
			cookie.MaxAge = -1
			cookie.Value = ""
			cookie.Path = "/"
			http.SetCookie(c.HttpResponseWriter(), cookie)
		}

		return goweb.Respond.WithRedirect(c, "/")

	})

	goweb.Map("GET", "/deauth", func(c context.Context) error {

		log.Printf("Example app: Forgetting user (deleting their auth token in the AuthStore)")

		sessionCookie, sessionCookieErr := c.HttpRequest().Cookie(CookieNameSession)

		// do they have a session ID?
		if sessionCookie == nil || sessionCookieErr != nil {
			// no - they will need to login
			return RedirectToLoginPage(c)
		}

		segs := strings.Split(sessionCookie.Value, ":")
		sessionId := segs[1]

		authManager.AuthStore().DeleteAuth(sessionId)

		// send them home
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
		authSession := authManager.NewSession(sessionId, nil)

		c.HttpResponseWriter().Write([]byte(`
			<html>
				<head>
					<title>Login</title>
					<link rel="stylesheet" href="/assets/authbuttons/auth-buttons.css">
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
						<li style='padding:5px;list-item:none'>
			`))
			c.HttpResponseWriter().Write([]byte(fmt.Sprintf("<a class='btn-auth btn-%s' href=\"%s\">Sign in with %s</a>", providerName, authUrl, provider.Name())))
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
		authSession := authManager.NewSession(returnedSessionId, provider)
		callbackErr := authSession.HandleCallback(c.HttpRequest())

		if callbackErr != nil {

			// save their session ID in the cookie
			cookie := &http.Cookie{Name: CookieNameSession,
				Value: fmt.Sprintf("%s:%s", provider.Name(), returnedSessionId),
				Path:  "/",
			}
			http.SetCookie(c.HttpResponseWriter(), cookie)

			log.Printf("Example app: User has been logged in with session ID: %s", returnedSessionId)
			log.Printf("  Set cookie: %s", cookie)

			targetUrl := gomniauth.TargetURLFromState(provider.AuthType(), state)

			if len(targetUrl) == 0 {
				targetUrl = "/"
			}

			goweb.Respond.WithRedirect(c, targetUrl)

		}

		return nil
	})

	// expose the assets too
	goweb.MapStatic("assets", "assets")

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
	log.Println("\n\n")

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
	auths map[string]*common.Auth
}

func (s *ExampleAuthStore) GetAuth(id string) (*common.Auth, error) {
	log.Printf("ExampleAuthStore: GetAuth %s", id)
	log.Printf("  returning: %v", s.auths[id])
	return s.auths[id], nil
}
func (s *ExampleAuthStore) PutAuth(id string, auth *common.Auth) error {
	log.Printf("ExampleAuthStore: PutAuth %s", id)
	log.Printf("  putting: %v", auth)
	s.auths[id] = auth
	return nil
}
func (s *ExampleAuthStore) DeleteAuth(id string) error {
	log.Printf("ExampleAuthStore: DeleteAuth: %s", id)
	delete(s.auths, id)
	return nil
}
