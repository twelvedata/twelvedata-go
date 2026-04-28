package twelvedata

import (
	"fmt"
	"net/http"
	"os"
)

const (
	apiKeyEnvVar  = "TWELVEDATA_API_KEY"
	baseURLEnvVar = "TWELVEDATA_API_BASE_URL"
	apiKeyPrefix  = "apikey"
)

// NewConfig returns a *Configuration pre-loaded with the Twelve Data API key
// and the X-API-Version header.
//
// The API key is taken from the first non-empty value of: the optional first
// argument, or the TWELVEDATA_API_KEY env variable. The base URL is taken from
// the optional second argument, or TWELVEDATA_API_BASE_URL, or the spec default.
//
// Returns an error if no API key is available.
func NewConfig(overrides ...string) (*Configuration, error) {
	apiKey := firstNonEmpty(pick(overrides, 0), os.Getenv(apiKeyEnvVar))
	if apiKey == "" {
		return nil, fmt.Errorf("%s environment variable is not set", apiKeyEnvVar)
	}

	cfg := NewConfiguration()
	cfg.DefaultHeader["Authorization"] = apiKeyPrefix + " " + apiKey
	cfg.DefaultHeader["X-API-Version"] = "last"
	cfg.HTTPClient = &http.Client{Transport: WrapTransport(http.DefaultTransport)}

	if baseURL := firstNonEmpty(pick(overrides, 1), os.Getenv(baseURLEnvVar)); baseURL != "" {
		cfg.Servers = ServerConfigurations{ServerConfiguration{URL: baseURL}}
	}
	return cfg, nil
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func pick(s []string, i int) string {
	if i < len(s) {
		return s[i]
	}
	return ""
}
