// Package log contains custom logger for our API
package log

import "net/http"

type responseWriter struct {
	http.ResponseWriter

	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
