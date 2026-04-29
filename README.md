# Twelve Data API client for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/twelvedata/twelvedata-go/twelvedata.svg)](https://pkg.go.dev/github.com/twelvedata/twelvedata-go/twelvedata)
[![Go Report Card](https://goreportcard.com/badge/github.com/twelvedata/twelvedata-go)](https://goreportcard.com/report/github.com/twelvedata/twelvedata-go)
[![License](https://img.shields.io/github/license/twelvedata/twelvedata-go)](https://github.com/twelvedata/twelvedata-go/blob/master/LICENSE.txt)
[![Issues](https://img.shields.io/github/issues/twelvedata/twelvedata-go)](https://github.com/twelvedata/twelvedata-go/issues)

Twelve Data official library. This package supports all main features of the service:
- Real-time and historical market data: time series, quotes, end-of-day prices, exchange rates, and market movers.
- Fundamentals: financial statements, earnings, dividends, splits, company profiles, and key statistics.
- 100+ technical indicators: SMA, EMA, RSI, MACD, Bollinger Bands, Ichimoku, and many more.
- Analysis & estimates: analyst ratings, price targets, EPS trends, revenue and growth estimates.
- ETFs and mutual funds: summaries, composition, performance, risk, and ratings.

🔑 **API key** is required. If you don't have it yet, get it by [signing up here](https://twelvedata.com/pricing).

## Requirements
Go 1.21 or later.

## Installation

```bash
go get github.com/twelvedata/twelvedata-go/twelvedata
```

🔗 View the package on [pkg.go.dev](https://pkg.go.dev/github.com/twelvedata/twelvedata-go/twelvedata).

## Quick start

### 1. Create `main.go`

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
    cfg, err := twelvedata.NewConfig("YOUR_API_KEY_HERE") // defaults to os.Getenv("TWELVEDATA_API_KEY")
    if err != nil {
        log.Fatal(err)
    }
    client := twelvedata.NewAPIClient(cfg)

    resp, _, err := client.MarketDataAPI.GetTimeSeries(context.Background()).
        Symbol("AAPL").Interval("1day").Outputsize(5).Execute()
    if err != nil {
        var apiErr twelvedata.TwelvedataApiError
        if errors.As(err, &apiErr) {
            log.Fatalf("API error [%d]: %s", apiErr.GetStatusCode(), apiErr.GetMessage())
        }
        log.Fatalf("Unexpected error: %v", err)
    }
    fmt.Println(resp)
}
```

### 2. Run

```bash
go run main.go
```

👀 Check the full example and other examples [here](https://github.com/twelvedata/twelvedata-go/tree/master/examples).

## Error Handling

See [error_handling.md](error_handling.md) for error types, fields, and usage examples.

For the complete list of API error codes and their meanings, see the [Twelve Data Errors documentation](https://twelvedata.com/docs/introduction/errors).

## API Reference

See [api_reference.md](api_reference.md) for the complete list of API endpoints and models.

## Integration tests

This repo ships a live-API integration suite at [`twelvedata/endpoints_test.go`](twelvedata/endpoints_test.go). It hits every supported endpoint once against the real Twelve Data API and is used to catch regressions before publishing a new client version. **Running the suite spends API quota.**

### Setup

```bash
export TWELVEDATA_API_KEY=<your-key>
```

### Run

```bash
go test -tags=integration -v -timeout=10m ./twelvedata/...
```

The build tag (`-tags=integration`) keeps the suite out of `go test ./...`. On success the output ends with `PASS`; a non-zero exit means at least one endpoint failed.

### Editor setup

The `//go:build integration` tag also hides `endpoints_test.go` from default Go tooling, so editors using `gopls` (VS Code, GoLand, Neovim, etc.) report "no packages found" and disable autocomplete in this file out of the box. To get full IntelliSense, add the `integration` build tag to your editor's Go config — for example, in VS Code's `settings.json`:

```json
{
    "go.buildTags": "integration",
    "gopls": {
        "build.buildFlags": ["-tags=integration"]
    }
}
```

## Documentation
Delve deeper with our [official documentation](https://twelvedata.com/docs).

## Examples
Explore practical scenarios in the [examples](https://github.com/twelvedata/twelvedata-go/tree/master/examples) directory.

## Support
- 🌐 Visit our [contact page](https://twelvedata.com/contact).
- 🛠 Or our [support center](https://support.twelvedata.com/).

## Stay updated
- Follow us on [𝕏](https://twitter.com/TwelveData).
- Follow us on [Telegram](https://t.me/twelvedata).

## Contributing
1. Fork and clone: `$ git clone https://github.com/twelvedata/twelvedata-go .`.
2. Branch out: `$ git checkout -b my-feature`.
3. Commit changes and test.
4. Push your branch and open a pull request with a comprehensive description.

For more guidance on contributing, see the [GitHub Docs](https://docs.github.com/en/get-started/quickstart/contributing-to-projects) on GitHub.

## License

This project is licensed under the MIT. See the `LICENSE.txt` file in this repository for more details.