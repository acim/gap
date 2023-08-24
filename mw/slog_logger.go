package mw

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
)

const Message = "request completed"

// SlogLogger is an idiomatic middleware to log requests using slog logger.
type SlogLogger struct {
	inner *slog.Logger
	next  http.Handler
}

// ServeHTTP implements http.Handler interface.
func (l *SlogLogger) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	writer := middleware.NewWrapResponseWriter(res, req.ProtoMajor)

	start := time.Now()

	l.next.ServeHTTP(writer, req)

	l.inner.Info(
		Message,
		"start", start,
		"method", req.Method,
		"uri", req.RequestURI,
		"status", writer.Status(),
		"bytes written", writer.BytesWritten(),
		"duration", time.Since(start))
}

func NewSlogLogger(next http.Handler, log *slog.Logger) *SlogLogger {
	return &SlogLogger{
		inner: log,
		next:  next,
	}
}
