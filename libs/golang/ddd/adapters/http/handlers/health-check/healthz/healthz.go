package healthz

import (
	"fmt"
	"net/http"
	"time"
)

// WebHealthzHandler handles HTTP requests for health checks,
// providing information about the server's uptime and readiness.
type WebHealthzHandler struct {
	startedAt    time.Time // The time the server started.
	timeProvider TimeProvider // The TimeProvider implementation for time-related functions.
}

// NewWebHealthzHandler creates and returns a new WebHealthzHandler instance
// with the specified TimeProvider. This allows for both real and mock time providers.
//
// Parameters:
//   - timeProvider: An implementation of the TimeProvider interface for time-related functions.
//
// Returns:
//   - A new instance of WebHealthzHandler.
func NewWebHealthzHandler(timeProvider TimeProvider) *WebHealthzHandler {
	return &WebHealthzHandler{
		startedAt:    timeProvider.Now(),
		timeProvider: timeProvider,
	}
}

// Healthz is an HTTP handler function that checks the health status of the server.
// If the server has been running for less than 5 seconds, it responds with a 500 Internal Server Error status.
// Otherwise, it responds with a 200 OK status.
//
// Parameters:
//   - w: The ResponseWriter to write the HTTP response.
//   - r: The HTTP request being handled.
//
// Returns:
//   - None.
//
// Example:
//
//   http.HandleFunc("/healthz", handler.Healthz)
//
func (h *WebHealthzHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	duration := h.timeProvider.Since(h.startedAt)
	if duration.Seconds() < 5 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Healthz check failed after %v seconds", duration.Seconds())))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healthz check passed"))
	}
}