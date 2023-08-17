package handler_helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	. "github.com/NGRsoftlab/ngr-logging"
)



//Recovering after panic in http handler.
func Recover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					Logger.Info("Make recovered error by default")
					err = errors.New("unknown error")
				}
				// we can use it for sending some errors
				Logger.Error(err.Error())
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(writer, request)
	})
}

//Logging requests.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Logger.Infof("Started %s %s", r.Method, r.RequestURI)

		start := time.Now()

		next.ServeHTTP(w, r)

		Logger.Infof("Request %s %s finished %v", r.Method, r.RequestURI, time.Now().Sub(start))
	})
}

//Common function for responding on http requests
func Respond(w http.ResponseWriter, data map[string]interface{}, stCode int) {
	w.Header().Add("Content-Type", "application/json")

	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods-Type", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	w.WriteHeader(stCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		if data != nil {
			Logger.Warning("RESPONSE ERROR: ", err, len(data))
		}

		// this line should be separated in such a way that it does not go beyond the margin
		err2 := json.NewEncoder(w).Encode(map[string]interface{}{"error": map[string]interface{}{"code": http.StatusBadRequest, "message": "no data"}})
		if err2 != nil {
			Logger.Warning("ERROR spare resp")
		}
	}
}


