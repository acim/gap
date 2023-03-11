package gap_test

import (
	"errors"
	"net/http"
	"testing"

	"go.acim.net/gap"
)

//nolint:exhaustruct,goerr113,funlen
func TestCommonErrors(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		err        gap.Error
		wantStatus int
		wantError  string
	}{
		"BadRequestError": {
			err:        gap.BadRequestError{},
			wantStatus: http.StatusBadRequest,
			wantError:  http.StatusText(http.StatusBadRequest),
		},
		"BadRequestError-custom": {
			err:        gap.BadRequestError{Err: errors.New("custom")},
			wantStatus: http.StatusBadRequest,
			wantError:  "custom",
		},
		"UnauthorizedError": {
			err:        gap.UnauthorizedError{},
			wantStatus: http.StatusUnauthorized,
			wantError:  http.StatusText(http.StatusUnauthorized),
		},
		"UnauthorizedError-custom": {
			err:        gap.UnauthorizedError{Err: errors.New("custom")},
			wantStatus: http.StatusUnauthorized,
			wantError:  "custom",
		},
		"ForbiddenError": {
			err:        gap.ForbiddenError{},
			wantStatus: http.StatusForbidden,
			wantError:  http.StatusText(http.StatusForbidden),
		},
		"ForbiddenError-custom": {
			err:        gap.ForbiddenError{Err: errors.New("custom")},
			wantStatus: http.StatusForbidden,
			wantError:  "custom",
		},
		"NotFoundError": {
			err:        gap.NotFoundError{},
			wantStatus: http.StatusNotFound,
			wantError:  http.StatusText(http.StatusNotFound),
		},
		"NotFoundError-custom": {
			err:        gap.NotFoundError{Err: errors.New("custom")},
			wantStatus: http.StatusNotFound,
			wantError:  "custom",
		},
		"MethodNotAllowedError": {
			err:        gap.MethodNotAllowedError{},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  http.StatusText(http.StatusMethodNotAllowed),
		},
		"MethodNotAllowedError-custom": {
			err:        gap.MethodNotAllowedError{Err: errors.New("custom")},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  "custom",
		},
		"InternalServerError": {
			err:        gap.InternalServerError{},
			wantStatus: http.StatusInternalServerError,
			wantError:  http.StatusText(http.StatusInternalServerError),
		},
		"InternalServerError-custom": {
			err:        gap.InternalServerError{Err: errors.New("custom")},
			wantStatus: http.StatusInternalServerError,
			wantError:  "custom",
		},
		"GenericError": {
			err:        gap.GenericError{},
			wantStatus: http.StatusTeapot,
			wantError:  http.StatusText(http.StatusTeapot),
		},
		"GenericError-custom": {
			err: gap.GenericError{
				StatusCode: http.StatusNotImplemented,
				Err:        errors.New("custom"),
			},
			wantStatus: http.StatusNotImplemented,
			wantError:  "custom",
		},
		"GenericError-customStatus": {
			err: gap.GenericError{
				StatusCode: http.StatusNotImplemented,
			},
			wantStatus: http.StatusNotImplemented,
			wantError:  http.StatusText(http.StatusNotImplemented),
		},
		"GenericError-unknown": {
			err: gap.GenericError{
				StatusCode: -1,
			},
			wantStatus: -1,
			wantError:  http.StatusText(http.StatusTeapot),
		},
	}

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if gotStatus := test.err.Status(); gotStatus != test.wantStatus {
				t.Errorf("Status()=%d; want %d", gotStatus, test.wantStatus)
			}

			if gotError := test.err.Error(); gotError != test.wantError {
				t.Errorf("Error()=%s; want %s", gotError, test.wantError)
			}
		})
	}
}
