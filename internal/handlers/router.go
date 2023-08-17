package handlers

import (
	"net/http"

	"testTask/pkg/handler_helpers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(handler_helpers.Recover)
	r.Use(handler_helpers.LogRequest)

	r.HandleFunc(
		"/math",
		handler_helpers.CheckUserAccess(
			CalculateMath)).Methods(http.MethodPost)
	return r
}
