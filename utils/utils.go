package utils

import (
	"encoding/json"
	"net/http"
)

// Message functions builds a json message
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"Status": status, "Message": message}
}

// Response function returns a response
func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
