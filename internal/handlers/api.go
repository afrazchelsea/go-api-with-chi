package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"go-api/internal/middleware"
	
)

// If  a function starts with a capital letter, it is exported and can be used by other packages.
// If a function starts with a lowercase letter, it is private and can only be used within the same package.

func Handler(r *chi.Mux){
	// Global middlwares
	r.Use(chimiddle.StripSlashes)

	// Routes
	r.Route("/api", func(router chi. Router){
		// Middleware to check if authorizaed
		router.Use(middleware.Authorization)
		// GET /api/coins
		router.Get("/coins", GetCoinBalance)
	})
}
