package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var authMessageRe = regexp.MustCompile(`(?i)api[\s_-]?key|unauthor|forbidden|invalid token`)

// classifyUpgradeFailure inspects the HTTP response from a failed WebSocket
// upgrade and decides whether it represents an auth failure (-> *AuthError)
// or a generic connection failure (-> *ConnectionError).
//
// Twelve Data sometimes returns HTTP 200 with a JSON error body
// ({"code":401,"message":"...","status":"error"}) for an invalid API key, and
// sometimes 401/403. Both the status line and the body are inspected.
func classifyUpgradeFailure(resp *http.Response, body []byte) error {
	status := 0
	if resp != nil {
		status = resp.StatusCode
	}

	var parsed struct {
		Code    int    `json:"code"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	if len(body) > 0 {
		_ = json.Unmarshal(body, &parsed)
	}

	isAuth := false
	if status == http.StatusUnauthorized || status == http.StatusForbidden {
		isAuth = true
	}
	if parsed.Code == http.StatusUnauthorized || parsed.Code == http.StatusForbidden {
		isAuth = true
	}
	if parsed.Message != "" && authMessageRe.MatchString(parsed.Message) {
		isAuth = true
	}

	if isAuth {
		msg := parsed.Message
		if msg == "" {
			msg = "authentication failed"
		}
		return &AuthError{Message: msg}
	}

	msg := parsed.Message
	if msg == "" {
		if status > 0 {
			msg = fmt.Sprintf("websocket upgrade failed: HTTP %d", status)
		} else {
			msg = "websocket upgrade failed"
		}
	}
	return &ConnectionError{Message: msg}
}

// readBody reads at most n bytes from r and closes it. Used to snapshot the
// upgrade-failure response body without risking unbounded reads.
func readBody(r io.ReadCloser, n int64) []byte {
	if r == nil {
		return nil
	}
	defer r.Close()
	buf := &bytes.Buffer{}
	_, _ = io.Copy(buf, io.LimitReader(r, n))
	return buf.Bytes()
}
