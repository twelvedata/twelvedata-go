# \RegulatoryAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDirectHolders**](RegulatoryAPI.md#GetDirectHolders) | **Get** /direct_holders | Direct holders
[**GetEdgarFilingsArchive**](RegulatoryAPI.md#GetEdgarFilingsArchive) | **Get** /edgar_filings/archive | EDGAR fillings
[**GetFundHolders**](RegulatoryAPI.md#GetFundHolders) | **Get** /fund_holders | Fund holders
[**GetInsiderTransactions**](RegulatoryAPI.md#GetInsiderTransactions) | **Get** /insider_transactions | Insider transaction
[**GetInstitutionalHolders**](RegulatoryAPI.md#GetInstitutionalHolders) | **Get** /institutional_holders | Institutional holders
[**GetSourceSanctionedEntities**](RegulatoryAPI.md#GetSourceSanctionedEntities) | **Get** /sanctions/{source} | Sanctioned entities
[**GetTaxInfo**](RegulatoryAPI.md#GetTaxInfo) | **Get** /tax_info | Tax information



## GetDirectHolders

> GetDirectHolders200Response GetDirectHolders(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()

Direct holders



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
	symbol := "7203" // string | Symbol ticker of instrument. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., `United States` or `US` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegulatoryAPI.GetDirectHolders(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegulatoryAPI.GetDirectHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectHolders`: GetDirectHolders200Response
	fmt.Fprintf(os.Stdout, "Response from `RegulatoryAPI.GetDirectHolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectHoldersRequest struct via the builder pattern


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

[**GetDirectHolders200Response**](GetDirectHolders200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEdgarFilingsArchive

> GetEdgarFilingsArchive200Response GetEdgarFilingsArchive(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).FormType(formType).FilledFrom(filledFrom).FilledTo(filledTo).Page(page).PageSize(pageSize).Execute()

EDGAR fillings



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
	exchange := "NASDAQ" // string | Filter by exchange name (optional)
	micCode := "XNGS" // string | Filter by market identifier code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	formType := "8-K" // string | Filter by form types, example `8-K`, `EX-1.1` (optional)
	filledFrom := "2024-01-01" // string | Filter by filled time from (optional)
	filledTo := "2024-01-01" // string | Filter by filled time to (optional)
	page := int64(789) // int64 | Page number (optional) (default to 1)
	pageSize := int64(789) // int64 | Number of records in response (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegulatoryAPI.GetEdgarFilingsArchive(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).FormType(formType).FilledFrom(filledFrom).FilledTo(filledTo).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegulatoryAPI.GetEdgarFilingsArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEdgarFilingsArchive`: GetEdgarFilingsArchive200Response
	fmt.Fprintf(os.Stdout, "Response from `RegulatoryAPI.GetEdgarFilingsArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgarFilingsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Filter by exchange name | 
 **micCode** | **string** | Filter by market identifier code (MIC) under ISO 10383 standard | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **formType** | **string** | Filter by form types, example &#x60;8-K&#x60;, &#x60;EX-1.1&#x60; | 
 **filledFrom** | **string** | Filter by filled time from | 
 **filledTo** | **string** | Filter by filled time to | 
 **page** | **int64** | Page number | [default to 1]
 **pageSize** | **int64** | Number of records in response | [default to 10]

### Return type

[**GetEdgarFilingsArchive200Response**](GetEdgarFilingsArchive200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundHolders

> GetFundHolders200Response GetFundHolders(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()

Fund holders



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
	resp, r, err := apiClient.RegulatoryAPI.GetFundHolders(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegulatoryAPI.GetFundHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFundHolders`: GetFundHolders200Response
	fmt.Fprintf(os.Stdout, "Response from `RegulatoryAPI.GetFundHolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFundHoldersRequest struct via the builder pattern


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

[**GetFundHolders200Response**](GetFundHolders200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsiderTransactions

> GetInsiderTransactions200Response GetInsiderTransactions(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()

Insider transaction



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
	symbol := "AAPL" // string | The ticker symbol of an instrument for which data is requested, e.g., `AAPL`, `TSLA`. For preffered stocks use dot(.) delimiter. E.g. `BRK.A` or `BRK.B` will be correct (optional)
	figi := "BBG000B9Y5X2" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded, e.g., `Nasdaq`, `NSE` (optional)
	micCode := "XNAS" // string | Market Identifier Code (MIC) under ISO 10383 standard (optional)
	country := "United States" // string | Country where instrument is traded, e.g., United States or US. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegulatoryAPI.GetInsiderTransactions(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegulatoryAPI.GetInsiderTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInsiderTransactions`: GetInsiderTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `RegulatoryAPI.GetInsiderTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInsiderTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested, e.g., &#x60;AAPL&#x60;, &#x60;TSLA&#x60;. For preffered stocks use dot(.) delimiter. E.g. &#x60;BRK.A&#x60; or &#x60;BRK.B&#x60; will be correct | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Exchange where instrument is traded, e.g., &#x60;Nasdaq&#x60;, &#x60;NSE&#x60; | 
 **micCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
 **country** | **string** | Country where instrument is traded, e.g., United States or US. | 

### Return type

[**GetInsiderTransactions200Response**](GetInsiderTransactions200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstitutionalHolders

> GetInstitutionalHolders200Response GetInstitutionalHolders(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()

Institutional holders



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
	resp, r, err := apiClient.RegulatoryAPI.GetInstitutionalHolders(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).MicCode(micCode).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegulatoryAPI.GetInstitutionalHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstitutionalHolders`: GetInstitutionalHolders200Response
	fmt.Fprintf(os.Stdout, "Response from `RegulatoryAPI.GetInstitutionalHolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstitutionalHoldersRequest struct via the builder pattern


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

[**GetInstitutionalHolders200Response**](GetInstitutionalHolders200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceSanctionedEntities

> GetSourceSanctionedEntities200Response GetSourceSanctionedEntities(ctx, source).Execute()

Sanctioned entities



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
	source := openapiclient.SourceEnum("ofac") // SourceEnum | Sanctions source

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegulatoryAPI.GetSourceSanctionedEntities(context.Background(), source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegulatoryAPI.GetSourceSanctionedEntities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceSanctionedEntities`: GetSourceSanctionedEntities200Response
	fmt.Fprintf(os.Stdout, "Response from `RegulatoryAPI.GetSourceSanctionedEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | [**SourceEnum**](.md) | Sanctions source | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceSanctionedEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSourceSanctionedEntities200Response**](GetSourceSanctionedEntities200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxInfo

> GetTaxInfo200Response GetTaxInfo(ctx).Symbol(symbol).Isin(isin).Figi(figi).Cusip(cusip).Exchange(exchange).MicCode(micCode).Execute()

Tax information



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
	symbol := "SKYQ" // string | The ticker symbol of an instrument for which data is requested, e.g., `SKYQ`, `AIRE`, `ALM:BME`, `HSI:HKEX`. (optional)
	isin := "US5949181045" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	figi := "BBG019XJT9D6" // string | The FIGI of an instrument for which data is requested. This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	exchange := "Nasdaq" // string | The exchange name where the instrument is traded, e.g., `Nasdaq`, `Euronext` (optional)
	micCode := "XNAS" // string | The Market Identifier Code (MIC) of the exchange where the instrument is traded, e.g., `XNAS`, `XLON` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegulatoryAPI.GetTaxInfo(context.Background()).Symbol(symbol).Isin(isin).Figi(figi).Cusip(cusip).Exchange(exchange).MicCode(micCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegulatoryAPI.GetTaxInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxInfo`: GetTaxInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `RegulatoryAPI.GetTaxInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The ticker symbol of an instrument for which data is requested, e.g., &#x60;SKYQ&#x60;, &#x60;AIRE&#x60;, &#x60;ALM:BME&#x60;, &#x60;HSI:HKEX&#x60;. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **figi** | **string** | The FIGI of an instrument for which data is requested. This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | The exchange name where the instrument is traded, e.g., &#x60;Nasdaq&#x60;, &#x60;Euronext&#x60; | 
 **micCode** | **string** | The Market Identifier Code (MIC) of the exchange where the instrument is traded, e.g., &#x60;XNAS&#x60;, &#x60;XLON&#x60; | 

### Return type

[**GetTaxInfo200Response**](GetTaxInfo200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

