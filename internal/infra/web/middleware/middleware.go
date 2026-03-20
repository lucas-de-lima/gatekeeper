package middleware

import (
	"net/http"
	"time"

	"log/slog"
)

// RequestLogger logs basic request details.
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		slog.Info("incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_ip", r.RemoteAddr,
			"duration_ms", duration.Milliseconds(),
		)
	})
}
