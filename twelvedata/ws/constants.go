package ws

import "time"

const (
	DefaultURL                    = "wss://ws.twelvedata.com/v1/quotes/price"
	DefaultHeartbeatInterval      = 10 * time.Second
	DefaultPingInterval           = 30 * time.Second
	DefaultPingTimeout            = 10 * time.Second
	DefaultReconnectInitialDelay  = 1 * time.Second
	DefaultReconnectMaxDelay      = 30 * time.Second
	DefaultReconnectMaxAttempts   = 10
	DefaultReconnectBackoffFactor = 2.0
	DefaultEventBuffer            = 256

	apiKeyEnvVar = "TWELVEDATA_API_KEY"
)
