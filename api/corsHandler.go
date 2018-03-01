package api

import (
	"net/http"
)

var allowedOrigin string

// HandleCORS is a function to handle CORS (Cross Origin Resource Sharing).
func HandleCORS(allowedHostPort string, next http.Handler) http.Handler {
	allowedOrigin = allowedHostPort
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		origin := req.Header.Get("Origin")
		if origin != "" {
			// define the host we will service
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			} else {
				return
			}
		}
		if req.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
			return
		}
		next.ServeHTTP(w, req)
	})
}
