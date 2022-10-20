package healthcheck

import (
	"fmt"
	"net/http"
)

// HealthCheck
// health check func
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "ok")
}
