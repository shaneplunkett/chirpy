package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Swap(0)
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		log.Errorf("Reset Write Failed: %v", err)
	}
	log.Info("Metrics Reset")
}
