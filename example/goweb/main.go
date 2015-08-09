package main

import (
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/gomniauth/providers/uber"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	// NOTE: Don't change this, the auth settings on the providers
	// are coded to this path for this example.
	Address string = ":8080"
)

func write(ctx context.Context, output string) {
	ctx.HttpResponseWriter().Write([]byte(output))
}

func writeHeader(ctx context.Context) {
	write(ctx, "Gomniauth - Example web app")
}

func respondWithError(ctx context.Context, errorMessage string) error {
	writeHeader(ctx)
	write(ctx, fmt.Sprintf("Error: %s", errorMessage))
	return nil
}

func main() {

	// setup the providers
	gomniauth.SetSecurityKey("yLiCQYG7CAflDavqGH461IO0MHp7TEbpg6TwHBWdJzNwYod1i5ZTbrIF5bEoO3oP") // NOTE: DO NOT COPY THIS - MAKE YOR OWN!
	gomniauth.WithProviders(
		github.New("3d1e6ba69036e0624b61", "7e8938928d802e7582908a5eadaaaf22d64babf1", "http://localhost:8080/auth/github/callback"),
		google.New("1051709296778.apps.googleusercontent.com", "7oZxBGwpCI3UgFMgCq80Kx94", "http://localhost:8080/auth/google/callback"),
		facebook.New("537611606322077", "f9f4d77b3d3f4f5775369f5c9f88f65e", "http://localhost:8080/auth/facebook/callback"),
		uber.New("UBERKEY", "UBERSECRET", "http://localhost:8080/auth/uber/callback"),
	)

	goweb.Map("/", func(ctx context.Context) error {

		return goweb.Respond.With(ctx, http.StatusOK, []byte(`
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
             <li>
              <a href="auth/uber/login">Uber</a>
            </li>
          </ul>
        </body>
      </html>
    `))

	})

	/*
	   GET /auth/{provider}/login

	   Redirects them to the fmtin page for the specified provider.
	*/
	goweb.Map("auth/{provider}/login", func(ctx context.Context) error {

		provider, err := gomniauth.Provider(ctx.PathValue("provider"))

		if err != nil {
			return err
		}

		state := gomniauth.NewState("after", "success")

		// if you want to request additional scopes from the provider,
		// pass them as login?scope=scope1,scope2
		//options := objx.MSI("scope", ctx.QueryValue("scope"))

		authUrl, err := provider.GetBeginAuthURL(state, nil)

		if err != nil {
			return err
		}

		// redirect
		return goweb.Respond.WithRedirect(ctx, authUrl)

	})

	goweb.Map("auth/{provider}/callback", func(ctx context.Context) error {

		provider, err := gomniauth.Provider(ctx.PathValue("provider"))

		if err != nil {
			return err
		}

		creds, err := provider.CompleteAuth(ctx.QueryParams())

		if err != nil {
			return err
		}

		/*
			// get the state
			state, stateErr := gomniauth.StateFromParam(ctx.QueryValue("state"))

			if stateErr != nil {
				return stateErr
			}

			// redirect to the 'after' URL
			afterUrl := state.GetStringOrDefault("after", "error?e=No after parameter was set in the state")

		*/

		// load the user
		user, userErr := provider.GetUser(creds)

		if userErr != nil {
			return userErr
		}

		return goweb.API.RespondWithData(ctx, user)

		// redirect
		//return goweb.Respond.WithRedirect(ctx, afterUrl)

	})

	/*
	   ----------------------------------------------------------------
	   START OF WEB SERVER CODE
	   ----------------------------------------------------------------
	*/

	log.Println("Starting...")
	fmt.Print("Gomniauth - Example web app\n")
	fmt.Print("by Mat Ryer and Tyler Bunnell\n")
	fmt.Print(" \n")
	fmt.Print("Starting Goweb powered server...\n")

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

	fmt.Printf("  visit: %s\n", Address)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", listenErr)
	}

	fmt.Println("\n")
	fmt.Println("Try some of these routes:\n")
	fmt.Printf("%s", goweb.DefaultHttpHandler())
	fmt.Println("\n\n")

	go func() {
		for _ = range c {

			// sig is a ^C, handle it

			// stop the HTTP server
			fmt.Print("Stopping the server...\n")
			listener.Close()

			/*
			   Tidy up and tear down
			*/
			fmt.Print("Tearing down...\n")

			// TODO: tidy code up here

			log.Fatal("Finished - bye bye.  ;-)\n")

		}
	}()

	// begin the server
	log.Fatalf("Error in Serve: %s\n", s.Serve(listener))

	/*
	   ----------------------------------------------------------------
	   END OF WEB SERVER CODE
	   ----------------------------------------------------------------
	*/

}
