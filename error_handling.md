## API errors

Non-2xx responses with a recognizable Twelve Data error body are surfaced as typed errors. Use `errors.As` to dispatch on the specific HTTP status, or on the `TwelvedataApiError` interface to catch any of them. The wiring is installed by `NewConfig` — no opt-in required.

| Error type | HTTP Status | Description |
|---|---|---|
| `*BadRequestError` | 400 | Invalid parameters were supplied to the API |
| `*UnauthorizedError` | 401 | Invalid or missing API key |
| `*ForbiddenError` | 403 | Access denied to the requested resource |
| `*NotFoundError` | 404 | The requested resource does not exist |
| `*ParameterTooLongError` | 414 | A request parameter exceeds the maximum allowed length |
| `*TooManyRequestsError` | 429 | API rate limit exceeded |
| `*InternalServerError` | 500 | Unexpected server-side error |

If the API returns an unrecognized HTTP status with a valid error body, a bare `*ApiError` is returned. Responses without a `{code, message, status}` body (HTML pages, malformed JSON, empty bodies) are not wrapped — the generator's default error handling surfaces them. Network and transport errors (DNS, timeouts, connection resets) propagate as the standard library's native errors.

Every typed error exposes the same four fields:
- `StatusCode` — HTTP status code (e.g. 400, 401)
- `Code` — Twelvedata-specific error code
- `Status` — Twelvedata error status string
- `Message` — Human-readable error description

For the complete list of API error codes and their meanings, see the [Twelve Data Errors documentation](https://twelvedata.com/docs/introduction/errors).

```go
package main

import (
    "context"
    "errors"
    "fmt"
    "log"

    twelvedata "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
    cfg, err := twelvedata.NewConfig("YOUR_API_KEY_HERE")
    if err != nil {
        log.Fatal(err)
    }
    client := twelvedata.NewAPIClient(cfg)

    _, _, err = client.MarketDataAPI.GetPrice(context.Background()).Symbol("AAPL").Execute()
    if err != nil {
        var rateLimited *twelvedata.TooManyRequestsError
        if errors.As(err, &rateLimited) {
            // API rate limit exceeded — back off and retry
            return
        }

        var unauthorized *twelvedata.UnauthorizedError
        if errors.As(err, &unauthorized) {
            // Invalid or missing API key
            return
        }

        var apiErr twelvedata.TwelvedataApiError
        if errors.As(err, &apiErr) {
            // Any other Twelve Data API error
            fmt.Printf("API error [%d]: %s\n", apiErr.GetStatusCode(), apiErr.GetMessage())
            return
        }

        // Network / transport / unknown error
        log.Fatal(err)
    }
}
```

Customizing the HTTP client (timeouts, proxy, etc.) keeps error handling installed as long as you wrap your transport with `WrapTransport`:

```go
cfg, _ := twelvedata.NewConfig("YOUR_API_KEY_HERE")
cfg.HTTPClient = &http.Client{
    Timeout:   10 * time.Second,
    Transport: twelvedata.WrapTransport(myCustomTransport), // pass nil to wrap http.DefaultTransport
}
```
