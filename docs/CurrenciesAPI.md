# \CurrenciesAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrencyConversion**](CurrenciesAPI.md#GetCurrencyConversion) | **Get** /currency_conversion | Currency conversion
[**GetExchangeRate**](CurrenciesAPI.md#GetExchangeRate) | **Get** /exchange_rate | Exchange rate



## GetCurrencyConversion

> GetCurrencyConversion200Response GetCurrencyConversion(ctx).Symbol(symbol).Amount(amount).Date(date).Format(format).Delimiter(delimiter).Dp(dp).Timezone(timezone).Execute()

Currency conversion



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
	symbol := "EUR/USD" // string | The currency pair you want to request can be either forex or cryptocurrency. Slash(`/`) delimiter is used. E.g. `EUR/USD` or `BTC/ETH` will be correct
	amount := float64(100) // float64 | Amount of base currency to be converted into quote currency. Supports values in the range from `0` and above
	date := "2006-01-02T15:04:05" // string | If not null, will use exchange rate from a specific date or time. Format `2006-01-02` or `2006-01-02T15:04:05`. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | Value can be `JSON` or `CSV`. Default `JSON` (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the `CSV` file. Default semicolon `;` (optional) (default to ";")
	dp := int64(789) // int64 | The number of decimal places for the data (optional) (default to 5)
	timezone := "UTC" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>Exchange</code> for local exchange time</li> <li>2. <code>UTC</code> for datetime at universal UTC standard</li> <li>3. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a>.</li> </ul> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrenciesAPI.GetCurrencyConversion(context.Background()).Symbol(symbol).Amount(amount).Date(date).Format(format).Delimiter(delimiter).Dp(dp).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesAPI.GetCurrencyConversion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrencyConversion`: GetCurrencyConversion200Response
	fmt.Fprintf(os.Stdout, "Response from `CurrenciesAPI.GetCurrencyConversion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrencyConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The currency pair you want to request can be either forex or cryptocurrency. Slash(&#x60;/&#x60;) delimiter is used. E.g. &#x60;EUR/USD&#x60; or &#x60;BTC/ETH&#x60; will be correct | 
 **amount** | **float64** | Amount of base currency to be converted into quote currency. Supports values in the range from &#x60;0&#x60; and above | 
 **date** | **string** | If not null, will use exchange rate from a specific date or time. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone | 
 **format** | [**FormatEnum**](FormatEnum.md) | Value can be &#x60;JSON&#x60; or &#x60;CSV&#x60;. Default &#x60;JSON&#x60; | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the &#x60;CSV&#x60; file. Default semicolon &#x60;;&#x60; | [default to &quot;;&quot;]
 **dp** | **int64** | The number of decimal places for the data | [default to 5]
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | 

### Return type

[**GetCurrencyConversion200Response**](GetCurrencyConversion200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeRate

> GetExchangeRate200Response GetExchangeRate(ctx).Symbol(symbol).Date(date).Format(format).Delimiter(delimiter).Dp(dp).Timezone(timezone).Execute()

Exchange rate



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
	symbol := "EUR/USD" // string | The currency pair you want to request can be either forex or cryptocurrency. Slash(`/`) delimiter is used. E.g. `EUR/USD` or `BTC/ETH` will be correct
	date := "2006-01-02T15:04:05" // string | If not null, will use exchange rate from a specific date or time. Format `2006-01-02` or `2006-01-02T15:04:05`. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | Value can be `JSON` or `CSV`. Default `JSON` (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the `CSV` file. Default semicolon `;` (optional) (default to ";")
	dp := int64(789) // int64 | The number of decimal places for the data (optional) (default to 5)
	timezone := "UTC" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>Exchange</code> for local exchange time</li> <li>2. <code>UTC</code> for datetime at universal UTC standard</li> <li>3. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a>.</li> </ul> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrenciesAPI.GetExchangeRate(context.Background()).Symbol(symbol).Date(date).Format(format).Delimiter(delimiter).Dp(dp).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesAPI.GetExchangeRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeRate`: GetExchangeRate200Response
	fmt.Fprintf(os.Stdout, "Response from `CurrenciesAPI.GetExchangeRate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The currency pair you want to request can be either forex or cryptocurrency. Slash(&#x60;/&#x60;) delimiter is used. E.g. &#x60;EUR/USD&#x60; or &#x60;BTC/ETH&#x60; will be correct | 
 **date** | **string** | If not null, will use exchange rate from a specific date or time. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60;. Is set in the local exchange time zone, use timezone parameter to specify a specific time zone | 
 **format** | [**FormatEnum**](FormatEnum.md) | Value can be &#x60;JSON&#x60; or &#x60;CSV&#x60;. Default &#x60;JSON&#x60; | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the &#x60;CSV&#x60; file. Default semicolon &#x60;;&#x60; | [default to &quot;;&quot;]
 **dp** | **int64** | The number of decimal places for the data | [default to 5]
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;Exchange&lt;/code&gt; for local exchange time&lt;/li&gt; &lt;li&gt;2. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;3. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | 

### Return type

[**GetExchangeRate200Response**](GetExchangeRate200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

