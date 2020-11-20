package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice" // New import
)

func (app *application) routes() http.Handler {
	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	router := mux.NewRouter()
	router.HandleFunc("/", app.TestHandler).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Return the 'standard' middleware chain followed by the servemux.
	return standardMiddleware.Then(router)
}
