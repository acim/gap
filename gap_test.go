package gap_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.acim.net/gap"
)

//nolint:exhaustruct,errcheck
func ExampleHandler_gap() {
	// This handler decodes user from JSON payload and again encodes it in the response.
	handler := gap.Wrap(func(w http.ResponseWriter, req *http.Request) error {
		var user User

		if err := gap.Decode(req.Body, &user); err != nil {
			return gap.BadRequestError{
				Err: err,
			}
		}

		gap.Response[User]{Data: &user}.Encode(w)

		return nil
	})

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	srv := &http.Server{
		Addr:              ":5000",
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}

	// This goroutine is used just to stop the server when running tests.
	go func() {
		time.Sleep(1 * time.Second)
		srv.Shutdown(context.Background())
	}()

	srv.ListenAndServe()

	// Output:
}

//nolint:exhaustruct,errcheck
func ExampleHandler_std() {
	// This handler decodes user from JSON payload and again encodes it in the response.
	handler := func(w http.ResponseWriter, req *http.Request) {
		var user User

		if err := gap.Decode(req.Body, &user); err != nil {
			gap.Response[gap.NoData]{
				Err: err.Error(),
			}.Encode(w)

			return
		}

		gap.Response[User]{Data: &user}.Encode(w)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	srv := &http.Server{
		Addr:              ":5000",
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
	}

	// This goroutine is used just to stop the server when running tests.
	go func() {
		time.Sleep(1 * time.Second)
		srv.Shutdown(context.Background())
	}()

	srv.ListenAndServe()

	// Output:
}

func TestEncode(t *testing.T) {
	t.Parallel()

	h := gap.Wrap(userHandler)

	want := User{Name: "alice", Address: "no name street 0"}

	data, err := json.Marshal(&want)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	w := httptest.NewRecorder()

	h(w, req)

	res := w.Result()

	defer res.Body.Close()

	var got gap.Response[User]

	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if *got.Data != want {
		t.Errorf("want %+v; got %+v", want, *got.Data)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("want status code 200; got %d", res.StatusCode)
	}

	if res.Header.Get("Content-Type") != "application/json" {
		t.Errorf("want content type application/json; got %s", res.Header.Get("Content-Type"))
	}
}

//nolint:goerr113,exhaustruct
func TestWrap(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		err  error
		want gap.Response[gap.NoData]
	}{
		"gap-error": {
			err: gap.InternalServerError{},
			want: gap.Response[gap.NoData]{
				StatusCode: http.StatusInternalServerError,
				Err:        gap.InternalServerError{}.Error(),
			},
		},
		"non-gap-error": {
			err: errors.New("unknown error"),
			want: gap.Response[gap.NoData]{
				StatusCode: http.StatusInternalServerError,
				Err:        http.StatusText(http.StatusInternalServerError),
			},
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			hf := gap.Wrap(func(w http.ResponseWriter, req *http.Request) error {
				return test.err
			})

			hf(w, req)

			res := w.Result()

			defer res.Body.Close()

			var got gap.Response[gap.NoData]

			if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
				t.Fatal(err)
			}

			if res.StatusCode != test.want.StatusCode {
				t.Errorf("want status code %d; got %d", test.want.StatusCode, res.StatusCode)
			}

			if got.Err != test.want.Err {
				t.Errorf("want error %v; got %v", test.want.Err, got.Err)
			}
		})
	}
}

type User struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func userHandler(w http.ResponseWriter, req *http.Request) error {
	var user User

	if err := gap.Decode(req.Body, &user); err != nil {
		return gap.BadRequestError{
			Err: err,
		}
	}

	gap.Response[User]{Data: &user}.Encode(w)

	return nil
}
