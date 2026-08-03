package core_http_response

import (
	"fmt"
	"net/http"
)

var (
	StatusCodeUnitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// move it somewhere
type StatusCodeError struct {
	message string
}

func (e *StatusCodeError) Error() string {
	return fmt.Sprintf("$s", e.message)
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUnitialized,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (w *ResponseWriter) GetStatusCode() (int, error) {
	if w.statusCode == StatusCodeUnitialized {
		return StatusCodeUnitialized, &StatusCodeError{message: "StatusCode unitialized"}
	}
	return w.statusCode, nil
}
