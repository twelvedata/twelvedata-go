// Package ws provides a WebSocket client for the Twelve Data real-time price
// stream (wss://ws.twelvedata.com/v1/quotes/price).
//
// The wrapper handles authentication, application-level heartbeat, WebSocket
// protocol ping/pong, exponential-backoff auto-reconnect with subscription
// replay, and typed errors. User code only needs to subscribe to symbols and
// read PriceEvent values off the channel returned by Client.Prices.
//
// Construct with NewClient, then Connect, then Subscribe. Event channels are
// buffered; consume promptly or events will be dropped (use Stats to monitor).
//
// See README.md in the repository root for a worked example.
package ws
