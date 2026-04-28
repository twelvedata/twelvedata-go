# \EtfsAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetETFsFamily**](EtfsAPI.md#GetETFsFamily) | **Get** /etfs/family | ETFs families
[**GetETFsList**](EtfsAPI.md#GetETFsList) | **Get** /etfs/list | ETFs directory
[**GetETFsType**](EtfsAPI.md#GetETFsType) | **Get** /etfs/type | ETFs types
[**GetETFsWorld**](EtfsAPI.md#GetETFsWorld) | **Get** /etfs/world | ETF full data
[**GetETFsWorldComposition**](EtfsAPI.md#GetETFsWorldComposition) | **Get** /etfs/world/composition | Composition
[**GetETFsWorldPerformance**](EtfsAPI.md#GetETFsWorldPerformance) | **Get** /etfs/world/performance | Performance
[**GetETFsWorldRisk**](EtfsAPI.md#GetETFsWorldRisk) | **Get** /etfs/world/risk | Risk
[**GetETFsWorldSummary**](EtfsAPI.md#GetETFsWorldSummary) | **Get** /etfs/world/summary | Summary



## GetETFsFamily

> GetETFsFamily200Response GetETFsFamily(ctx).Country(country).FundFamily(fundFamily).Execute()

ETFs families



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
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	fundFamily := "iShares" // string | Filter by investment company that manages the fund (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsFamily(context.Background()).Country(country).FundFamily(fundFamily).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsFamily``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsFamily`: GetETFsFamily200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsFamily`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsFamilyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **fundFamily** | **string** | Filter by investment company that manages the fund | 

### Return type

[**GetETFsFamily200Response**](GetETFsFamily200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsList

> GetETFsList200Response GetETFsList(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Country(country).FundFamily(fundFamily).FundType(fundType).Page(page).Outputsize(outputsize).Execute()

ETFs directory



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
	symbol := "IVV" // string | Filter by symbol (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cik := "95953" // string | The CIK of an instrument for which data is requested (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	fundFamily := "iShares" // string | Filter by investment company that manages the fund (optional)
	fundType := "Large Blend" // string | Filter by the type of fund (optional)
	page := int64(789) // int64 | Page number (optional) (default to 1)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsList(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Country(country).FundFamily(fundFamily).FundType(fundType).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsList`: GetETFsList200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cik** | **string** | The CIK of an instrument for which data is requested | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **fundFamily** | **string** | Filter by investment company that manages the fund | 
 **fundType** | **string** | Filter by the type of fund | 
 **page** | **int64** | Page number | [default to 1]
 **outputsize** | **int64** | Number of records in response | [default to 50]

### Return type

[**GetETFsList200Response**](GetETFsList200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsType

> GetETFsType200Response GetETFsType(ctx).Country(country).FundType(fundType).Execute()

ETFs types



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
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	fundType := "Large Blend" // string | Filter by the type of fund (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsType(context.Background()).Country(country).FundType(fundType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsType`: GetETFsType200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **fundType** | **string** | Filter by the type of fund | 

### Return type

[**GetETFsType200Response**](GetETFsType200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsWorld

> GetETFsWorld200Response GetETFsWorld(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

ETF full data



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
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorld(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorld`: GetETFsWorld200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorld`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorld200Response**](GetETFsWorld200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsWorldComposition

> GetETFsWorldComposition200Response GetETFsWorldComposition(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Composition



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
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorldComposition(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorldComposition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorldComposition`: GetETFsWorldComposition200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorldComposition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorldComposition200Response**](GetETFsWorldComposition200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsWorldPerformance

> GetETFsWorldPerformance200Response GetETFsWorldPerformance(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Performance



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
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorldPerformance(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorldPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorldPerformance`: GetETFsWorldPerformance200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorldPerformance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorldPerformance200Response**](GetETFsWorldPerformance200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsWorldRisk

> GetETFsWorldRisk200Response GetETFsWorldRisk(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Risk



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
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorldRisk(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorldRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorldRisk`: GetETFsWorldRisk200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorldRisk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorldRisk200Response**](GetETFsWorldRisk200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetETFsWorldSummary

> GetETFsWorldSummary200Response GetETFsWorldSummary(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Summary



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
	symbol := "IVV" // string | Symbol ticker of etf (optional)
	figi := "BBG000BVZ697" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US4642872000" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "464287200" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtfsAPI.GetETFsWorldSummary(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtfsAPI.GetETFsWorldSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetETFsWorldSummary`: GetETFsWorldSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `EtfsAPI.GetETFsWorldSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetETFsWorldSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of etf | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetETFsWorldSummary200Response**](GetETFsWorldSummary200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

