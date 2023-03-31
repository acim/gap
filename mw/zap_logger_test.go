package mw_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.acim.net/gap/mw"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestZapLogger(t *testing.T) {
	t.Parallel()

	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()
	buf := &bytes.Buffer{}

	//nolint:exhaustruct
	log := zap.New(
		zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{}), zapcore.AddSync(buf), zapcore.DebugLevel),
		zap.ErrorOutput(zapcore.AddSync(buf)),
	)

	middleware := mw.NewZapLogger(handler, log)

	middleware.ServeHTTP(rec, req)

	res := rec.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("ServeHTTP() status code=%d; want %d", res.StatusCode, http.StatusOK)
	}

	var got zapRecord

	if err := json.NewDecoder(buf).Decode(&got); err != nil {
		t.Fatalf("decode json: %v", err)
	}

	got.isValid(t)
}

type zapRecord struct {
	StartTime    uint          `json:"start"`
	Method       string        `json:"method"`
	URI          string        `json:"uri"`
	StatusCode   int           `json:"status"`
	BytesWritten uint          `json:"bytes written"` //nolint:tagliatelle
	Duration     time.Duration `json:"duration"`
}

func (r *zapRecord) isValid(t *testing.T) {
	t.Helper()

	if r.StartTime <= 0 {
		t.Errorf("ServeHTTP() start time=%d; want >0", r.StartTime)
	}

	if r.Method != http.MethodGet {
		t.Errorf("ServeHTTP() method=%q; want %q", r.Method, http.MethodGet)
	}

	if r.URI != "/" {
		t.Errorf("ServeHTTP() uri=%q; want %q", r.URI, "")
	}

	if r.StatusCode != http.StatusOK {
		t.Errorf("ServeHTTP() status code=%d; want %d", r.StatusCode, http.StatusOK)
	}

	if r.BytesWritten != 0 {
		t.Errorf("ServeHTTP() bytes written=%d; want %d", r.BytesWritten, 0)
	}

	if r.Duration <= 0 {
		t.Error("ServeHTTP() duration=0; want > 0")
	}
}
