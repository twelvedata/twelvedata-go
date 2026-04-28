# \MarketDataAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEod**](MarketDataAPI.md#GetEod) | **Get** /eod | End of day price
[**GetMarketMovers**](MarketDataAPI.md#GetMarketMovers) | **Get** /market_movers/{market} | Market movers
[**GetPrice**](MarketDataAPI.md#GetPrice) | **Get** /price | Latest price
[**GetQuote**](MarketDataAPI.md#GetQuote) | **Get** /quote | Quote
[**GetTimeSeries**](MarketDataAPI.md#GetTimeSeries) | **Get** /time_series | Time series
[**GetTimeSeriesCross**](MarketDataAPI.md#GetTimeSeriesCross) | **Get** /time_series/cross | Time series cross



## GetEod

> GetEod200Response GetEod(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Date(date).Prepost(prepost).Dp(dp).Execute()

End of day price



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
	symbol := "AAPL" // string | Symbol ticker of the instrument (optional)
	figi := "BBG000BHTMY7" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	type_ := openapiclient.TypeEnum("American Depositary Receipt") // TypeEnum | The asset class to which the instrument belongs (optional)
	date := "2006-01-02" // string | If not null, then return data from a specific date (optional)
	prepost := true // bool | Parameter is optional. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume (optional) (default to false)
	dp := int64(789) // int64 | Specifies the number of decimal places for floating values Should be in range [0,11] inclusive (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetEod(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Date(date).Prepost(prepost).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetEod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEod`: GetEod200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetEod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of the instrument | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **type_** | [**TypeEnum**](TypeEnum.md) | The asset class to which the instrument belongs | 
 **date** | **string** | If not null, then return data from a specific date | 
 **prepost** | **bool** | Parameter is optional. Only for the &#x60;Pro&#x60; plan (individual) and &#x60;Venture&#x60; plan (business) and above. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume | [default to false]
 **dp** | **int64** | Specifies the number of decimal places for floating values Should be in range [0,11] inclusive | [default to 5]

### Return type

[**GetEod200Response**](GetEod200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketMovers

> MarketMoversResponseBody GetMarketMovers(ctx, market).Direction(direction).Outputsize(outputsize).Country(country).PriceGreaterThan(priceGreaterThan).Dp(dp).Execute()

Market movers



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
	market := openapiclient.MarketEnum("stocks") // MarketEnum | Maket type
	direction := openapiclient.DirectionEnum("gainers") // DirectionEnum | Specifies direction of the snapshot gainers or losers (optional) (default to "gainers")
	outputsize := int64(789) // int64 | Specifies the size of the snapshot. Can be in a range from `1` to `50` (optional) (default to 30)
	country := "country_example" // string | Country of the snapshot, applicable to non-currencies only. Takes country name or alpha code (optional) (default to "USA")
	priceGreaterThan := "175.5" // string | Takes values with price grater than specified value (optional)
	dp := "dp_example" // string | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive (optional) (default to "5")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketMovers(context.Background(), market).Direction(direction).Outputsize(outputsize).Country(country).PriceGreaterThan(priceGreaterThan).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketMovers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketMovers`: MarketMoversResponseBody
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketMovers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**market** | [**MarketEnum**](.md) | Maket type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketMoversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | [**DirectionEnum**](DirectionEnum.md) | Specifies direction of the snapshot gainers or losers | [default to &quot;gainers&quot;]
 **outputsize** | **int64** | Specifies the size of the snapshot. Can be in a range from &#x60;1&#x60; to &#x60;50&#x60; | [default to 30]
 **country** | **string** | Country of the snapshot, applicable to non-currencies only. Takes country name or alpha code | [default to &quot;USA&quot;]
 **priceGreaterThan** | **string** | Takes values with price grater than specified value | 
 **dp** | **string** | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive | [default to &quot;5&quot;]

### Return type

[**MarketMoversResponseBody**](MarketMoversResponseBody.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrice

> GetPrice200Response GetPrice(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Format(format).Delimiter(delimiter).Prepost(prepost).Dp(dp).Execute()

Latest price



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
	symbol := "AAPL" // string | Symbol ticker of the instrument (optional)
	figi := "BBG000BHTMY7" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	type_ := openapiclient.TypeEnum("American Depositary Receipt") // TypeEnum | The asset class to which the instrument belongs (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | Value can be JSON or CSV (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the CSV file (optional) (default to ";")
	prepost := true // bool | Parameter is optional. Only for Pro or Venture, and above plans. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume. (optional) (default to false)
	dp := int64(789) // int64 | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetPrice(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Format(format).Delimiter(delimiter).Prepost(prepost).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrice`: GetPrice200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of the instrument | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **type_** | [**TypeEnum**](TypeEnum.md) | The asset class to which the instrument belongs | 
 **format** | [**FormatEnum**](FormatEnum.md) | Value can be JSON or CSV | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the CSV file | [default to &quot;;&quot;]
 **prepost** | **bool** | Parameter is optional. Only for Pro or Venture, and above plans. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume. | [default to false]
 **dp** | **int64** | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive | [default to 5]

### Return type

[**GetPrice200Response**](GetPrice200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuote

> GetQuote200Response GetQuote(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Interval(interval).Exchange(exchange).MicCode(micCode).Country(country).VolumeTimePeriod(volumeTimePeriod).Type_(type_).Format(format).Delimiter(delimiter).Prepost(prepost).Eod(eod).RollingPeriod(rollingPeriod).Dp(dp).Timezone(timezone).Execute()

Quote



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
	symbol := "AAPL" // string | Symbol ticker of the instrument (optional)
	figi := "BBG000BHTMY7" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	interval := openapiclient.IntervalEnum("1min") // IntervalEnum | Interval of the quote (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	volumeTimePeriod := int64(789) // int64 | Number of periods for Average Volume (optional) (default to 9)
	type_ := openapiclient.TypeEnum("American Depositary Receipt") // TypeEnum | The asset class to which the instrument belongs (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | Value can be JSON or CSV Default JSON (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the CSV file (optional) (default to ";")
	prepost := true // bool | Parameter is optional. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume. (optional) (default to false)
	eod := true // bool | If true, then return data for closed day (optional) (default to false)
	rollingPeriod := int64(789) // int64 | Number of hours for calculate rolling change at period. By default set to 24, it can be in range [1, 168]. (optional) (default to 24)
	dp := int64(789) // int64 | Specifies the number of decimal places for floating values Should be in range [0,11] inclusive (optional) (default to 5)
	timezone := "timezone_example" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>Exchange</code> for local exchange time</li> <li>2. <code>UTC</code> for datetime at universal UTC standard</li> <li>3. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a>.</li> </ul> <p>Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time.</p> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional) (default to "Exchange")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetQuote(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Interval(interval).Exchange(exchange).MicCode(micCode).Country(country).VolumeTimePeriod(volumeTimePeriod).Type_(type_).Format(format).Delimiter(delimiter).Prepost(prepost).Eod(eod).RollingPeriod(rollingPeriod).Dp(dp).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuote`: GetQuote200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of the instrument | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **interval** | [**IntervalEnum**](IntervalEnum.md) | Interval of the quote | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **volumeTimePeriod** | **int64** | Number of periods for Average Volume | [default to 9]
 **type_** | [**TypeEnum**](TypeEnum.md) | The asset class to which the instrument belongs | 
 **format** | [**FormatEnum**](FormatEnum.md) | Value can be JSON or CSV Default JSON | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the CSV file | [default to &quot;;&quot;]
 **prepost** | **bool** | Parameter is optional. Only for the &#x60;Pro&#x60; plan (individual) and &#x60;Venture&#x60; plan (business) and above. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume. | [default to false]
 **eod** | **bool** | If true, then return data for closed day | [default to false]
 **rollingPeriod** | **int64** | Number of hours for calculate rolling change at period. By default set to 24, it can be in range [1, 168]. | [default to 24]
 **dp** | **int64** | Specifies the number of decimal places for floating values Should be in range [0,11] inclusive | [default to 5]
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time.&lt;/p&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | [default to &quot;Exchange&quot;]

### Return type

[**GetQuote200Response**](GetQuote200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeSeries

> GetTimeSeries200Response GetTimeSeries(ctx).Interval(interval).Symbol(symbol).Isin(isin).Figi(figi).Cusip(cusip).Outputsize(outputsize).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Timezone(timezone).StartDate(startDate).EndDate(endDate).Date(date).Order(order).Prepost(prepost).Format(format).Delimiter(delimiter).Dp(dp).PreviousClose(previousClose).Adjust(adjust).Execute()

Time series



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
	interval := openapiclient.IntervalEnum("1min") // IntervalEnum | Interval between two consecutive points in time series
	symbol := "AAPL" // string | Symbol ticker of the instrument. E.g. `AAPL`, `EUR/USD`, `ETH/BTC`, ... (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	figi := "BBG000B9Y5X2" // string | The FIGI of an instrument for which data is requested. This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	outputsize := int64(789) // int64 | Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum (optional) (default to 30)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | The country where the instrument is traded, e.g., `United States` or `US` (optional)
	type_ := openapiclient.TypeEnum("American Depositary Receipt") // TypeEnum | The asset class to which the instrument belongs (optional)
	timezone := "timezone_example" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>Exchange</code> for local exchange time</li> <li>2. <code>UTC</code> for datetime at universal UTC standard</li> <li>3. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a></li> </ul> <p>Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time.</p> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional) (default to "Exchange")
	startDate := "2024-08-22T15:04:05" // string | Can be used separately and together with `end_date`. Format `2006-01-02` or `2006-01-02T15:04:05`  Default location: <ul> <li>Forex and Cryptocurrencies - <code>UTC</code></li> <li>Stocks - where exchange is located (e.g. for AAPL it will be <code>America/New_York</code>)</li> </ul> Both parameters take into account if <code>timezone</code> parameter is provided.<br/> If <code>timezone</code> is given then, <code>start_date</code> and <code>end_date</code> will be used in the specified location  Examples: <ul> <li>1. <code>&symbol=AAPL&start_date=2019-08-09T15:50:00&…</code><br/> Returns all records starting from 2019-08-09T15:50:00 New York time up to current date</li> <li>2. <code>&symbol=EUR/USD&timezone=Asia/Singapore&start_date=2019-08-09T15:50:00&…</code><br/> Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date</li> <li>3. <code>&symbol=ETH/BTC&timezone=Europe/Zurich&start_date=2019-08-09T15:50:00&end_date=2019-08-09T15:55:00&...</code><br/> Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00</li> </ul> (optional)
	endDate := "2024-08-22T16:04:05" // string | The ending date and time for data selection, see `start_date` description for details. (optional)
	date := "2021-10-27" // string | Specifies the exact date to get the data for. Could be the exact date, e.g. `2021-10-27`, or in human language `today` or `yesterday` (optional)
	order := openapiclient.OrderEnum("asc") // OrderEnum | Sorting order of the output (optional) (default to "desc")
	prepost := true // bool | Returns quotes that include pre-market and post-market data. Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume (optional) (default to false)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	dp := int64(789) // int64 | Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided (optional) (default to -1)
	previousClose := true // bool | A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object (optional) (default to false)
	adjust := openapiclient.AdjustEnum("all") // AdjustEnum | Adjusting mode for prices (optional) (default to "splits")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetTimeSeries(context.Background()).Interval(interval).Symbol(symbol).Isin(isin).Figi(figi).Cusip(cusip).Outputsize(outputsize).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Timezone(timezone).StartDate(startDate).EndDate(endDate).Date(date).Order(order).Prepost(prepost).Format(format).Delimiter(delimiter).Dp(dp).PreviousClose(previousClose).Adjust(adjust).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetTimeSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeSeries`: GetTimeSeries200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetTimeSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interval** | [**IntervalEnum**](IntervalEnum.md) | Interval between two consecutive points in time series | 
 **symbol** | **string** | Symbol ticker of the instrument. E.g. &#x60;AAPL&#x60;, &#x60;EUR/USD&#x60;, &#x60;ETH/BTC&#x60;, ... | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **figi** | **string** | The FIGI of an instrument for which data is requested. This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **outputsize** | **int64** | Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;5000&#x60;. Default &#x60;30&#x60; when no date parameters are set, otherwise set to maximum | [default to 30]
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | The country where the instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **type_** | [**TypeEnum**](TypeEnum.md) | The asset class to which the instrument belongs | 
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Interval Limitation: The timezone parameter is only applicable for intraday intervals (less than 1 day). For intervals of 1day, 1week, or 1month, the timezone parameter is ignored, and data is strictly returned in the Exchange local time.&lt;/p&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | [default to &quot;Exchange&quot;]
 **startDate** | **string** | Can be used separately and together with &#x60;end_date&#x60;. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;  Default location: &lt;ul&gt; &lt;li&gt;Forex and Cryptocurrencies - &lt;code&gt;UTC&lt;/code&gt;&lt;/li&gt; &lt;li&gt;Stocks - where exchange is located (e.g. for AAPL it will be &lt;code&gt;America/New_York&lt;/code&gt;)&lt;/li&gt; &lt;/ul&gt; Both parameters take into account if &lt;code&gt;timezone&lt;/code&gt; parameter is provided.&lt;br/&gt; If &lt;code&gt;timezone&lt;/code&gt; is given then, &lt;code&gt;start_date&lt;/code&gt; and &lt;code&gt;end_date&lt;/code&gt; will be used in the specified location  Examples: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;&amp;symbol&#x3D;AAPL&amp;start_date&#x3D;2019-08-09T15:50:00&amp;…&lt;/code&gt;&lt;br/&gt; Returns all records starting from 2019-08-09T15:50:00 New York time up to current date&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;&amp;symbol&#x3D;EUR/USD&amp;timezone&#x3D;Asia/Singapore&amp;start_date&#x3D;2019-08-09T15:50:00&amp;…&lt;/code&gt;&lt;br/&gt; Returns all records starting from 2019-08-09T15:50:00 Singapore time up to current date&lt;/li&gt; &lt;li&gt;3. &lt;code&gt;&amp;symbol&#x3D;ETH/BTC&amp;timezone&#x3D;Europe/Zurich&amp;start_date&#x3D;2019-08-09T15:50:00&amp;end_date&#x3D;2019-08-09T15:55:00&amp;...&lt;/code&gt;&lt;br/&gt; Returns all records starting from 2019-08-09T15:50:00 Zurich time up to 2019-08-09T15:55:00&lt;/li&gt; &lt;/ul&gt; | 
 **endDate** | **string** | The ending date and time for data selection, see &#x60;start_date&#x60; description for details. | 
 **date** | **string** | Specifies the exact date to get the data for. Could be the exact date, e.g. &#x60;2021-10-27&#x60;, or in human language &#x60;today&#x60; or &#x60;yesterday&#x60; | 
 **order** | [**OrderEnum**](OrderEnum.md) | Sorting order of the output | [default to &quot;desc&quot;]
 **prepost** | **bool** | Returns quotes that include pre-market and post-market data. Only for the &#x60;Pro&#x60; plan (individual) and &#x60;Venture&#x60; plan (business) and above. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume | [default to false]
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **dp** | **int64** | Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. By default, the number of decimal places is automatically determined based on the values provided | [default to -1]
 **previousClose** | **bool** | A boolean parameter to include the previous closing price in the time_series data. If true, adds previous bar close price value to the current object | [default to false]
 **adjust** | [**AdjustEnum**](AdjustEnum.md) | Adjusting mode for prices | [default to &quot;splits&quot;]

### Return type

[**GetTimeSeries200Response**](GetTimeSeries200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeSeriesCross

> GetTimeSeriesCross200Response GetTimeSeriesCross(ctx).Base(base).Quote(quote).Interval(interval).BaseType(baseType).BaseExchange(baseExchange).BaseMicCode(baseMicCode).QuoteType(quoteType).QuoteExchange(quoteExchange).QuoteMicCode(quoteMicCode).Outputsize(outputsize).Format(format).Delimiter(delimiter).Prepost(prepost).StartDate(startDate).EndDate(endDate).Adjust(adjust).Dp(dp).Timezone(timezone).Execute()

Time series cross



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
	base := "JPY" // string | Base currency symbol
	quote := "BTC" // string | Quote currency symbol
	interval := openapiclient.IntervalEnum("1min") // IntervalEnum | Interval between two consecutive points in time series
	baseType := "Physical Currency" // string | Base instrument type according to the `/instrument_type` endpoint (optional)
	baseExchange := "Binance" // string | Base exchange (optional)
	baseMicCode := "XNGS" // string | Base MIC code (optional)
	quoteType := "Digital Currency" // string | Quote instrument type according to the `/instrument_type` endpoint (optional)
	quoteExchange := "Coinbase" // string | Quote exchange (optional)
	quoteMicCode := "XNYS" // string | Quote MIC code (optional)
	outputsize := int64(30) // int64 | Number of data points to retrieve. Supports values in the range from `1` to `5000`. Default `30` when no date parameters are set, otherwise set to maximum (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | Format of the response data (optional) (default to "JSON")
	delimiter := ";" // string | Delimiter used in CSV file (optional) (default to ";")
	prepost := true // bool | Only for the `Pro` plan (individual) and `Venture` plan (business) and above. Available at the `1min`, `5min`, `15min`, and `30min` intervals for US equities. Open, high, low, close values are supplied without volume. (optional) (default to false)
	startDate := "2025-01-01" // string | Start date for the time series data (optional)
	endDate := "2025-01-31" // string | End date for the time series data (optional)
	adjust := true // bool | Specifies if there should be an adjustment (optional) (default to true)
	dp := int64(5) // int64 | Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. (optional) (default to 5)
	timezone := "UTC" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>Exchange</code> for local exchange time</li> <li>2. <code>UTC</code> for datetime at universal UTC standard</li> <li>3. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a>.</li> </ul> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetTimeSeriesCross(context.Background()).Base(base).Quote(quote).Interval(interval).BaseType(baseType).BaseExchange(baseExchange).BaseMicCode(baseMicCode).QuoteType(quoteType).QuoteExchange(quoteExchange).QuoteMicCode(quoteMicCode).Outputsize(outputsize).Format(format).Delimiter(delimiter).Prepost(prepost).StartDate(startDate).EndDate(endDate).Adjust(adjust).Dp(dp).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetTimeSeriesCross``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeSeriesCross`: GetTimeSeriesCross200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetTimeSeriesCross`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeSeriesCrossRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **base** | **string** | Base currency symbol | 
 **quote** | **string** | Quote currency symbol | 
 **interval** | [**IntervalEnum**](IntervalEnum.md) | Interval between two consecutive points in time series | 
 **baseType** | **string** | Base instrument type according to the &#x60;/instrument_type&#x60; endpoint | 
 **baseExchange** | **string** | Base exchange | 
 **baseMicCode** | **string** | Base MIC code | 
 **quoteType** | **string** | Quote instrument type according to the &#x60;/instrument_type&#x60; endpoint | 
 **quoteExchange** | **string** | Quote exchange | 
 **quoteMicCode** | **string** | Quote MIC code | 
 **outputsize** | **int64** | Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;5000&#x60;. Default &#x60;30&#x60; when no date parameters are set, otherwise set to maximum | 
 **format** | [**FormatEnum**](FormatEnum.md) | Format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Delimiter used in CSV file | [default to &quot;;&quot;]
 **prepost** | **bool** | Only for the &#x60;Pro&#x60; plan (individual) and &#x60;Venture&#x60; plan (business) and above. Available at the &#x60;1min&#x60;, &#x60;5min&#x60;, &#x60;15min&#x60;, and &#x60;30min&#x60; intervals for US equities. Open, high, low, close values are supplied without volume. | [default to false]
 **startDate** | **string** | Start date for the time series data | 
 **endDate** | **string** | End date for the time series data | 
 **adjust** | **bool** | Specifies if there should be an adjustment | [default to true]
 **dp** | **int64** | Specifies the number of decimal places for floating values. Should be in range [0, 11] inclusive. | [default to 5]
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | 

### Return type

[**GetTimeSeriesCross200Response**](GetTimeSeriesCross200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

