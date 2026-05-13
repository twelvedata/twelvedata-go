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
	sourceParam   = "source"
	// DefaultSourceValue is the `source` query value used by NewConfig and
	// WrapTransport. Tools built on top of this SDK (e.g. the Twelve Data
	// CLI) should call NewConfigWithSource/WrapTransportWithSource to
	// attribute their traffic instead.
	DefaultSourceValue = "client-go"
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
	return NewConfigWithSource(DefaultSourceValue, overrides...)
}

// NewConfigWithSource is like NewConfig but tags outbound requests with the
// caller-supplied `source` query value (e.g. "cli" for the Twelve Data CLI),
// so the API can attribute traffic to the wrapping tool rather than the raw
// Go SDK.
func NewConfigWithSource(source string, overrides ...string) (*Configuration, error) {
	apiKey := firstNonEmpty(pick(overrides, 0), os.Getenv(apiKeyEnvVar))
	if apiKey == "" {
		return nil, fmt.Errorf("%s environment variable is not set", apiKeyEnvVar)
	}

	cfg := NewConfiguration()
	cfg.DefaultHeader["Authorization"] = apiKeyPrefix + " " + apiKey
	cfg.DefaultHeader["X-API-Version"] = "last"
	cfg.HTTPClient = &http.Client{Transport: WrapTransportWithSource(http.DefaultTransport, source)}

	if baseURL := firstNonEmpty(pick(overrides, 1), os.Getenv(baseURLEnvVar)); baseURL != "" {
		cfg.Servers = ServerConfigurations{ServerConfiguration{URL: baseURL}}
	}
	return cfg, nil
}

// sourceParamRoundTripper attaches the configured `source` query parameter to
// every outbound request, so the API can attribute traffic to the calling
// client. User-supplied values for the same parameter take precedence.
type sourceParamRoundTripper struct {
	base   http.RoundTripper
	source string
}

func (t *sourceParamRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL == nil {
		return t.base.RoundTrip(req)
	}
	q := req.URL.Query()
	if q.Get(sourceParam) != "" {
		return t.base.RoundTrip(req)
	}
	newReq := req.Clone(req.Context())
	q.Set(sourceParam, t.source)
	newReq.URL.RawQuery = q.Encode()
	return t.base.RoundTrip(newReq)
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
