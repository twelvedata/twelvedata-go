package ws

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

// Client is the Twelve Data real-time WebSocket client. Construct with
// NewClient, then Connect, then Subscribe. Event channels are accessible via
// the Prices / SubscribeStatuses / Heartbeats / Messages / Errors /
// Reconnecting / Reconnected accessors.
type Client struct {
	apiKey            string
	url               string
	heartbeatInterval time.Duration
	pingInterval      time.Duration
	pingTimeout       time.Duration
	reconnectOpts     ReconnectOptions
	headers           http.Header

	prices       chan PriceEvent
	statuses     chan SubscribeStatusEvent
	heartbeats   chan HeartbeatEvent
	messages     chan json.RawMessage
	errs         chan error
	reconnecting chan ReconnectingEvent
	reconnected  chan ReconnectedEvent

	droppedPrices       atomic.Int64
	droppedStatuses     atomic.Int64
	droppedHeartbeats   atomic.Int64
	droppedMessages     atomic.Int64
	droppedErrors       atomic.Int64
	droppedReconnecting atomic.Int64
	droppedReconnected  atomic.Int64

	mu               sync.Mutex
	conn             *websocket.Conn
	connectingDone   chan struct{}
	connectErr       error
	hasEverConnected bool
	fatalErr         error
	userClosed       bool

	runCtx    context.Context
	runCancel context.CancelFunc
	runDone   chan struct{}

	writeMu sync.Mutex // serializes WriteJSON / WriteMessage on the active conn

	subsMu sync.Mutex
	subs   map[string]any // value is string or Symbol
}

// NewClient validates options and prepares the event channels. The underlying
// socket is not opened until Connect is called.
//
// API key resolution: opts.APIKey takes precedence over the
// TWELVEDATA_API_KEY environment variable. If neither is set, NewClient
// returns *AuthError synchronously.
func NewClient(opts Options) (*Client, error) {
	apiKey := opts.APIKey
	if apiKey == "" {
		apiKey = os.Getenv(apiKeyEnvVar)
	}
	if apiKey == "" {
		return nil, &AuthError{Message: apiKeyEnvVar + " is not set and no APIKey was provided"}
	}

	rawURL := opts.URL
	if rawURL == "" {
		rawURL = DefaultURL
	}

	heartbeat := opts.HeartbeatInterval
	if heartbeat == 0 {
		heartbeat = DefaultHeartbeatInterval
	}
	pingI := opts.PingInterval
	if pingI == 0 {
		pingI = DefaultPingInterval
	}
	pingT := opts.PingTimeout
	if pingT == 0 {
		pingT = DefaultPingTimeout
	}

	rec := opts.Reconnect
	if rec.InitialDelay == 0 {
		rec.InitialDelay = DefaultReconnectInitialDelay
	}
	if rec.MaxDelay == 0 {
		rec.MaxDelay = DefaultReconnectMaxDelay
	}
	if rec.MaxAttempts == 0 {
		rec.MaxAttempts = DefaultReconnectMaxAttempts
	}
	if rec.BackoffFactor == 0 {
		rec.BackoffFactor = DefaultReconnectBackoffFactor
	}

	buf := opts.EventBuffer
	if buf <= 0 {
		buf = DefaultEventBuffer
	}

	return &Client{
		apiKey:            apiKey,
		url:               rawURL,
		heartbeatInterval: heartbeat,
		pingInterval:      pingI,
		pingTimeout:       pingT,
		reconnectOpts:     rec,
		headers:           opts.HTTPHeader,

		prices:       make(chan PriceEvent, buf),
		statuses:     make(chan SubscribeStatusEvent, buf),
		heartbeats:   make(chan HeartbeatEvent, buf),
		messages:     make(chan json.RawMessage, buf),
		errs:         make(chan error, buf),
		reconnecting: make(chan ReconnectingEvent, buf),
		reconnected:  make(chan ReconnectedEvent, buf),

		subs: map[string]any{},
	}, nil
}

func (c *Client) Prices() <-chan PriceEvent                      { return c.prices }
func (c *Client) SubscribeStatuses() <-chan SubscribeStatusEvent { return c.statuses }
func (c *Client) Heartbeats() <-chan HeartbeatEvent              { return c.heartbeats }
func (c *Client) Messages() <-chan json.RawMessage               { return c.messages }
func (c *Client) Errors() <-chan error                           { return c.errs }
func (c *Client) Reconnecting() <-chan ReconnectingEvent         { return c.reconnecting }
func (c *Client) Reconnected() <-chan ReconnectedEvent           { return c.reconnected }

// Stats reports per-channel drop counters since the Client was constructed.
func (c *Client) Stats() Stats {
	return Stats{
		DroppedPrices:       c.droppedPrices.Load(),
		DroppedStatuses:     c.droppedStatuses.Load(),
		DroppedHeartbeats:   c.droppedHeartbeats.Load(),
		DroppedMessages:     c.droppedMessages.Load(),
		DroppedErrors:       c.droppedErrors.Load(),
		DroppedReconnecting: c.droppedReconnecting.Load(),
		DroppedReconnected:  c.droppedReconnected.Load(),
	}
}

// IsConnected returns true when there is an active underlying socket.
func (c *Client) IsConnected() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn != nil
}

// Connect dials the WebSocket and starts the supervisor goroutine.
//
// Concurrent calls fold into a single in-flight attempt. After Disconnect,
// Connect can be called again to resume; tracked subscriptions are preserved
// and replayed on the new connection. A failed initial Connect does not
// engage the auto-reconnect loop — only post-success failures do.
func (c *Client) Connect(ctx context.Context) error {
	c.mu.Lock()
	if c.fatalErr != nil {
		err := c.fatalErr
		c.mu.Unlock()
		return err
	}
	if c.conn != nil {
		c.mu.Unlock()
		return nil
	}
	if c.connectingDone != nil {
		done := c.connectingDone
		c.mu.Unlock()
		select {
		case <-done:
		case <-ctx.Done():
			return ctx.Err()
		}
		c.mu.Lock()
		err := c.connectErr
		c.mu.Unlock()
		return err
	}
	done := make(chan struct{})
	c.connectingDone = done
	c.connectErr = nil
	c.userClosed = false
	c.mu.Unlock()

	err := c.dial(ctx)

	c.mu.Lock()
	c.connectErr = err
	c.connectingDone = nil
	if err == nil {
		c.hasEverConnected = true
		c.runCtx, c.runCancel = context.WithCancel(context.Background())
		c.runDone = make(chan struct{})
		go c.supervise()
	}
	c.mu.Unlock()
	close(done)
	return err
}

// dial performs a single connection attempt. On success c.conn is set.
func (c *Client) dial(ctx context.Context) error {
	fullURL, err := buildURL(c.url, c.apiKey)
	if err != nil {
		return &ConnectionError{Message: "invalid URL", Cause: err}
	}
	dialer := *websocket.DefaultDialer
	conn, resp, err := dialer.DialContext(ctx, fullURL, c.headers)
	if err != nil {
		var body []byte
		if resp != nil {
			body = readBody(resp.Body, 4096)
		}
		if errors.Is(err, websocket.ErrBadHandshake) {
			return classifyUpgradeFailure(resp, body)
		}
		return &ConnectionError{Message: "dial failed", Cause: err}
	}

	c.mu.Lock()
	c.conn = conn
	c.mu.Unlock()
	return nil
}

// Disconnect closes the underlying socket cleanly, disables the
// auto-reconnect loop, and waits for goroutines to exit. Tracked
// subscriptions are preserved for the next Connect; use Reset to discard them.
func (c *Client) Disconnect() {
	c.mu.Lock()
	c.userClosed = true
	cancel := c.runCancel
	done := c.runDone
	conn := c.conn
	c.mu.Unlock()

	if cancel != nil {
		cancel()
	}
	if conn != nil {
		_ = conn.Close()
	}
	if done != nil {
		<-done
	}
}

// Subscribe tracks the given inputs locally (for replay on reconnect) and,
// if connected, sends a subscribe action immediately. Accepted input types:
// string (bare or comma-separated tickers), Symbol (extended form), or
// slices thereof — see normalize for details.
func (c *Client) Subscribe(input ...any) error {
	items, err := normalize(input)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

	c.subsMu.Lock()
	for _, it := range items {
		c.subs[canonicalKey(it)] = it
	}
	c.subsMu.Unlock()

	if !c.IsConnected() {
		return nil
	}
	return c.sendSubscribeAction("subscribe", items)
}

// Unsubscribe untracks the given inputs and, if connected, sends an
// unsubscribe action. Untracked items are silently ignored.
func (c *Client) Unsubscribe(input ...any) error {
	items, err := normalize(input)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

	c.subsMu.Lock()
	for _, it := range items {
		delete(c.subs, canonicalKey(it))
	}
	c.subsMu.Unlock()

	if !c.IsConnected() {
		return nil
	}
	return c.sendSubscribeAction("unsubscribe", items)
}

// Reset clears the local subscription set and, if connected, sends a reset
// action. Returns nil when the client is not connected (the local state was
// still cleared).
func (c *Client) Reset() error {
	c.subsMu.Lock()
	c.subs = map[string]any{}
	c.subsMu.Unlock()
	if !c.IsConnected() {
		return nil
	}
	return c.send(map[string]any{"action": "reset"})
}

// Heartbeat sends one application-level heartbeat frame. The auto-heartbeat
// timer (driven by HeartbeatInterval) calls this internally as well.
func (c *Client) Heartbeat() error {
	return c.send(map[string]any{"action": "heartbeat"})
}

// supervise owns the connection lifecycle: read loop, then on disconnect the
// reconnect loop, until userClosed or a fatal error.
func (c *Client) supervise() {
	defer close(c.runDone)

	for {
		c.replaySubscriptions()

		c.runConnection()

		c.mu.Lock()
		conn := c.conn
		c.conn = nil
		userClosed := c.userClosed
		fatal := c.fatalErr
		c.mu.Unlock()
		if conn != nil {
			_ = conn.Close()
		}

		if userClosed || fatal != nil {
			return
		}
		if c.reconnectOpts.Disabled {
			return
		}
		if !c.runReconnect() {
			return
		}
	}
}

// runConnection sets up timers and the read loop on the active c.conn.
// Returns when the connection ends.
func (c *Client) runConnection() {
	c.mu.Lock()
	conn := c.conn
	c.mu.Unlock()
	if conn == nil {
		return
	}

	if c.pingInterval > 0 {
		readDeadline := c.pingInterval + c.pingTimeout
		_ = conn.SetReadDeadline(time.Now().Add(readDeadline))
		conn.SetPongHandler(func(string) error {
			return conn.SetReadDeadline(time.Now().Add(readDeadline))
		})
	}

	var wg sync.WaitGroup
	stop := make(chan struct{})

	if c.heartbeatInterval > 0 {
		wg.Add(1)
		go c.heartbeatLoop(conn, stop, &wg)
	}
	if c.pingInterval > 0 {
		wg.Add(1)
		go c.pingLoop(conn, stop, &wg)
	}

	// Watcher: closes the conn when runCtx is cancelled (Disconnect). Without
	// this, ReadMessage could block indefinitely on a clean shutdown.
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-c.runCtx.Done():
			_ = conn.Close()
		case <-stop:
		}
	}()

	c.readLoop(conn)

	close(stop)
	_ = conn.Close()
	wg.Wait()
}

func (c *Client) readLoop(conn *websocket.Conn) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			c.mu.Lock()
			userClosed := c.userClosed
			c.mu.Unlock()
			if userClosed {
				return
			}
			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				c.emitError(&TimeoutError{Timeout: c.pingTimeout})
			} else if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				c.emitError(&ConnectionError{Message: "read", Cause: err})
			}
			return
		}
		c.dispatch(data)
	}
}

func (c *Client) heartbeatLoop(conn *websocket.Conn, stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(c.heartbeatInterval)
	defer ticker.Stop()
	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			if err := c.writeJSON(conn, map[string]string{"action": "heartbeat"}); err != nil {
				return
			}
		}
	}
}

func (c *Client) pingLoop(conn *websocket.Conn, stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(c.pingInterval)
	defer ticker.Stop()
	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			deadline := time.Now().Add(c.pingTimeout)
			if err := conn.WriteControl(websocket.PingMessage, nil, deadline); err != nil {
				_ = conn.Close()
				return
			}
		}
	}
}

func (c *Client) dispatch(data []byte) {
	var probe struct {
		Event string `json:"event"`
	}
	_ = json.Unmarshal(data, &probe)

	switch probe.Event {
	case "price":
		var ev PriceEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			c.emitMessage(data)
			return
		}
		c.emitPrice(ev)
	case "subscribe-status":
		var ev SubscribeStatusEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			c.emitMessage(data)
			return
		}
		c.emitStatus(ev)
	case "heartbeat":
		var ev HeartbeatEvent
		if err := json.Unmarshal(data, &ev); err != nil {
			c.emitMessage(data)
			return
		}
		c.emitHeartbeat(ev)
	default:
		c.emitMessage(data)
	}
}

// runReconnect runs the exponential-backoff reconnect loop. Returns true on
// successful reconnect (caller continues the supervisor loop), false on
// shutdown / fatal / max-attempts exhaustion.
func (c *Client) runReconnect() bool {
	delay := c.reconnectOpts.InitialDelay
	maxDelay := c.reconnectOpts.MaxDelay
	factor := c.reconnectOpts.BackoffFactor
	maxAttempts := c.reconnectOpts.MaxAttempts

	var lastErr error
	for attempt := 1; ; attempt++ {
		if maxAttempts > 0 && attempt > maxAttempts {
			err := &ReconnectError{Attempts: attempt - 1, Cause: lastErr}
			c.mu.Lock()
			c.fatalErr = err
			c.mu.Unlock()
			c.emitError(err)
			return false
		}

		c.emitReconnecting(ReconnectingEvent{Attempt: attempt, Delay: delay})

		select {
		case <-c.runCtx.Done():
			return false
		case <-time.After(delay):
		}

		c.mu.Lock()
		if c.userClosed {
			c.mu.Unlock()
			return false
		}
		c.mu.Unlock()

		err := c.dial(c.runCtx)
		if err == nil {
			c.emitReconnected(ReconnectedEvent{Attempt: attempt})
			return true
		}
		lastErr = err

		var authErr *AuthError
		if errors.As(err, &authErr) {
			c.mu.Lock()
			c.fatalErr = err
			c.mu.Unlock()
			c.emitError(err)
			return false
		}

		next := float64(delay) * factor
		if next > float64(maxDelay) {
			next = float64(maxDelay)
		}
		if math.IsInf(next, 0) || math.IsNaN(next) {
			next = float64(maxDelay)
		}
		delay = time.Duration(next)
	}
}

// replaySubscriptions sends the tracked subscription set on a freshly opened
// connection.
func (c *Client) replaySubscriptions() {
	c.subsMu.Lock()
	if len(c.subs) == 0 {
		c.subsMu.Unlock()
		return
	}
	items := make([]any, 0, len(c.subs))
	for _, v := range c.subs {
		items = append(items, v)
	}
	c.subsMu.Unlock()

	if err := c.sendSubscribeAction("subscribe", items); err != nil {
		c.emitError(err)
	}
}

// sendSubscribeAction builds the {action, params:{symbols}} payload. When
// every item is a bare string, symbols is sent as a comma-separated string;
// otherwise it is sent as the protocol's array form (which accepts mixed
// strings and objects).
func (c *Client) sendSubscribeAction(action string, items []any) error {
	allString := true
	for _, it := range items {
		if _, ok := it.(string); !ok {
			allString = false
			break
		}
	}
	var symbols any
	if allString {
		ss := make([]string, len(items))
		for i, it := range items {
			ss[i] = it.(string)
		}
		symbols = strings.Join(ss, ",")
	} else {
		symbols = items
	}
	return c.send(map[string]any{
		"action": action,
		"params": map[string]any{"symbols": symbols},
	})
}

func (c *Client) send(payload any) error {
	c.mu.Lock()
	conn := c.conn
	c.mu.Unlock()
	if conn == nil {
		return &ConnectionError{Message: "not connected"}
	}
	return c.writeJSON(conn, payload)
}

func (c *Client) writeJSON(conn *websocket.Conn, payload any) error {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	if err := conn.WriteJSON(payload); err != nil {
		return &ConnectionError{Message: "send failed", Cause: err}
	}
	return nil
}

func buildURL(rawURL, apiKey string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("apikey", apiKey)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// Non-blocking emit helpers. A full channel means the consumer is too slow;
// the event is dropped and the corresponding counter is bumped.

func (c *Client) emitPrice(ev PriceEvent) {
	select {
	case c.prices <- ev:
	default:
		c.droppedPrices.Add(1)
	}
}
func (c *Client) emitStatus(ev SubscribeStatusEvent) {
	select {
	case c.statuses <- ev:
	default:
		c.droppedStatuses.Add(1)
	}
}
func (c *Client) emitHeartbeat(ev HeartbeatEvent) {
	select {
	case c.heartbeats <- ev:
	default:
		c.droppedHeartbeats.Add(1)
	}
}
func (c *Client) emitMessage(data []byte) {
	cp := make(json.RawMessage, len(data))
	copy(cp, data)
	select {
	case c.messages <- cp:
	default:
		c.droppedMessages.Add(1)
	}
}
func (c *Client) emitError(err error) {
	select {
	case c.errs <- err:
	default:
		c.droppedErrors.Add(1)
	}
}
func (c *Client) emitReconnecting(ev ReconnectingEvent) {
	select {
	case c.reconnecting <- ev:
	default:
		c.droppedReconnecting.Add(1)
	}
}
func (c *Client) emitReconnected(ev ReconnectedEvent) {
	select {
	case c.reconnected <- ev:
	default:
		c.droppedReconnected.Add(1)
	}
}
