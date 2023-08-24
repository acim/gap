package mw_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.acim.net/gap/mw"
)

func TestSlogLogger(t *testing.T) {
	t.Parallel()

	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()
	buf := &bytes.Buffer{}

	log := slog.New(slog.NewJSONHandler(buf, nil))
	middleware := mw.NewSlogLogger(handler, log)

	middleware.ServeHTTP(rec, req)

	res := rec.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("ServeHTTP() status code=%d; want %d", res.StatusCode, http.StatusOK)
	}

	var got slogRecord

	if err := json.NewDecoder(buf).Decode(&got); err != nil {
		t.Fatalf("decode json: %v", err)
	}

	got.isValid(t)
}

type slogRecord struct {
	Time         *time.Time    `json:"time"`
	Level        string        `json:"level"`
	Message      string        `json:"msg"`
	StartTime    *time.Time    `json:"start"`
	Method       string        `json:"method"`
	URI          string        `json:"uri"`
	StatusCode   int           `json:"status"`
	BytesWritten uint          `json:"bytes written"` //nolint:tagliatelle
	Duration     time.Duration `json:"duration"`
}

func (r *slogRecord) isValid(t *testing.T) {
	t.Helper()

	if r.StartTime.After(*r.Time) {
		t.Errorf("ServeHTTP() log time=%q; start time %q", r.Time, r.StartTime)
	}

	if r.Level != slog.LevelInfo.String() {
		t.Errorf("ServeHTTP() log level=%q; want %q", r.Level, slog.LevelInfo.String())
	}

	if r.Message != mw.Message {
		t.Errorf("ServeHTTP() message=%q; want %q", r.Message, mw.Message)
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
