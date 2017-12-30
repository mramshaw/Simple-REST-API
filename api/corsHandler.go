package api

import (
    "net/http"
)

// HandleCORS is a function to handle CORS (Cross Origin Resource Sharing).
func HandleCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        origin := req.Header.Get("Origin")
        if origin != "" {
            // define the hosts we will service
            if origin == "http://localhost:3200" {
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
