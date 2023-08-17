package handlers

import (
	"net/http"

	"testTask/pkg/handler_helpers"

	"github.com/gorilla/mux"
)

//Http Router
func Router() *mux.Router {
	r := mux.NewRouter()

	// set recover func to router for basic fault-tolerance
	r.Use(handler_helpers.Recover)

	// set common logger function for all requests
	r.Use(handler_helpers.LogRequest)

	// handler for calculating mathematical expressions with User-Access wrapper
	r.HandleFunc(
		"/math",
		handler_helpers.CheckUserAccess(
			CalculateMath)).Methods(http.MethodPost)
	return r
}
