package pkg

import (
	"net/http"

	"github.com/alamin-mahamud/arya/api/pkg/utl"
)

// Ping - Responds Ok with 200 Status.
func Ping(w http.ResponseWriter, r *http.Request) {
	briefMessage := "PONG"
	payload := map[string]interface{}{
		"user": map[string]interface{}{
			"id":   123,
			"name": "ABC",
		},
	}
	utl.SendJSONSuccessResponse(w, http.StatusOK, briefMessage, payload)
}

// PingError - Responds Error with custom code Status.
func PingError(w http.ResponseWriter, r *http.Request) {
	briefMessage := "PONG ERROR"
	errors := map[string]interface{}{
		"code":              "0x0001",
		"message":           "Permission Denied",
		"developer_message": "Use Sudo",
	}
	utl.SendJSONErrResponse(w, http.StatusUnauthorized, briefMessage, errors)
}
