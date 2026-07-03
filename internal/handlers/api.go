package handlers

import (
	"goapi/internal/middleware"

	"github.com/go-chi/chi"
)

func Handler(r *chi.Mux) {
	// REMOVED r.Use(chimiddle.StripSlashes) from here!

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
