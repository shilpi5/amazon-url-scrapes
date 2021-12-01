package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// parse body
func ParseBody(r *http.Request, body interface{}) {
	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		json.NewDecoder(r.Body).Decode(&body)
	}
}

// application Json
func Json(str string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, str)
	}
}
