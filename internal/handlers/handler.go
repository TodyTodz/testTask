package handlers

import (
	"encoding/json"
	"net/http"

	"testTask/pkg/handler_helpers"

	"github.com/Maldris/mathparse"

	. "github.com/NGRsoftlab/ngr-logging"
)

func CalculateMath(resp http.ResponseWriter, req *http.Request) {
	Logger.Info("starting calculate...")

	var sendData handler_helpers.Math
	err := json.NewDecoder(req.Body).Decode(&sendData)
	if err != nil {
		Logger.Error(err)
		handler_helpers.Respond(resp, map[string]interface{}{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	expression := sendData.Operation
	p := mathparse.NewParser(expression)

	p.Resolve()

	if p.FoundResult() {
		var result float64
		result = p.GetValueResult()
		Logger.Debug(result)

		handler_helpers.Respond(resp, map[string]interface{}{"res": result}, http.StatusOK)

	} else {
		handler_helpers.Respond(resp, map[string]interface{}{"error": "can't solve expression"}, http.StatusBadRequest)
	}

}
