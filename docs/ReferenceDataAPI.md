# \ReferenceDataAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBonds**](ReferenceDataAPI.md#GetBonds) | **Get** /bonds | Fixed income
[**GetCommodities**](ReferenceDataAPI.md#GetCommodities) | **Get** /commodities | Commodities
[**GetCountries**](ReferenceDataAPI.md#GetCountries) | **Get** /countries | Countries
[**GetCrossListings**](ReferenceDataAPI.md#GetCrossListings) | **Get** /cross_listings | Cross listings
[**GetCryptocurrencies**](ReferenceDataAPI.md#GetCryptocurrencies) | **Get** /cryptocurrencies | Cryptocurrency pairs
[**GetCryptocurrencyExchanges**](ReferenceDataAPI.md#GetCryptocurrencyExchanges) | **Get** /cryptocurrency_exchanges | Cryptocurrency exchanges
[**GetEarliestTimestamp**](ReferenceDataAPI.md#GetEarliestTimestamp) | **Get** /earliest_timestamp | Earliest timestamp
[**GetEtf**](ReferenceDataAPI.md#GetEtf) | **Get** /etfs | ETFs
[**GetExchangeSchedule**](ReferenceDataAPI.md#GetExchangeSchedule) | **Get** /exchange_schedule | Exchanges schedule
[**GetExchanges**](ReferenceDataAPI.md#GetExchanges) | **Get** /exchanges | Exchanges
[**GetForexPairs**](ReferenceDataAPI.md#GetForexPairs) | **Get** /forex_pairs | Forex pairs
[**GetFunds**](ReferenceDataAPI.md#GetFunds) | **Get** /funds | Funds
[**GetInstrumentType**](ReferenceDataAPI.md#GetInstrumentType) | **Get** /instrument_type | Instrument type
[**GetIntervals**](ReferenceDataAPI.md#GetIntervals) | **Get** /intervals | Intervals List
[**GetMarketState**](ReferenceDataAPI.md#GetMarketState) | **Get** /market_state | Market state
[**GetStocks**](ReferenceDataAPI.md#GetStocks) | **Get** /stocks | Stocks
[**GetSymbolSearch**](ReferenceDataAPI.md#GetSymbolSearch) | **Get** /symbol_search | Symbol search
[**GetTechnicalIndicators**](ReferenceDataAPI.md#GetTechnicalIndicators) | **Get** /technical_indicators | Technical indicators



## GetBonds

> GetBonds200Response GetBonds(ctx).Symbol(symbol).Exchange(exchange).Country(country).Format(format).Delimiter(delimiter).ShowPlan(showPlan).Page(page).Outputsize(outputsize).Execute()

Fixed income



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	symbol := "US2Y" // string | The ticker symbol of an instrument for which data is requested (optional)
	exchange := "NYSE" // string | Filter by exchange name (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	showPlan := true // bool | Adds info on which plan symbol is available (optional) (default to false)
	page := int64(789) // int64 | Page number of the results to fetch (optional) (default to 1)
	outputsize := int64(789) // int64 | Determines the number of data points returned in the output (optional) (default to 5000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetBonds(context.Background()).Symbol(symbol).Exchange(exchange).Country(country).Format(format).Delimiter(delimiter).ShowPlan(showPlan).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetBonds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBonds`: GetBonds200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetBonds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBondsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested | 
 **exchange** | **string** | Filter by exchange name | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **showPlan** | **bool** | Adds info on which plan symbol is available | [default to false]
 **page** | **int64** | Page number of the results to fetch | [default to 1]
 **outputsize** | **int64** | Determines the number of data points returned in the output | [default to 5000]

### Return type

[**GetBonds200Response**](GetBonds200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommodities

> GetCommodities200Response GetCommodities(ctx).Symbol(symbol).Category(category).Format(format).Delimiter(delimiter).Page(page).Outputsize(outputsize).Execute()

Commodities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	symbol := "XAU/USD" // string | The ticker symbol of an instrument for which data is requested (optional)
	category := "Precious Metal" // string | Filter by category of commodity (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	page := int64(789) // int64 | Page number of the results to fetch (optional) (default to 1)
	outputsize := int64(789) // int64 | Determines the number of data points returned in the output (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetCommodities(context.Background()).Symbol(symbol).Category(category).Format(format).Delimiter(delimiter).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetCommodities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommodities`: GetCommodities200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetCommodities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommoditiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested | 
 **category** | **string** | Filter by category of commodity | 
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **page** | **int64** | Page number of the results to fetch | [default to 1]
 **outputsize** | **int64** | Determines the number of data points returned in the output | 

### Return type

[**GetCommodities200Response**](GetCommodities200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountries

> GetCountries200Response GetCountries(ctx).Execute()

Countries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountries`: GetCountries200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


### Return type

[**GetCountries200Response**](GetCountries200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossListings

> GetCrossListings200Response GetCrossListings(ctx).Symbol(symbol).Exchange(exchange).MicCode(micCode).Country(country).Execute()

Cross listings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	symbol := "AAPL" // string | The ticker symbol of an instrument for which data is requested
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNGS" // string | Market identifier code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country to which stock exchange belongs to (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetCrossListings(context.Background()).Symbol(symbol).Exchange(exchange).MicCode(micCode).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetCrossListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossListings`: GetCrossListings200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetCrossListings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market identifier code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country to which stock exchange belongs to | 

### Return type

[**GetCrossListings200Response**](GetCrossListings200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCryptocurrencies

> GetCryptocurrencies200Response GetCryptocurrencies(ctx).Symbol(symbol).Exchange(exchange).CurrencyBase(currencyBase).CurrencyQuote(currencyQuote).Format(format).Delimiter(delimiter).Page(page).Outputsize(outputsize).Execute()

Cryptocurrency pairs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	symbol := "BTC/USD" // string | The ticker symbol of an instrument for which data is requested (optional)
	exchange := "Binance" // string | Filter by exchange name. E.g. `Binance`, `Coinbase`, etc. (optional)
	currencyBase := "BTC" // string | Filter by currency base (optional)
	currencyQuote := "USD" // string | Filter by currency quote (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	page := int64(789) // int64 | Page number of the results to fetch (optional) (default to 1)
	outputsize := int64(789) // int64 | Determines the number of data points returned in the output (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetCryptocurrencies(context.Background()).Symbol(symbol).Exchange(exchange).CurrencyBase(currencyBase).CurrencyQuote(currencyQuote).Format(format).Delimiter(delimiter).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetCryptocurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCryptocurrencies`: GetCryptocurrencies200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetCryptocurrencies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCryptocurrenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested | 
 **exchange** | **string** | Filter by exchange name. E.g. &#x60;Binance&#x60;, &#x60;Coinbase&#x60;, etc. | 
 **currencyBase** | **string** | Filter by currency base | 
 **currencyQuote** | **string** | Filter by currency quote | 
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **page** | **int64** | Page number of the results to fetch | [default to 1]
 **outputsize** | **int64** | Determines the number of data points returned in the output | 

### Return type

[**GetCryptocurrencies200Response**](GetCryptocurrencies200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCryptocurrencyExchanges

> GetCryptocurrencyExchanges200Response GetCryptocurrencyExchanges(ctx).Format(format).Delimiter(delimiter).Execute()

Cryptocurrency exchanges



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the CSV file (optional) (default to ";")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetCryptocurrencyExchanges(context.Background()).Format(format).Delimiter(delimiter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetCryptocurrencyExchanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCryptocurrencyExchanges`: GetCryptocurrencyExchanges200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetCryptocurrencyExchanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCryptocurrencyExchangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the CSV file | [default to &quot;;&quot;]

### Return type

[**GetCryptocurrencyExchanges200Response**](GetCryptocurrencyExchanges200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEarliestTimestamp

> GetEarliestTimestamp200Response GetEarliestTimestamp(ctx).Interval(interval).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Timezone(timezone).Execute()

Earliest timestamp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	interval := openapiclient.IntervalEnum("1min") // IntervalEnum | Interval between two consecutive points in time series.
	symbol := "AAPL" // string | Symbol ticker of the instrument. (optional)
	figi := "BBG000B9XRY4" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "Nasdaq" // string | Exchange where instrument is traded. (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard. (optional)
	timezone := "timezone_example" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>Exchange</code> for local exchange time</li> <li>2. <code>UTC</code> for datetime at universal UTC standard</li> <li>3. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a>.</li> </ul> <p>Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time.</p> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional) (default to "Exchange")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetEarliestTimestamp(context.Background()).Interval(interval).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetEarliestTimestamp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEarliestTimestamp`: GetEarliestTimestamp200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetEarliestTimestamp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEarliestTimestampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**IntervalEnum**](IntervalEnum.md) | Interval between two consecutive points in time series. | 
 **symbol** | **string** | Symbol ticker of the instrument. | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded. | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard. | 
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time.&lt;/p&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | [default to &quot;Exchange&quot;]

### Return type

[**GetEarliestTimestamp200Response**](GetEarliestTimestamp200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEtf

> GetEtf200Response GetEtf(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Exchange(exchange).MicCode(micCode).Country(country).Format(format).Delimiter(delimiter).ShowPlan(showPlan).IncludeDelisted(includeDelisted).Page(page).Outputsize(outputsize).Execute()

ETFs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	symbol := "SPY" // string | The ticker symbol of an instrument for which data is requested (optional)
	figi := "BBG000BDTF76" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cik := "95953" // string | The CIK of an instrument for which data is requested (optional)
	exchange := "NYSE" // string | Filter by exchange name (optional)
	micCode := "XNYS" // string | Filter by market identifier code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	showPlan := true // bool | Adds info on which plan symbol is available (optional) (default to false)
	includeDelisted := true // bool | Include delisted identifiers (optional) (default to false)
	page := int64(789) // int64 | Page number of the results to fetch (optional) (default to 1)
	outputsize := int64(789) // int64 | Determines the number of data points returned in the output (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetEtf(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Exchange(exchange).MicCode(micCode).Country(country).Format(format).Delimiter(delimiter).ShowPlan(showPlan).IncludeDelisted(includeDelisted).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetEtf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEtf`: GetEtf200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetEtf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEtfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cik** | **string** | The CIK of an instrument for which data is requested | 
 **exchange** | **string** | Filter by exchange name | 
 **micCode** | **string** | Filter by market identifier code (MIC) under ISO 10383 standard | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **showPlan** | **bool** | Adds info on which plan symbol is available | [default to false]
 **includeDelisted** | **bool** | Include delisted identifiers | [default to false]
 **page** | **int64** | Page number of the results to fetch | [default to 1]
 **outputsize** | **int64** | Determines the number of data points returned in the output | 

### Return type

[**GetEtf200Response**](GetEtf200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeSchedule

> GetExchangeSchedule200Response GetExchangeSchedule(ctx).MicName(micName).MicCode(micCode).Country(country).Date(date).Execute()

Exchanges schedule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	micName := "NASDAQ" // string | Filter by exchange name (optional)
	micCode := "XNGS" // string | Filter by market identifier code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	date := "2021-10-27" // string | <p> If a date is provided, the API returns the schedule for the specified date; otherwise, it returns the default (common) schedule. </p> The date can be specified in one of the following formats: <ul> <li>An exact date (e.g., <code>2021-10-27</code>)</li> <li>A human-readable keyword: <code>today</code> or <code>yesterday</code></li> <li>A full datetime string in UTC (e.g., <code>2025-04-11T20:00:00</code>) to retrieve the schedule corresponding to the day in the specified time.</li> </ul> When using a datetime value, the resulting schedule will correspond to the local calendar day at the specified time. For example, <code>2025-04-11T20:00:00 UTC</code> corresponds to: <ul> <li><code>2025-04-11</code> in the <code>America/New_York</code> timezone</li> <li><code>2025-04-12</code> in the <code>Australia/Sydney</code> timezone</li> </ul> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetExchangeSchedule(context.Background()).MicName(micName).MicCode(micCode).Country(country).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetExchangeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeSchedule`: GetExchangeSchedule200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetExchangeSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **micName** | **string** | Filter by exchange name | 
 **micCode** | **string** | Filter by market identifier code (MIC) under ISO 10383 standard | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **date** | **string** | &lt;p&gt; If a date is provided, the API returns the schedule for the specified date; otherwise, it returns the default (common) schedule. &lt;/p&gt; The date can be specified in one of the following formats: &lt;ul&gt; &lt;li&gt;An exact date (e.g., &lt;code&gt;2021-10-27&lt;/code&gt;)&lt;/li&gt; &lt;li&gt;A human-readable keyword: &lt;code&gt;today&lt;/code&gt; or &lt;code&gt;yesterday&lt;/code&gt;&lt;/li&gt; &lt;li&gt;A full datetime string in UTC (e.g., &lt;code&gt;2025-04-11T20:00:00&lt;/code&gt;) to retrieve the schedule corresponding to the day in the specified time.&lt;/li&gt; &lt;/ul&gt; When using a datetime value, the resulting schedule will correspond to the local calendar day at the specified time. For example, &lt;code&gt;2025-04-11T20:00:00 UTC&lt;/code&gt; corresponds to: &lt;ul&gt; &lt;li&gt;&lt;code&gt;2025-04-11&lt;/code&gt; in the &lt;code&gt;America/New_York&lt;/code&gt; timezone&lt;/li&gt; &lt;li&gt;&lt;code&gt;2025-04-12&lt;/code&gt; in the &lt;code&gt;Australia/Sydney&lt;/code&gt; timezone&lt;/li&gt; &lt;/ul&gt; | 

### Return type

[**GetExchangeSchedule200Response**](GetExchangeSchedule200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchanges

> GetExchanges200Response GetExchanges(ctx).Type_(type_).Name(name).Code(code).Country(country).Format(format).Delimiter(delimiter).ShowPlan(showPlan).Execute()

Exchanges



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	type_ := openapiclient.TypeEnum("American Depositary Receipt") // TypeEnum | The asset class to which the instrument belongs (optional)
	name := "NASDAQ" // string | Filter by exchange name (optional)
	code := "XBUE" // string | Filter by market identifier code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	showPlan := true // bool | Adds info on which plan symbol is available (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetExchanges(context.Background()).Type_(type_).Name(name).Code(code).Country(country).Format(format).Delimiter(delimiter).ShowPlan(showPlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetExchanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchanges`: GetExchanges200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetExchanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**TypeEnum**](TypeEnum.md) | The asset class to which the instrument belongs | 
 **name** | **string** | Filter by exchange name | 
 **code** | **string** | Filter by market identifier code (MIC) under ISO 10383 standard | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **showPlan** | **bool** | Adds info on which plan symbol is available | [default to false]

### Return type

[**GetExchanges200Response**](GetExchanges200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForexPairs

> GetForexPairs200Response GetForexPairs(ctx).Symbol(symbol).CurrencyBase(currencyBase).CurrencyQuote(currencyQuote).Format(format).Delimiter(delimiter).Page(page).Outputsize(outputsize).Execute()

Forex pairs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	symbol := "EUR/USD" // string | The ticker symbol of an instrument for which data is requested (optional)
	currencyBase := "EUR" // string | Filter by currency base (optional)
	currencyQuote := "USD" // string | Filter by currency quote (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	page := int64(789) // int64 | Page number of the results to fetch (optional) (default to 1)
	outputsize := int64(789) // int64 | Determines the number of data points returned in the output (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetForexPairs(context.Background()).Symbol(symbol).CurrencyBase(currencyBase).CurrencyQuote(currencyQuote).Format(format).Delimiter(delimiter).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetForexPairs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForexPairs`: GetForexPairs200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetForexPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetForexPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested | 
 **currencyBase** | **string** | Filter by currency base | 
 **currencyQuote** | **string** | Filter by currency quote | 
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **page** | **int64** | Page number of the results to fetch | [default to 1]
 **outputsize** | **int64** | Determines the number of data points returned in the output | 

### Return type

[**GetForexPairs200Response**](GetForexPairs200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunds

> GetFunds200Response GetFunds(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Exchange(exchange).Country(country).Format(format).Delimiter(delimiter).ShowPlan(showPlan).Page(page).Outputsize(outputsize).Execute()

Funds



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	symbol := "FXAIX" // string | The ticker symbol of an instrument for which data is requested (optional)
	figi := "BBG000BHTMY7" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cik := "95953" // string | The CIK of an instrument for which data is requested (optional)
	exchange := "Nasdaq" // string | Filter by exchange name (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	showPlan := true // bool | Adds info on which plan symbol is available (optional) (default to false)
	page := int64(789) // int64 | Page number of the results to fetch (optional) (default to 1)
	outputsize := int64(789) // int64 | Determines the number of data points returned in the output (optional) (default to 5000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetFunds(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Exchange(exchange).Country(country).Format(format).Delimiter(delimiter).ShowPlan(showPlan).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetFunds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunds`: GetFunds200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetFunds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cik** | **string** | The CIK of an instrument for which data is requested | 
 **exchange** | **string** | Filter by exchange name | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **showPlan** | **bool** | Adds info on which plan symbol is available | [default to false]
 **page** | **int64** | Page number of the results to fetch | [default to 1]
 **outputsize** | **int64** | Determines the number of data points returned in the output | [default to 5000]

### Return type

[**GetFunds200Response**](GetFunds200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstrumentType

> GetInstrumentType200Response GetInstrumentType(ctx).Execute()

Instrument type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetInstrumentType(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetInstrumentType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstrumentType`: GetInstrumentType200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetInstrumentType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstrumentTypeRequest struct via the builder pattern


### Return type

[**GetInstrumentType200Response**](GetInstrumentType200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntervals

> GetIntervals200Response GetIntervals(ctx).Execute()

Intervals List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetIntervals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetIntervals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntervals`: GetIntervals200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetIntervals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntervalsRequest struct via the builder pattern


### Return type

[**GetIntervals200Response**](GetIntervals200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketState

> []MarketStateResponseItem GetMarketState(ctx).Exchange(exchange).Code(code).Country(country).Execute()

Market state



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	exchange := "NYSE" // string | The exchange name where the instrument is traded. (optional)
	code := "XNYS" // string | The Market Identifier Code (MIC) of the exchange where the instrument is traded. (optional)
	country := "United States" // string | The country where the exchange is located. Takes country name or alpha code. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetMarketState(context.Background()).Exchange(exchange).Code(code).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetMarketState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketState`: []MarketStateResponseItem
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetMarketState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchange** | **string** | The exchange name where the instrument is traded. | 
 **code** | **string** | The Market Identifier Code (MIC) of the exchange where the instrument is traded. | 
 **country** | **string** | The country where the exchange is located. Takes country name or alpha code. | 

### Return type

[**[]MarketStateResponseItem**](MarketStateResponseItem.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStocks

> GetStocks200Response GetStocks(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Format(format).Delimiter(delimiter).ShowPlan(showPlan).IncludeDelisted(includeDelisted).Page(page).Outputsize(outputsize).Execute()

Stocks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	symbol := "AAPL" // string | The ticker symbol of an instrument for which data is requested (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cik := "95953" // string | The CIK of an instrument for which data is requested (optional)
	exchange := "NASDAQ" // string | Filter by exchange name (optional)
	micCode := "XNGS" // string | Filter by market identifier code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	type_ := openapiclient.TypeEnum("American Depositary Receipt") // TypeEnum | The asset class to which the instrument belongs (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	showPlan := true // bool | Adds info on which plan symbol is available (optional) (default to false)
	includeDelisted := true // bool | Include delisted identifiers (optional) (default to false)
	page := int64(789) // int64 | Page number of the results to fetch (optional) (default to 1)
	outputsize := int64(789) // int64 | Determines the number of data points returned in the output (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetStocks(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Format(format).Delimiter(delimiter).ShowPlan(showPlan).IncludeDelisted(includeDelisted).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetStocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStocks`: GetStocks200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetStocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cik** | **string** | The CIK of an instrument for which data is requested | 
 **exchange** | **string** | Filter by exchange name | 
 **micCode** | **string** | Filter by market identifier code (MIC) under ISO 10383 standard | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **type_** | [**TypeEnum**](TypeEnum.md) | The asset class to which the instrument belongs | 
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **showPlan** | **bool** | Adds info on which plan symbol is available | [default to false]
 **includeDelisted** | **bool** | Include delisted identifiers | [default to false]
 **page** | **int64** | Page number of the results to fetch | [default to 1]
 **outputsize** | **int64** | Determines the number of data points returned in the output | 

### Return type

[**GetStocks200Response**](GetStocks200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSymbolSearch

> GetSymbolSearch200Response GetSymbolSearch(ctx).Symbol(symbol).Outputsize(outputsize).ShowPlan(showPlan).Execute()

Symbol search



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	symbol := "AAPL" // string | Symbol to search. Supports: <ul> <li>Ticker symbol of instrument.</li> <li>International securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section</li> <li>The FIGI (Financial Instrument Global Identifier) parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above.</li> <li>Composite FIGI parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above.</li> <li>Share Class FIGI parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above.</li> </ul>
	outputsize := int64(789) // int64 | Number of matches in response. Max <code>120</code> (optional) (default to 30)
	showPlan := true // bool | Adds info on which plan symbol is available. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetSymbolSearch(context.Background()).Symbol(symbol).Outputsize(outputsize).ShowPlan(showPlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetSymbolSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSymbolSearch`: GetSymbolSearch200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetSymbolSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSymbolSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol to search. Supports: &lt;ul&gt; &lt;li&gt;Ticker symbol of instrument.&lt;/li&gt; &lt;li&gt;International securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section&lt;/li&gt; &lt;li&gt;The FIGI (Financial Instrument Global Identifier) parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.&lt;/li&gt; &lt;li&gt;Composite FIGI parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.&lt;/li&gt; &lt;li&gt;Share Class FIGI parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above.&lt;/li&gt; &lt;/ul&gt; | 
 **outputsize** | **int64** | Number of matches in response. Max &lt;code&gt;120&lt;/code&gt; | [default to 30]
 **showPlan** | **bool** | Adds info on which plan symbol is available. | [default to false]

### Return type

[**GetSymbolSearch200Response**](GetSymbolSearch200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTechnicalIndicators

> GetTechnicalIndicators200Response GetTechnicalIndicators(ctx).Execute()

Technical indicators



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataAPI.GetTechnicalIndicators(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.GetTechnicalIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTechnicalIndicators`: GetTechnicalIndicators200Response
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.GetTechnicalIndicators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTechnicalIndicatorsRequest struct via the builder pattern


### Return type

[**GetTechnicalIndicators200Response**](GetTechnicalIndicators200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

