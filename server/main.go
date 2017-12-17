package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"crypto/worlds/server/config"
	"crypto/worlds/server/state"
)

func main() {
	config, err := config.Get()
	if err != nil {
		log.Fatalf("Err loading config: %v", err)
	}

	log.Printf("Starting server with config: %+v\n", config)

	stateManager := state.NewManager(
		config.EthereumClientAddresses,
		config.EthereumContractAddress,
		time.Duration(config.StateRefreshIntervalSeconds)*time.Second,
	)
	err = stateManager.Init()
	if err != nil {
		log.Fatalf("Err initializing state manager: %v", err)
	}

	landingPage, err := template.ParseFiles("static/landing.html")
	if err != nil {
		log.Fatalf("Err parsiing landing page template: %v", err)
	}

	// Setup background server to redirect http --> https
	sslRedirectMux := http.NewServeMux()
	sslRedirectMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// remove/add not default ports from r.Host
		target := "https://" + r.Host + r.URL.Path
		if len(r.URL.RawQuery) > 0 {
			target += "?" + r.URL.RawQuery
		}
		log.Printf("redirecting to: %s", target)
		http.Redirect(w, r, target, http.StatusTemporaryRedirect)
	})
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", config.HTTPPort), sslRedirectMux))
	}()

	// Setup primary server for serving landing page over HTTPS
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		landingPage.Execute(w, stateManager.Get())
	})

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	log.Fatal(
		http.Serve(
			autocert.NewListener(
				"worldsgreatesthuman.com", "biggestg.com",
			),
			mux,
		),
	)
}
