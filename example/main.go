package main

import (
	"fmt"
	"github.com/stretchr/gomniauth"
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

	gomniauth.WithProviders()

	/*
	   GET /auth/{provider}/login

	   Redirects them to the fmtin page for the specified provider.
	*/
	goweb.Map("auth/{provider}/login", func(ctx context.Context) error {

		provider, err := gomniauth.Provider(ctx.PathParam("provider"))

		if err != nil {
			return err
		}

		authUrl, err := provider.GetBeginAuthURL(nil)

		if err != nil {
			return err
		}

		// redirect
		return goweb.Respond.WithRedirect(ctx, authUrl)

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
