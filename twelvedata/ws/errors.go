package ws

import (
	"fmt"
	"time"
)

// TwelvedataWebSocketError is the marker interface satisfied by every typed
// error this package returns. Use errors.As against it to catch any of them;
// use errors.As with a concrete pointer (e.g. *AuthError) to handle a
// specific failure mode. The interface is sealed via an unexported method.
type TwelvedataWebSocketError interface {
	error
	twelvedataWebSocketError()
}

// AuthError is returned when the Twelve Data API key is missing or rejected
// by the server. AuthErrors are fatal — the auto-reconnect loop is suppressed.
type AuthError struct {
	Message string
}

func (e *AuthError) Error() string             { return "twelvedata websocket auth: " + e.Message }
func (*AuthError) twelvedataWebSocketError()   {}

// ConnectionError covers socket-level failures: failed upgrade (non-auth),
// failed send, unexpected non-auth server response. Cause holds the underlying
// network error when available.
type ConnectionError struct {
	Message string
	Cause   error
}

func (e *ConnectionError) Error() string {
	if e.Cause != nil {
		return "twelvedata websocket: " + e.Message + ": " + e.Cause.Error()
	}
	return "twelvedata websocket: " + e.Message
}
func (e *ConnectionError) Unwrap() error             { return e.Cause }
func (*ConnectionError) twelvedataWebSocketError()   {}

// TimeoutError is returned when no pong arrived within the configured
// timeout; the connection is treated as dead and a reconnect is triggered.
type TimeoutError struct {
	Timeout time.Duration
}

func (e *TimeoutError) Error() string {
	return fmt.Sprintf("twelvedata websocket: no pong within %s", e.Timeout)
}
func (*TimeoutError) twelvedataWebSocketError() {}

// ReconnectError is returned when the auto-reconnect loop has exhausted its
// configured maximum attempts. The most recent dial error is in Cause.
type ReconnectError struct {
	Attempts int
	Cause    error
}

func (e *ReconnectError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("twelvedata websocket: reconnect failed after %d attempts: %s", e.Attempts, e.Cause.Error())
	}
	return fmt.Sprintf("twelvedata websocket: reconnect failed after %d attempts", e.Attempts)
}
func (e *ReconnectError) Unwrap() error             { return e.Cause }
func (*ReconnectError) twelvedataWebSocketError()   {}
