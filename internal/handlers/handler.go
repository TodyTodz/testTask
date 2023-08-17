package handlers

import (
	"encoding/json"
	"net/http"
	"testTask/pkg/handler_helpers"

	. "github.com/NGRsoftlab/ngr-logging"
)

func CalculateMath(resp http.ResponseWriter, req *http.Request) {
	Logger.Info("starting calculate...")

	var sendData handler_helpers.Math
	err := json.NewDecoder(req.Body).Decode(&sendData)
	if err != nil {
		Logger.Error(err)
		httphelpers.SendErrorResponse(resp, errorCustom.GlobalErrors.ErrBadUnmarshal(), http.StatusBadRequest)
		return
	}



}
