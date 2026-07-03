package main

import (
	"fmt"
	"net/http"

	"goapi/internal/handlers"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

// CORSMiddleware ensures that the browser does not block the webpage requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

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

	// STEP 1: Register ALL global middleware first
	r.Use(CORSMiddleware)
	r.Use(chimiddle.StripSlashes) // Moved here to satisfy Chi's strict order!

	// STEP 2: Define your static home route next
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	// STEP 3: Register your external handlers package routes
	handlers.Handler(r)

	// STEP 4: Start the server
	fmt.Println("Starting GO API service on http://localhost:8000...")
	err := http.ListenAndServe("0.0.0.0:8000", r)
	if err != nil {
		log.Error(err)
	}
}
