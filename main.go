package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/charmbracelet/log"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		log.Info("Request Received")

		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(fmt.Appendf(nil, "Hits: %v", cfg.fileserverHits.Load()))
	if err != nil {
		log.Errorf("Metric Write Failed: %v", err)
	}
	log.Infof("Metrics Returned Successful Response, count = %v", cfg.fileserverHits.Load())
}

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

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		log.Errorf("Health Write Failed: %v", err)
	}
	log.Info("Request on /healthz responded with OK")
}

func main() {
	const filepath = "."
	const port = "8080"

	cfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	mux.Handle(
		"/app/",
		http.StripPrefix(
			"/app/",
			cfg.middlewareMetricsInc(http.FileServer(http.Dir(filepath))),
		),
	)
	mux.HandleFunc("/healthz", handlerHealthz)
	mux.HandleFunc("/metrics", cfg.handlerMetrics)
	mux.HandleFunc("/reset", cfg.handlerReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Infof("Server running from: %s on port: %s\n", filepath, port)
	log.Fatal(srv.ListenAndServe())
}
