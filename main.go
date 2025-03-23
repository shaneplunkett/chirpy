package main

import (
	"net/http"
	"sync/atomic"

	"github.com/charmbracelet/log"
)

type apiConfig struct {
	fileserverHits atomic.Int32
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
	mux.HandleFunc("GET /api/healthz", handlerHealthz)
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidate)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Infof("Server running from: %s on port: %s\n", filepath, port)
	log.Fatal(srv.ListenAndServe())
}
