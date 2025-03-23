package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {
	mux := http.NewServeMux()
	svr := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Infof("Server running on: %s", svr.Addr)
	err := svr.ListenAndServe()
	if err != nil {
		log.Errorf("Error serving connection: %s", err)
	}
}
