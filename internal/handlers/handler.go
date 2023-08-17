package handlers

import (
	"encoding/json"
	"net/http"

	"testTask/pkg/handler_helpers"

	"github.com/Maldris/mathparse"

	. "github.com/NGRsoftlab/ngr-logging"
)

//CalculateMath - handler for calculating mathematical expressions and returning the result
func CalculateMath(resp http.ResponseWriter, req *http.Request) {
	Logger.Info("starting calculate...")

	// parse json data
	var sendData handler_helpers.Math
	err := json.NewDecoder(req.Body).Decode(&sendData)
	if err != nil {
		Logger.Error(err)
		handler_helpers.Respond(resp, map[string]interface{}{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	// use imported lib to parse expression
	expression := sendData.Operation
	p := mathparse.NewParser(expression)

	// calculate expression, but it's support * and / too..
	//  it's use regexp and i can:
	//   - fork from this lib for removing unnecessary support of * and / :)
	//   - or rewrite this lib more optimally if i have more time
	p.Resolve()

	// checking - is result calculated
	if p.FoundResult() {
		var result float64
		result = p.GetValueResult()

		// log result
		Logger.Debug(result)

		// and send 200 OK, json {"res":<result>}
		handler_helpers.Respond(resp, map[string]interface{}{"res": result}, http.StatusOK)

	} else {
		// send error because we can't calculated the result
		handler_helpers.Respond(resp, map[string]interface{}{"error": "can't solve expression"}, http.StatusBadRequest)
	}

}
