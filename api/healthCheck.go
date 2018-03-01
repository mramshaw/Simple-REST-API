package api

import (
	"fmt"
	"net/http"
)

// HealthCheck implements a health check function suitable for use in Consul or Kubernetes.
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong\n")
}
