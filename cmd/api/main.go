package main

import (
	"fmt"
	"net/http"

	"goapi/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// CORSMiddleware ensures that the browser does not block the webpage requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle browser preflight requests immediately
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()

	// 1. Inject the CORS middleware before loading your handlers
	r.Use(CORSMiddleware)

	handlers.Handler(r)

	fmt.Println("Starting GO API service on http://localhost:8000...")

	// 1. Inject the CORS middleware before loading your handlers
	r.Use(CORSMiddleware)

	// NEW: Serve index.html as the home page directly from the container
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	handlers.Handler(r)

	fmt.Println("Starting GO API service on http://localhost:8000...")

	// 2. Changed from "localhost:8000" to "0.0.0.0:8000" to reliably accept local requests
	err := http.ListenAndServe("0.0.0.0:8000", r)
	if err != nil {
		log.Error(err)
	}
}
