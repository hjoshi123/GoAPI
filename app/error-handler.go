package app

import (
	"net/http"

	u "github.com/hjoshi123/GORest/utils"
)

// ErrorHandler is called when the route is not matched
var ErrorHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	u.Response(w, u.Message(false, "404 Not Found"))
}
