package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TwelvedataApiError is the marker interface satisfied by *ApiError and every
// status-specific error returned by this client. Use errors.As with this
// interface to detect any Twelve Data API error; use errors.As with a concrete
// pointer type (e.g. *UnauthorizedError) to handle a specific HTTP status.
type TwelvedataApiError interface {
	error
	GetStatusCode() int
	GetCode() int64
	GetStatus() string
	GetMessage() string
}

// ApiError is the base type returned for any non-2xx Twelve Data API response
// with a recognizable error body. Status-specific errors embed it.
type ApiError struct {
	StatusCode int
	Code       int64
	Status     string
	Message    string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("twelvedata API error [%d]: %s", e.StatusCode, e.Message)
}

func (e *ApiError) GetStatusCode() int { return e.StatusCode }
func (e *ApiError) GetCode() int64     { return e.Code }
func (e *ApiError) GetStatus() string  { return e.Status }
func (e *ApiError) GetMessage() string { return e.Message }

type BadRequestError struct{ ApiError }
type UnauthorizedError struct{ ApiError }
type ForbiddenError struct{ ApiError }
type NotFoundError struct{ ApiError }
type ParameterTooLongError struct{ ApiError }
type TooManyRequestsError struct{ ApiError }
type InternalServerError struct{ ApiError }

// errorMap dispatches an HTTP status code to a constructor of the matching
// typed error. One-to-one for now; the function-valued shape leaves room for
// future logic that branches on response-body fields. Keep this mirrored with
// the Api*ErrorResponseBody models the generator emits.
var errorMap = map[int]func(ApiError) error{
	400: func(a ApiError) error { return &BadRequestError{a} },
	401: func(a ApiError) error { return &UnauthorizedError{a} },
	403: func(a ApiError) error { return &ForbiddenError{a} },
	404: func(a ApiError) error { return &NotFoundError{a} },
	414: func(a ApiError) error { return &ParameterTooLongError{a} },
	429: func(a ApiError) error { return &TooManyRequestsError{a} },
	500: func(a ApiError) error { return &InternalServerError{a} },
}

// errorResponseBody is the shared shape of every Twelve Data error response
// body. Pointer fields let us tell "field absent" apart from "field present
// with the zero value" — required to safely detect responses we do not own.
type errorResponseBody struct {
	Code    *int64  `json:"code"`
	Message *string `json:"message"`
	Status  *string `json:"status"`
}

func (b *errorResponseBody) valid() bool {
	return b != nil && b.Code != nil && b.Message != nil && b.Status != nil
}

type errorRoundTripper struct {
	base http.RoundTripper
}

func (t *errorRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.base.RoundTrip(req)
	if err != nil || resp == nil {
		return resp, err
	}
	if resp.StatusCode < 400 {
		return resp, nil
	}

	buf, readErr := io.ReadAll(resp.Body)
	resp.Body.Close()
	if readErr != nil {
		resp.Body = io.NopCloser(bytes.NewReader(nil))
		return resp, nil
	}

	var body errorResponseBody
	if jsonErr := json.Unmarshal(buf, &body); jsonErr != nil || !body.valid() {
		resp.Body = io.NopCloser(bytes.NewReader(buf))
		return resp, nil
	}

	apiErr := ApiError{
		StatusCode: resp.StatusCode,
		Code:       *body.Code,
		Status:     *body.Status,
		Message:    *body.Message,
	}
	if resolver, ok := errorMap[resp.StatusCode]; ok {
		return nil, resolver(apiErr)
	}
	return nil, &apiErr
}

// WrapTransport returns an http.RoundTripper that translates non-2xx Twelve
// Data API responses with a recognizable error body into typed errors. Pass
// nil to wrap http.DefaultTransport. Users with a custom *http.Client should
// wrap their own transport so error handling stays installed.
func WrapTransport(base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &errorRoundTripper{base: base}
}
