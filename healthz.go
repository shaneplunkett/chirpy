package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		log.Errorf("Health Write Failed: %v", err)
	}
	log.Info("Request on /healthz responded with OK")
}
