package utl

import (
	"encoding/json"
	"net/http"
	"time"
)

// SendJSONSuccessResponse - ...
func SendJSONSuccessResponse(w http.ResponseWriter, code int, briefMessage string, payload interface{}) {
	SendJSONResponse(w, true, code, briefMessage, payload, nil)
}

// SendJSONResponse - ...
func SendJSONResponse(w http.ResponseWriter, success bool, code int, briefMessage string, payload interface{}, errors interface{}) {

	responseFormat := map[string]interface{}{
		"program":        applicationName,
		"version":        appicationVersion,
		"release":        applicationRelease,
		"datetime":       time.Now().Format(time.RFC3339),
		"unix_timestamp": int32(time.Now().Unix()),
		"success":        success,
		"code":           code,
		"message":        briefMessage,
	}

	if success {
		responseFormat["payload"] = payload
		responseFormat["errors"] = []interface{}{}
	} else {
		responseFormat["payload"] = []interface{}{}
		responseFormat["errors"] = errors
	}

	response, _ := json.Marshal(responseFormat)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// SendJSONErrResponse - ...
func SendJSONErrResponse(w http.ResponseWriter, code int, briefMessage string, errors interface{}) {
	/**

	Errors should be in following Format

	errors = []map[string]interface{}{
		{
			code: 0x0123,
			message: "short message",
			developer_message: "Message for developers"
		},
		{
			code: 0x0123,
			message: "short message",
			developer_message: "Message for developers"
		}
	}

	**/
	SendJSONResponse(w, false, code, briefMessage, nil, errors)
}
