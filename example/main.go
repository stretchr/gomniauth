package main

import (
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	Address string = ":9090"
)

// mapRoutes contains lots of examples of how to map things in
// Goweb.  It is in its own function so that test code can call it
// without having to run main().
func mapRoutes(g *gomniauth.Gomniauth) {

	/*
		Map a specific route that will redirect
	*/
	goweb.Map("GET", "~auth/[provider]/login", func(c context.Context) error {

		provider, exists := c.PathParams().Get("provider").(string)
		if !exists {
			return goweb.Respond.WithStatus(c, http.StatusInternalServerError)
		}

		return goweb.Respond.WithRedirect(c, g.RedirectURL(provider, fmt.Sprintf("State for %s", provider)))
	})

	/*
		/people (with optional ID)
	*/
	goweb.Map("GET", "~auth/[provider]/callback", func(c context.Context) error {

		fmt.Println("Callback fired!")

		provider, exists := c.PathParams().Get("provider").(string)
		if !exists {
			return goweb.Respond.WithStatus(c, http.StatusInternalServerError)
		}

		code := c.FormValue("code")
		state := c.FormValue("state")

		fmt.Printf("Got a redirect with state info: %s\n", state)

		g.Exchange(provider, code)

		response, err := g.Get(provider, "https://api.github.com/user")
		if err != nil {
			return goweb.Respond.WithStatus(c, http.StatusInternalServerError)
		}

		body, _ := ioutil.ReadAll(response.Body)
		return goweb.Respond.With(c, http.StatusOK, body)

	})

}

func main() {

	// set up gomniauth

	g := gomniauth.MakeGomniauth("http://localhost:9090/~auth/")
	g.AddProvider(gomniauth.Github, "3d1e6ba69036e0624b61", "7e8938928d802e7582908a5eadaaaf22d64babf1", "user")

	// map the routes
	mapRoutes(g)

	/*

	   START OF WEB SERVER CODE

	*/

	log.Print("Gomniauth Example Application")
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
	log.Print("Some things to try in your browser:")
	log.Printf("\t  http://localhost%s/~auth/github/login", Address)

	log.Println("")

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
