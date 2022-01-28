package main

import (
	"log"
	"net/http"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/kamp-us/graphql/internal/clients"
	"github.com/kamp-us/graphql/internal/handler"
	"github.com/kamp-us/graphql/internal/loader"
	"github.com/kamp-us/graphql/internal/resolver"
	"github.com/kamp-us/graphql/internal/schema"
	"go.uber.org/zap"
)

func main() {
	var (
		addr              = ":8000"
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)
	start := time.Now()

	log.SetFlags(log.LstdFlags | log.Llongfile)

	s, err := schema.String()
	if err != nil {
		log.Fatalf("reading embedded schema contents: %s", err)
	}

	c, err := clients.NewClients()
	if err != nil {
		log.Fatalf("initializing clients: %s", err)
	}

	root, err := resolver.NewRoot(c)
	if err != nil {
		log.Fatalf("creating root resolver: %s", err)
	}

	h := handler.GraphQL{
		Schema:  graphql.MustParseSchema(s, root),
		Loaders: loader.Initialize(nil),
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler.GraphiQL{})
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h)

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	go func() {
		log.Printf("Listening for requests on %s %v", srv.Addr, zap.Duration("elapsed", time.Since(start)))
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Println("server.ListenAndServe:", err)
	}

	// TODO: intercept shutdown signals for cleanup of connections.
	log.Println("Shut down.")
}
