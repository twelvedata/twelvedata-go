# \FundamentalsAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalanceSheet**](FundamentalsAPI.md#GetBalanceSheet) | **Get** /balance_sheet | Balance sheet
[**GetBalanceSheetConsolidated**](FundamentalsAPI.md#GetBalanceSheetConsolidated) | **Get** /balance_sheet/consolidated | Balance sheet consolidated
[**GetCashFlow**](FundamentalsAPI.md#GetCashFlow) | **Get** /cash_flow | Cash flow
[**GetCashFlowConsolidated**](FundamentalsAPI.md#GetCashFlowConsolidated) | **Get** /cash_flow/consolidated | Cash flow consolidated
[**GetDividends**](FundamentalsAPI.md#GetDividends) | **Get** /dividends | Dividends
[**GetDividendsCalendar**](FundamentalsAPI.md#GetDividendsCalendar) | **Get** /dividends_calendar | Dividends calendar
[**GetEarnings**](FundamentalsAPI.md#GetEarnings) | **Get** /earnings | Earnings
[**GetEarningsCalendar**](FundamentalsAPI.md#GetEarningsCalendar) | **Get** /earnings_calendar | Earnings calendar
[**GetIncomeStatement**](FundamentalsAPI.md#GetIncomeStatement) | **Get** /income_statement | Income statement
[**GetIncomeStatementConsolidated**](FundamentalsAPI.md#GetIncomeStatementConsolidated) | **Get** /income_statement/consolidated | Income statement consolidated
[**GetIpoCalendar**](FundamentalsAPI.md#GetIpoCalendar) | **Get** /ipo_calendar | IPO calendar
[**GetKeyExecutives**](FundamentalsAPI.md#GetKeyExecutives) | **Get** /key_executives | Key executives
[**GetLastChanges**](FundamentalsAPI.md#GetLastChanges) | **Get** /last_change/{endpoint} | Last changes
[**GetLogo**](FundamentalsAPI.md#GetLogo) | **Get** /logo | Logo
[**GetMarketCap**](FundamentalsAPI.md#GetMarketCap) | **Get** /market_cap | Market capitalization
[**GetProfile**](FundamentalsAPI.md#GetProfile) | **Get** /profile | Profile
[**GetSplits**](FundamentalsAPI.md#GetSplits) | **Get** /splits | Splits
[**GetSplitsCalendar**](FundamentalsAPI.md#GetSplitsCalendar) | **Get** /splits_calendar | Splits calendar
[**GetStatistics**](FundamentalsAPI.md#GetStatistics) | **Get** /statistics | Statistics
[**PressReleasesListParameters**](FundamentalsAPI.md#PressReleasesListParameters) | **Get** /press_releases | Press releases



## GetBalanceSheet

> GetBalanceSheet200Response GetBalanceSheet(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()

Balance sheet



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	period := openapiclient.PeriodEnum("annual") // PeriodEnum | The reporting period for the balane sheet data (optional) (default to "annual")
	startDate := "2024-01-01" // string | Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format `2006-01-02` (optional)
	endDate := "2024-05-01" // string | End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format `2006-01-02` (optional)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 6)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetBalanceSheet(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetBalanceSheet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalanceSheet`: GetBalanceSheet200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetBalanceSheet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **period** | [**PeriodEnum**](PeriodEnum.md) | The reporting period for the balane sheet data | [default to &quot;annual&quot;]
 **startDate** | **string** | Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60; | 
 **outputsize** | **int64** | Number of records in response | [default to 6]

### Return type

[**GetBalanceSheet200Response**](GetBalanceSheet200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalanceSheetConsolidated

> GetBalanceSheetConsolidated200Response GetBalanceSheetConsolidated(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()

Balance sheet consolidated



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	period := openapiclient.PeriodEnum("annual") // PeriodEnum | The reporting period for the balance sheet data. (optional) (default to "annual")
	startDate := "startDate_example" // string | Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format `2006-01-02` (optional)
	endDate := "endDate_example" // string | End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format `2006-01-02` (optional)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 6)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetBalanceSheetConsolidated(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetBalanceSheetConsolidated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalanceSheetConsolidated`: GetBalanceSheetConsolidated200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetBalanceSheetConsolidated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceSheetConsolidatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **period** | [**PeriodEnum**](PeriodEnum.md) | The reporting period for the balance sheet data. | [default to &quot;annual&quot;]
 **startDate** | **string** | Begin date for filtering items by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | End date for filtering items by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60; | 
 **outputsize** | **int64** | Number of records in response | [default to 6]

### Return type

[**GetBalanceSheetConsolidated200Response**](GetBalanceSheetConsolidated200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCashFlow

> GetCashFlow200Response GetCashFlow(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()

Cash flow



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	period := openapiclient.PeriodEnum("annual") // PeriodEnum | The reporting period for the cash flow statements (optional) (default to "annual")
	startDate := "2024-01-01" // string | Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format `2006-01-02` (optional)
	endDate := "2024-12-31" // string | End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format `2006-01-02` (optional)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 6)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetCashFlow(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetCashFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCashFlow`: GetCashFlow200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetCashFlow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCashFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **period** | [**PeriodEnum**](PeriodEnum.md) | The reporting period for the cash flow statements | [default to &quot;annual&quot;]
 **startDate** | **string** | Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format &#x60;2006-01-02&#x60; | 
 **outputsize** | **int64** | Number of records in response | [default to 6]

### Return type

[**GetCashFlow200Response**](GetCashFlow200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCashFlowConsolidated

> GetCashFlowConsolidated200Response GetCashFlowConsolidated(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()

Cash flow consolidated



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	period := openapiclient.PeriodEnum("annual") // PeriodEnum | The reporting period for the cash flow statements (optional) (default to "annual")
	startDate := "2024-01-01" // string | Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format `2006-01-02` (optional)
	endDate := "2024-12-31" // string | End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format `2006-01-02` (optional)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 6)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetCashFlowConsolidated(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetCashFlowConsolidated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCashFlowConsolidated`: GetCashFlowConsolidated200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetCashFlowConsolidated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCashFlowConsolidatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **period** | [**PeriodEnum**](PeriodEnum.md) | The reporting period for the cash flow statements | [default to &quot;annual&quot;]
 **startDate** | **string** | Start date for filtering cash flow statements. Only cash flow statements with fiscal dates on or after this date will be included. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | End date for filtering cash flow statements. Only cash flow statements with fiscal dates on or before this date will be included. Format &#x60;2006-01-02&#x60; | 
 **outputsize** | **int64** | Number of records in response | [default to 6]

### Return type

[**GetCashFlowConsolidated200Response**](GetCashFlowConsolidated200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDividends

> GetDividends200Response GetDividends(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Range_(range_).StartDate(startDate).EndDate(endDate).Adjust(adjust).Execute()

Dividends



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "US" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	range_ := openapiclient.RangeEnum("last") // RangeEnum | Specifies the time range for which to retrieve dividend data. Accepts values such as `last` (most recent dividend), `next` (upcoming dividend), `1m` - `5y` for respective periods, or `full` for all available data. If provided together with `start_date` and/or `end_date`, this parameter takes precedence. (optional) (default to "last")
	startDate := "2024-01-01" // string | Start date for the dividend data query. Only dividends with dates on or after this date will be returned. Format `2006-01-02`. If provided together with `range` parameter, `range` will take precedence. (optional)
	endDate := "2024-12-31" // string | End date for the dividend data query. Only dividends with dates on or after this date will be returned. Format `2006-01-02`. If provided together with `range` parameter, `range` will take precedence. (optional)
	adjust := true // bool | Specifies if there should be an adjustment (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetDividends(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Range_(range_).StartDate(startDate).EndDate(endDate).Adjust(adjust).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetDividends``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDividends`: GetDividends200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetDividends`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDividendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **range_** | [**RangeEnum**](RangeEnum.md) | Specifies the time range for which to retrieve dividend data. Accepts values such as &#x60;last&#x60; (most recent dividend), &#x60;next&#x60; (upcoming dividend), &#x60;1m&#x60; - &#x60;5y&#x60; for respective periods, or &#x60;full&#x60; for all available data. If provided together with &#x60;start_date&#x60; and/or &#x60;end_date&#x60;, this parameter takes precedence. | [default to &quot;last&quot;]
 **startDate** | **string** | Start date for the dividend data query. Only dividends with dates on or after this date will be returned. Format &#x60;2006-01-02&#x60;. If provided together with &#x60;range&#x60; parameter, &#x60;range&#x60; will take precedence. | 
 **endDate** | **string** | End date for the dividend data query. Only dividends with dates on or after this date will be returned. Format &#x60;2006-01-02&#x60;. If provided together with &#x60;range&#x60; parameter, &#x60;range&#x60; will take precedence. | 
 **adjust** | **bool** | Specifies if there should be an adjustment | [default to true]

### Return type

[**GetDividends200Response**](GetDividends200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDividendsCalendar

> []DividendsCalendarItem GetDividendsCalendar(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Page(page).Execute()

Dividends calendar



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "US" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	startDate := "2024-01-01" // string | Start date for the dividends calendar query. Only dividends with ex-dates on or after this date will be returned. Format `2006-01-02` (optional)
	endDate := "2024-12-31" // string | End date for the dividends calendar query. Only dividends with ex-dates on or before this date will be returned. Format `2006-01-02` (optional)
	outputsize := int64(789) // int64 | Number of data points to retrieve. Supports values in the range from `1` to `500`. Default `100` when no date parameters are set, otherwise set to maximum (optional) (default to 100)
	page := int64(789) // int64 | Page number (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetDividendsCalendar(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetDividendsCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDividendsCalendar`: []DividendsCalendarItem
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetDividendsCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDividendsCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **startDate** | **string** | Start date for the dividends calendar query. Only dividends with ex-dates on or after this date will be returned. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | End date for the dividends calendar query. Only dividends with ex-dates on or before this date will be returned. Format &#x60;2006-01-02&#x60; | 
 **outputsize** | **int64** | Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;500&#x60;. Default &#x60;100&#x60; when no date parameters are set, otherwise set to maximum | [default to 100]
 **page** | **int64** | Page number | [default to 1]

### Return type

[**[]DividendsCalendarItem**](DividendsCalendarItem.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEarnings

> GetEarnings200Response GetEarnings(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Period(period).Outputsize(outputsize).Format(format).Delimiter(delimiter).StartDate(startDate).EndDate(endDate).Dp(dp).Execute()

Earnings



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	type_ := openapiclient.TypeEnum("American Depositary Receipt") // TypeEnum | The asset class to which the instrument belongs (optional)
	period := openapiclient.PeriodEarningsEnum("latest") // PeriodEarningsEnum | Type of earning, returns only 1 record. When is not empty, dates and outputsize parameters are ignored (optional)
	outputsize := int64(789) // int64 | Number of data points to retrieve. Supports values in the range from `1` to `1000`. Default `10` when no date parameters are set, otherwise set to maximum (optional) (default to 10)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | The format of the response data (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | The separator used in the CSV response data (optional) (default to ";")
	startDate := "2024-04-01" // string | The date from which the data is requested. The date format is `YYYY-MM-DD`. (optional)
	endDate := "2024-04-30" // string | The date to which the data is requested. The date format is `YYYY-MM-DD`. (optional)
	dp := int64(789) // int64 | The number of decimal places in the response data. Should be in range [0,11] inclusive (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetEarnings(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Type_(type_).Period(period).Outputsize(outputsize).Format(format).Delimiter(delimiter).StartDate(startDate).EndDate(endDate).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetEarnings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEarnings`: GetEarnings200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetEarnings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEarningsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **type_** | [**TypeEnum**](TypeEnum.md) | The asset class to which the instrument belongs | 
 **period** | [**PeriodEarningsEnum**](PeriodEarningsEnum.md) | Type of earning, returns only 1 record. When is not empty, dates and outputsize parameters are ignored | 
 **outputsize** | **int64** | Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;1000&#x60;. Default &#x60;10&#x60; when no date parameters are set, otherwise set to maximum | [default to 10]
 **format** | [**FormatEnum**](FormatEnum.md) | The format of the response data | [default to &quot;JSON&quot;]
 **delimiter** | **string** | The separator used in the CSV response data | [default to &quot;;&quot;]
 **startDate** | **string** | The date from which the data is requested. The date format is &#x60;YYYY-MM-DD&#x60;. | 
 **endDate** | **string** | The date to which the data is requested. The date format is &#x60;YYYY-MM-DD&#x60;. | 
 **dp** | **int64** | The number of decimal places in the response data. Should be in range [0,11] inclusive | [default to 2]

### Return type

[**GetEarnings200Response**](GetEarnings200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEarningsCalendar

> GetEarningsCalendar200Response GetEarningsCalendar(ctx).Exchange(exchange).MicCode(micCode).Country(country).Format(format).Delimiter(delimiter).StartDate(startDate).EndDate(endDate).Dp(dp).Execute()

Earnings calendar



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
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	format := openapiclient.FormatEnum("JSON") // FormatEnum | Value can be JSON or CSV (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the CSV file (optional) (default to ";")
	startDate := "2024-04-01" // string | Can be used separately and together with end_date. Format `2006-01-02` or `2006-01-02T15:04:05` (optional)
	endDate := "2024-04-30" // string | Can be used separately and together with start_date. Format `2006-01-02` or `2006-01-02T15:04:05` (optional)
	dp := int64(789) // int64 | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetEarningsCalendar(context.Background()).Exchange(exchange).MicCode(micCode).Country(country).Format(format).Delimiter(delimiter).StartDate(startDate).EndDate(endDate).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetEarningsCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEarningsCalendar`: GetEarningsCalendar200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetEarningsCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEarningsCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **format** | [**FormatEnum**](FormatEnum.md) | Value can be JSON or CSV | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the CSV file | [default to &quot;;&quot;]
 **startDate** | **string** | Can be used separately and together with end_date. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60; | 
 **endDate** | **string** | Can be used separately and together with start_date. Format &#x60;2006-01-02&#x60; or &#x60;2006-01-02T15:04:05&#x60; | 
 **dp** | **int64** | Specifies the number of decimal places for floating values. Should be in range [0,11] inclusive | [default to 2]

### Return type

[**GetEarningsCalendar200Response**](GetEarningsCalendar200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomeStatement

> GetIncomeStatement200Response GetIncomeStatement(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()

Income statement



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	period := openapiclient.PeriodEnum("annual") // PeriodEnum | The reporting period for the income statement data (optional) (default to "annual")
	startDate := "2024-01-01" // string | Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format `2006-01-02` (optional)
	endDate := "2024-12-31" // string | End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format `2006-01-02` (optional)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 6)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetIncomeStatement(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetIncomeStatement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomeStatement`: GetIncomeStatement200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetIncomeStatement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomeStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **period** | [**PeriodEnum**](PeriodEnum.md) | The reporting period for the income statement data | [default to &quot;annual&quot;]
 **startDate** | **string** | Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60; | 
 **outputsize** | **int64** | Number of records in response | [default to 6]

### Return type

[**GetIncomeStatement200Response**](GetIncomeStatement200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomeStatementConsolidated

> GetIncomeStatementConsolidated200Response GetIncomeStatementConsolidated(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()

Income statement consolidated



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	period := openapiclient.PeriodEnum("annual") // PeriodEnum | The reporting period for the income statement data (optional) (default to "annual")
	startDate := "2024-01-01" // string | Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format `2006-01-02` (optional)
	endDate := "2024-12-31" // string | End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format `2006-01-02` (optional)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 6)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetIncomeStatementConsolidated(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Period(period).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetIncomeStatementConsolidated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomeStatementConsolidated`: GetIncomeStatementConsolidated200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetIncomeStatementConsolidated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomeStatementConsolidatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **period** | [**PeriodEnum**](PeriodEnum.md) | The reporting period for the income statement data | [default to &quot;annual&quot;]
 **startDate** | **string** | Begin date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or after this date. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | End date for filtering income statements by fiscal date. Returns income statements with fiscal dates on or before this date. Format &#x60;2006-01-02&#x60; | 
 **outputsize** | **int64** | Number of records in response | [default to 6]

### Return type

[**GetIncomeStatementConsolidated200Response**](GetIncomeStatementConsolidated200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpoCalendar

> map[string][]GetIpoCalendar200ResponseValueInner GetIpoCalendar(ctx).Exchange(exchange).MicCode(micCode).Country(country).StartDate(startDate).EndDate(endDate).Execute()

IPO calendar



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
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	startDate := "2021-01-01" // string | The earliest IPO date to include in the results. Format: `2006-01-02` (optional)
	endDate := "2021-12-31" // string | The latest IPO date to include in the results. Format: `2006-01-02` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetIpoCalendar(context.Background()).Exchange(exchange).MicCode(micCode).Country(country).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetIpoCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpoCalendar`: map[string][]GetIpoCalendar200ResponseValueInner
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetIpoCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpoCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **startDate** | **string** | The earliest IPO date to include in the results. Format: &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | The latest IPO date to include in the results. Format: &#x60;2006-01-02&#x60; | 

### Return type

[**map[string][]GetIpoCalendar200ResponseValueInner**](array.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyExecutives

> GetKeyExecutives200Response GetKeyExecutives(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()

Key executives



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetKeyExecutives(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetKeyExecutives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyExecutives`: GetKeyExecutives200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetKeyExecutives`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyExecutivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 

### Return type

[**GetKeyExecutives200Response**](GetKeyExecutives200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastChanges

> GetLastChanges200Response GetLastChanges(ctx, endpoint).StartDate(startDate).Symbol(symbol).Exchange(exchange).MicCode(micCode).Country(country).Page(page).Outputsize(outputsize).Execute()

Last changes



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
	endpoint := openapiclient.EndpointEnum("price_target") // EndpointEnum | Endpoint name
	startDate := "2023-10-14T00:00:00" // string | The starting date and time for data selection, in `2006-01-02T15:04:05` format (optional)
	symbol := "AAPL" // string | Filter by symbol (optional)
	exchange := "NASDAQ" // string | Filter by exchange name (optional)
	micCode := "XNAS" // string | Filter by market identifier code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	page := int64(789) // int64 | Page number (optional) (default to 1)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetLastChanges(context.Background(), endpoint).StartDate(startDate).Symbol(symbol).Exchange(exchange).MicCode(micCode).Country(country).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetLastChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLastChanges`: GetLastChanges200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetLastChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpoint** | [**EndpointEnum**](.md) | Endpoint name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | The starting date and time for data selection, in &#x60;2006-01-02T15:04:05&#x60; format | 
 **symbol** | **string** | Filter by symbol | 
 **exchange** | **string** | Filter by exchange name | 
 **micCode** | **string** | Filter by market identifier code (MIC) under ISO 10383 standard | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **page** | **int64** | Page number | [default to 1]
 **outputsize** | **int64** | Number of records in response | [default to 30]

### Return type

[**GetLastChanges200Response**](GetLastChanges200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogo

> GetLogo200Response GetLogo(ctx).Symbol(symbol).Exchange(exchange).MicCode(micCode).Country(country).Execute()

Logo



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
	symbol := "BTC/USD" // string | The ticker symbol of an instrument for which data is requested, e.g., `AAPL`, `BTC/USD`, `EUR/USD`.
	exchange := "NASDAQ" // string | The exchange name where the instrument is traded, e.g., `NASDAQ`, `NSE` (optional)
	micCode := "XNAS" // string | The Market Identifier Code (MIC) of the exchange where the instrument is traded, e.g., `XNAS`, `XLON` (optional)
	country := "United States" // string | The country where the instrument is traded, e.g., `United States` or `US` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetLogo(context.Background()).Symbol(symbol).Exchange(exchange).MicCode(micCode).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogo`: GetLogo200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetLogo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested, e.g., &#x60;AAPL&#x60;, &#x60;BTC/USD&#x60;, &#x60;EUR/USD&#x60;. | 
 **exchange** | **string** | The exchange name where the instrument is traded, e.g., &#x60;NASDAQ&#x60;, &#x60;NSE&#x60; | 
 **micCode** | **string** | The Market Identifier Code (MIC) of the exchange where the instrument is traded, e.g., &#x60;XNAS&#x60;, &#x60;XLON&#x60; | 
 **country** | **string** | The country where the instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 

### Return type

[**GetLogo200Response**](GetLogo200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketCap

> GetMarketCap200Response GetMarketCap(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).StartDate(startDate).EndDate(endDate).Page(page).Outputsize(outputsize).Execute()

Market capitalization



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
	symbol := "AAPL" // string | Filter by symbol (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Filter by exchange name (optional)
	micCode := "XNAS" // string | Filter by market identifier code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	startDate := "2023-01-01" // string | Start date for market capitalization data retrieval. Data will be returned from this date onwards. Format `2006-01-02` (optional)
	endDate := "2023-12-31" // string | End date for market capitalization data retrieval. Data will be returned up to and including this date. Format `2006-01-02` (optional)
	page := int64(789) // int64 | Page number (optional) (default to 1)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetMarketCap(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).StartDate(startDate).EndDate(endDate).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetMarketCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketCap`: GetMarketCap200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetMarketCap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketCapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Filter by exchange name | 
 **micCode** | **string** | Filter by market identifier code (MIC) under ISO 10383 standard | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **startDate** | **string** | Start date for market capitalization data retrieval. Data will be returned from this date onwards. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | End date for market capitalization data retrieval. Data will be returned up to and including this date. Format &#x60;2006-01-02&#x60; | 
 **page** | **int64** | Page number | [default to 1]
 **outputsize** | **int64** | Number of records in response | [default to 10]

### Return type

[**GetMarketCap200Response**](GetMarketCap200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> GetProfile200Response GetProfile(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()

Profile



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetProfile(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: GetProfile200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 

### Return type

[**GetProfile200Response**](GetProfile200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSplits

> GetSplits200Response GetSplits(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Range_(range_).StartDate(startDate).EndDate(endDate).Execute()

Splits



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	range_ := openapiclient.RangeSplitsEnum("last") // RangeSplitsEnum | Range of data to be returned (optional) (default to "last")
	startDate := "2020-01-01" // string | The starting date for data selection. Format `2006-01-02` (optional)
	endDate := "2020-12-31" // string | The ending date for data selection. Format `2006-01-02` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetSplits(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Range_(range_).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetSplits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSplits`: GetSplits200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetSplits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSplitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **range_** | [**RangeSplitsEnum**](RangeSplitsEnum.md) | Range of data to be returned | [default to &quot;last&quot;]
 **startDate** | **string** | The starting date for data selection. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | The ending date for data selection. Format &#x60;2006-01-02&#x60; | 

### Return type

[**GetSplits200Response**](GetSplits200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSplitsCalendar

> []SplitsCalendarResponseItem GetSplitsCalendar(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Page(page).Execute()

Splits calendar



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)
	startDate := "2024-01-01" // string | The starting date (inclusive) for filtering split events in the calendar. Format `2006-01-02` (optional)
	endDate := "2024-12-31" // string | The ending date (inclusive) for filtering split events in the calendar. Format `2006-01-02` (optional)
	outputsize := int64(789) // int64 | Number of data points to retrieve. Supports values in the range from `1` to `500`. Default `100` when no date parameters are set, otherwise set to maximum (optional) (default to 100)
	page := "page_example" // string | Page number (optional) (default to "1")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetSplitsCalendar(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).StartDate(startDate).EndDate(endDate).Outputsize(outputsize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetSplitsCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSplitsCalendar`: []SplitsCalendarResponseItem
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetSplitsCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSplitsCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **startDate** | **string** | The starting date (inclusive) for filtering split events in the calendar. Format &#x60;2006-01-02&#x60; | 
 **endDate** | **string** | The ending date (inclusive) for filtering split events in the calendar. Format &#x60;2006-01-02&#x60; | 
 **outputsize** | **int64** | Number of data points to retrieve. Supports values in the range from &#x60;1&#x60; to &#x60;500&#x60;. Default &#x60;100&#x60; when no date parameters are set, otherwise set to maximum | [default to 100]
 **page** | **string** | Page number | [default to &quot;1&quot;]

### Return type

[**[]SplitsCalendarResponseItem**](SplitsCalendarResponseItem.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatistics

> GetStatistics200Response GetStatistics(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()

Statistics



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.GetStatistics(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.GetStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatistics`: GetStatistics200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.GetStatistics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 

### Return type

[**GetStatistics200Response**](GetStatistics200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PressReleasesListParameters

> PressReleasesListParameters200Response PressReleasesListParameters(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).StartDate(startDate).EndDate(endDate).Timezone(timezone).Language(language).Outputsize(outputsize).Execute()

Press releases



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
	symbol := "AAPL" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	startDate := "2025-12-01T00:00:00" // string | Begin date for filtering items. Returns press releases with release date on or after this date. Format `2025-12-24T02:07:00` (optional)
	endDate := "2025-12-31T23:59:00" // string | End date for filtering items. Returns press releases with release date on or before this date. Format `2025-12-24T02:07:00` (optional)
	timezone := "America/New_York" // string | Time zone for date filtering. Default is the identifier time zone. (optional)
	language := "en,en-US" // string | Comma-separated list of languages to filter press releases by language. (optional)
	outputsize := int64(5) // int64 | Number of latest press releases returned. Only used if no data range is specified. Maximum value is `10`.  type: number (optional) (default to 2)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundamentalsAPI.PressReleasesListParameters(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).StartDate(startDate).EndDate(endDate).Timezone(timezone).Language(language).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundamentalsAPI.PressReleasesListParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PressReleasesListParameters`: PressReleasesListParameters200Response
	fmt.Fprintf(os.Stdout, "Response from `FundamentalsAPI.PressReleasesListParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPressReleasesListParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **startDate** | **string** | Begin date for filtering items. Returns press releases with release date on or after this date. Format &#x60;2025-12-24T02:07:00&#x60; | 
 **endDate** | **string** | End date for filtering items. Returns press releases with release date on or before this date. Format &#x60;2025-12-24T02:07:00&#x60; | 
 **timezone** | **string** | Time zone for date filtering. Default is the identifier time zone. | 
 **language** | **string** | Comma-separated list of languages to filter press releases by language. | 
 **outputsize** | **int64** | Number of latest press releases returned. Only used if no data range is specified. Maximum value is &#x60;10&#x60;.  type: number | [default to 2]

### Return type

[**PressReleasesListParameters200Response**](PressReleasesListParameters200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

