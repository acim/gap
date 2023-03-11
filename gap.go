package gap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Decode[T any](body io.ReadCloser, v *T) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	defer func() {
		err = errors.Join(err, body.Close())
	}()

	return nil
}

type Response[T any] struct {
	StatusCode int                           `json:"-"`
	Data       *T                            `json:"data,omitempty"`
	Err        string                        `json:"error,omitempty"`
	Log        func(msg string, args ...any) `json:"-"`
}

func (j Response[T]) Encode(w http.ResponseWriter) {
	if j.Log == nil {
		j.Log = logger
	}

	w.Header().Set("Content-Type", "application/json")

	if j.StatusCode > 0 {
		w.WriteHeader(j.StatusCode)
	}

	data, err := json.Marshal(j)
	if err != nil {
		j.Log("json encode", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := w.Write(data); err != nil {
		j.Log("write response", err)
	}
}

type HandlerFunc func(w http.ResponseWriter, req *http.Request) error

func Wrap(next HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := next(w, req); err != nil {
			var e Error
			if errors.As(err, &e) {
				Response[NoData]{StatusCode: e.Status(), Err: e.Error()}.Encode(w)

				return
			}

			Response[NoData]{
				StatusCode: http.StatusInternalServerError,
				Err:        http.StatusText(http.StatusInternalServerError),
			}.Encode(w)
		}
	})
}

type NoData struct{}

func logger(msg string, args ...any) {
	log.Println(append([]any{msg}, args...))
}
