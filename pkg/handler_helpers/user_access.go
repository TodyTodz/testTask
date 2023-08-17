//File with User-Access wrapper
package handler_helpers

import (
	"context"
	"net/http"

	. "github.com/NGRsoftlab/ngr-logging"
)

const (
	UserAccessHeader = "User-Access"
	UserAccessValue = "superuser"
)

// Checking User-Access wrap.
func CheckUserAccess(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		userAccess, ok := request.Header[UserAccessHeader]
		if !ok {
			Logger.Error("Request have NO User-Access header")
			Respond(writer, map[string]interface{}{"error": http.StatusForbidden}, http.StatusForbidden)
			return
		} else if userAccess[0] != UserAccessValue {
			Logger.Error("Error: bad User-Access key")
			Respond(writer, map[string]interface{}{"error": http.StatusForbidden}, http.StatusForbidden)
			return
		}

		// add User-Access variables for future use (authentication etc.)
		request = request.WithContext(context.WithValue(request.Context(), UserAccessHeader, UserAccessValue))

		// run main handler for math
		h.ServeHTTP(writer, request)
	}
}
