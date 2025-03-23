package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		log.Info("Request Received")

		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := `<html>
      <body>
        <h1>Welcome, Chirpy Admin</h1>
        <p>Chirpy has been visited %d times!</p>
      </body>
    </html>`
	htmlContent := fmt.Sprintf(htmlTemplate, cfg.fileserverHits.Load())
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(fmt.Appendf(nil, htmlContent))
	if err != nil {
		log.Errorf("Metric Write Failed: %v", err)
	}
	log.Infof("Metrics Returned Successful Response, count = %v", cfg.fileserverHits.Load())
}
