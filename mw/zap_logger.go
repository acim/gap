package mw

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// SlogLogger is an idiomatic middleware to log requests using zap logger.
type ZapLogger struct {
	inner *zap.Logger
	next  http.Handler
}

// ServeHTTP implements http.Handler interface.
func (l *ZapLogger) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	writer := middleware.NewWrapResponseWriter(res, req.ProtoMajor)

	start := time.Now()

	l.next.ServeHTTP(writer, req)

	l.inner.Info(
		Message,
		zap.Time("start", start),
		zap.String("method", req.Method),
		zap.String("uri", req.RequestURI),
		zap.Int("status", writer.Status()),
		zap.Int("bytes written", writer.BytesWritten()),
		zap.Duration("duration", time.Since(start)))
}

func NewZapLogger(next http.Handler, log *zap.Logger) *ZapLogger {
	return &ZapLogger{
		inner: log,
		next:  next,
	}
}
