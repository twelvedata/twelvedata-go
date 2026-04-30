package ws

import (
	"encoding/json"
	"net/http"
	"time"
)

// Options configures a Client. A zero-valued field falls back to the
// package-level Default* constant; HeartbeatInterval and PingInterval set to
// a negative value disable the corresponding mechanism.
type Options struct {
	APIKey            string        // optional; falls back to $TWELVEDATA_API_KEY
	URL               string        // optional; defaults to DefaultURL
	HeartbeatInterval time.Duration // <= 0 disables app-level heartbeat
	PingInterval      time.Duration // <= 0 disables protocol ping/pong
	PingTimeout       time.Duration
	Reconnect         ReconnectOptions
	EventBuffer       int         // per-channel buffer size; default DefaultEventBuffer
	HTTPHeader        http.Header // additional headers for the upgrade request
}

// ReconnectOptions controls the auto-reconnect loop. The loop only engages
// after the initial Connect has succeeded — a failed initial Connect never
// spawns background retries.
type ReconnectOptions struct {
	Disabled      bool
	InitialDelay  time.Duration
	MaxDelay      time.Duration
	MaxAttempts   int // 0 falls back to Default; -1 means retry forever
	BackoffFactor float64
}

// Symbol is the extended subscription form. Only Symbol is required; the
// other fields disambiguate when multiple instruments share a ticker.
type Symbol struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange,omitempty"`
	MicCode  string `json:"mic_code,omitempty"`
	Type     string `json:"type,omitempty"`
}

// SubscribeStatusItem is a single entry in the success/fails arrays of a
// SubscribeStatusEvent. The server may include additional metadata fields
// (e.g. country) on top of the four canonical Symbol fields.
type SubscribeStatusItem struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange,omitempty"`
	MicCode  string `json:"mic_code,omitempty"`
	Country  string `json:"country,omitempty"`
	Type     string `json:"type,omitempty"`
}

// PriceEvent matches a server "price" event. DayVolume / Bid / Ask are
// pointers because they are only present on certain instrument classes.
type PriceEvent struct {
	Event         string   `json:"event"`
	Symbol        string   `json:"symbol"`
	Exchange      string   `json:"exchange,omitempty"`
	MicCode       string   `json:"mic_code,omitempty"`
	Type          string   `json:"type,omitempty"`
	Currency      string   `json:"currency,omitempty"`
	CurrencyBase  string   `json:"currency_base,omitempty"`
	CurrencyQuote string   `json:"currency_quote,omitempty"`
	Timestamp     int64    `json:"timestamp"`
	Price         float64  `json:"price"`
	DayVolume     *int64   `json:"day_volume,omitempty"`
	Bid           *float64 `json:"bid,omitempty"`
	Ask           *float64 `json:"ask,omitempty"`
}

// SubscribeStatusEvent matches a server "subscribe-status" event. Success
// and Fails are normalized from a possibly-null JSON value to an empty slice.
type SubscribeStatusEvent struct {
	Event   string                `json:"event"`
	Status  string                `json:"status"`
	Success []SubscribeStatusItem `json:"success"`
	Fails   []SubscribeStatusItem `json:"fails"`
}

func (e *SubscribeStatusEvent) UnmarshalJSON(data []byte) error {
	type alias struct {
		Event   string                 `json:"event"`
		Status  string                 `json:"status"`
		Success *[]SubscribeStatusItem `json:"success"`
		Fails   *[]SubscribeStatusItem `json:"fails"`
	}
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	e.Event = a.Event
	e.Status = a.Status
	if a.Success != nil {
		e.Success = *a.Success
	} else {
		e.Success = []SubscribeStatusItem{}
	}
	if a.Fails != nil {
		e.Fails = *a.Fails
	} else {
		e.Fails = []SubscribeStatusItem{}
	}
	return nil
}

// HeartbeatEvent matches a server "heartbeat" event.
type HeartbeatEvent struct {
	Event   string `json:"event"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// ReconnectingEvent fires before each reconnect attempt.
type ReconnectingEvent struct {
	Attempt int
	Delay   time.Duration
}

// ReconnectedEvent fires after a successful reconnect.
type ReconnectedEvent struct {
	Attempt int
}

// Stats reports per-channel drop counters. Useful for diagnosing slow
// consumers — a non-zero counter means the corresponding channel was full
// when the client tried to deliver an event, and the event was discarded.
type Stats struct {
	DroppedPrices       int64
	DroppedStatuses     int64
	DroppedHeartbeats   int64
	DroppedMessages     int64
	DroppedErrors       int64
	DroppedReconnecting int64
	DroppedReconnected  int64
}
