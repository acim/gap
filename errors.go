package gap

import "net/http"

type Error interface {
	error
	Status() int
}

type BadRequestError struct {
	Err error
}

func (s BadRequestError) Error() string {
	if s.Err == nil {
		return http.StatusText(http.StatusBadRequest)
	}

	return s.Err.Error()
}

func (s BadRequestError) Status() int {
	return http.StatusBadRequest
}

type UnauthorizedError struct {
	Err error
}

func (s UnauthorizedError) Error() string {
	if s.Err == nil {
		return http.StatusText(http.StatusUnauthorized)
	}

	return s.Err.Error()
}

func (s UnauthorizedError) Status() int {
	return http.StatusUnauthorized
}

type ForbiddenError struct {
	Err error
}

func (s ForbiddenError) Error() string {
	if s.Err == nil {
		return http.StatusText(http.StatusForbidden)
	}

	return s.Err.Error()
}

func (s ForbiddenError) Status() int {
	return http.StatusForbidden
}

type NotFoundError struct {
	Err error
}

func (s NotFoundError) Error() string {
	if s.Err == nil {
		return http.StatusText(http.StatusNotFound)
	}

	return s.Err.Error()
}

func (s NotFoundError) Status() int {
	return http.StatusNotFound
}

type MethodNotAllowedError struct {
	Err error
}

func (s MethodNotAllowedError) Error() string {
	if s.Err == nil {
		return http.StatusText(http.StatusMethodNotAllowed)
	}

	return s.Err.Error()
}

func (s MethodNotAllowedError) Status() int {
	return http.StatusMethodNotAllowed
}

type InternalServerError struct {
	Err error
}

func (s InternalServerError) Error() string {
	if s.Err == nil {
		return http.StatusText(http.StatusInternalServerError)
	}

	return s.Err.Error()
}

func (s InternalServerError) Status() int {
	return http.StatusInternalServerError
}

type GenericError struct {
	StatusCode int
	Err        error
}

func (s GenericError) Error() string {
	if s.Err == nil {
		if st := http.StatusText(s.StatusCode); st != "" {
			return st
		}

		return http.StatusText(http.StatusTeapot)
	}

	return s.Err.Error()
}

func (s GenericError) Status() int {
	if s.StatusCode == 0 {
		return http.StatusTeapot
	}

	return s.StatusCode
}
