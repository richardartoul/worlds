package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"crypto/worlds/server/glide_sucks/crypto/acme/autocert"

	"crypto/worlds/server/state"
)

func main() {
	// TODO: Config/Environment variables
	port := os.Args[1]
	connectionURL := os.Args[2]
	contractAddress := os.Args[3]

	stateManager := state.NewManager(connectionURL, contractAddress, 5*time.Second)
	err := stateManager.Init()
	if err != nil {
		log.Fatalf("Err initializing state manager: %v", err)
	}

	landingPage, err := template.ParseFiles("static/landing.html")
	if err != nil {
		log.Fatalf("Err parsiing landing page template: %v", err)
	}

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
