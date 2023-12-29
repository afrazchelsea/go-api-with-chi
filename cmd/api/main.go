package main

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi" // For web development
	"go-api/internal/handlers" // From our own module within the project
	log "github.com/sirupsen/logrus" // Log statements for error debugging
)
	
func main(){
	log.SetReportCaller(true) // Set to true to see the line number of the error
	var r *chi.Mux = chi.NewRouter() // Create a new router
	handlers.Handler(r) // Call the handler function from handlers.go

	fmt.Println("Starting GO API Service")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}