package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {
	const filepath = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepath)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Infof("Server running from: %s on port: %s\n", filepath, port)
	log.Fatal(srv.ListenAndServe())
}
