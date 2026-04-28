# \MutualFundsAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMutualFundsFamily**](MutualFundsAPI.md#GetMutualFundsFamily) | **Get** /mutual_funds/family | MFs families
[**GetMutualFundsList**](MutualFundsAPI.md#GetMutualFundsList) | **Get** /mutual_funds/list | MFs directory
[**GetMutualFundsType**](MutualFundsAPI.md#GetMutualFundsType) | **Get** /mutual_funds/type | MFs types
[**GetMutualFundsWorld**](MutualFundsAPI.md#GetMutualFundsWorld) | **Get** /mutual_funds/world | MF full data
[**GetMutualFundsWorldComposition**](MutualFundsAPI.md#GetMutualFundsWorldComposition) | **Get** /mutual_funds/world/composition | Composition
[**GetMutualFundsWorldPerformance**](MutualFundsAPI.md#GetMutualFundsWorldPerformance) | **Get** /mutual_funds/world/performance | Performance
[**GetMutualFundsWorldPurchaseInfo**](MutualFundsAPI.md#GetMutualFundsWorldPurchaseInfo) | **Get** /mutual_funds/world/purchase_info | Purchase info
[**GetMutualFundsWorldRatings**](MutualFundsAPI.md#GetMutualFundsWorldRatings) | **Get** /mutual_funds/world/ratings | Ratings
[**GetMutualFundsWorldRisk**](MutualFundsAPI.md#GetMutualFundsWorldRisk) | **Get** /mutual_funds/world/risk | Risk
[**GetMutualFundsWorldSummary**](MutualFundsAPI.md#GetMutualFundsWorldSummary) | **Get** /mutual_funds/world/summary | Summary
[**GetMutualFundsWorldSustainability**](MutualFundsAPI.md#GetMutualFundsWorldSustainability) | **Get** /mutual_funds/world/sustainability | Sustainability



## GetMutualFundsFamily

> GetMutualFundsFamily200Response GetMutualFundsFamily(ctx).FundFamily(fundFamily).Country(country).Execute()

MFs families



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
	fundFamily := "Jackson National" // string | Filter by investment company that manages the fund (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsFamily(context.Background()).FundFamily(fundFamily).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsFamily``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsFamily`: GetMutualFundsFamily200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsFamily`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsFamilyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fundFamily** | **string** | Filter by investment company that manages the fund | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 

### Return type

[**GetMutualFundsFamily200Response**](GetMutualFundsFamily200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsList

> GetMutualFundsList200Response GetMutualFundsList(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Country(country).FundFamily(fundFamily).FundType(fundType).PerformanceRating(performanceRating).RiskRating(riskRating).Page(page).Outputsize(outputsize).Execute()

MFs directory



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
	symbol := "1535462D" // string | Filter by symbol (optional)
	figi := "BBG00HMMLCH1" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "LU1206782309" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "120678230" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cik := "95953" // string | The CIK of an instrument for which data is requested (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	fundFamily := "Jackson National" // string | Filter by investment company that manages the fund (optional)
	fundType := "Small Blend" // string | Filter by the type of fund (optional)
	performanceRating := int64(4) // int64 | Filter by performance rating from `0` to `5` (optional)
	riskRating := int64(4) // int64 | Filter by risk rating from `0` to `5` (optional)
	page := int64(789) // int64 | Page number (optional) (default to 1)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsList(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Cik(cik).Country(country).FundFamily(fundFamily).FundType(fundType).PerformanceRating(performanceRating).RiskRating(riskRating).Page(page).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsList`: GetMutualFundsList200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsListRequest struct via the builder pattern


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
 **performanceRating** | **int64** | Filter by performance rating from &#x60;0&#x60; to &#x60;5&#x60; | 
 **riskRating** | **int64** | Filter by risk rating from &#x60;0&#x60; to &#x60;5&#x60; | 
 **page** | **int64** | Page number | [default to 1]
 **outputsize** | **int64** | Number of records in response | [default to 100]

### Return type

[**GetMutualFundsList200Response**](GetMutualFundsList200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsType

> GetMutualFundsType200Response GetMutualFundsType(ctx).FundType(fundType).Country(country).Execute()

MFs types



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
	fundType := "Jackson National" // string | Filter by the type of fund (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsType(context.Background()).FundType(fundType).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsType`: GetMutualFundsType200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fundType** | **string** | Filter by the type of fund | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 

### Return type

[**GetMutualFundsType200Response**](GetMutualFundsType200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsWorld

> GetMutualFundsWorld200Response GetMutualFundsWorld(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

MF full data



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
	symbol := "1535462D" // string | Symbol ticker of mutual fund (optional)
	figi := "BBG00HMMLCH1" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "LU1206782309" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "120678230" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsWorld(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsWorld`: GetMutualFundsWorld200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsWorld`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsWorldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of mutual fund | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetMutualFundsWorld200Response**](GetMutualFundsWorld200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsWorldComposition

> GetMutualFundsWorldComposition200Response GetMutualFundsWorldComposition(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

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
	symbol := "1535462D" // string | Symbol ticker of mutual fund (optional)
	figi := "BBG00HMMLCH1" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "LU1206782309" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "120678230" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsWorldComposition(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsWorldComposition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsWorldComposition`: GetMutualFundsWorldComposition200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsWorldComposition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsWorldCompositionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of mutual fund | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetMutualFundsWorldComposition200Response**](GetMutualFundsWorldComposition200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsWorldPerformance

> GetMutualFundsWorldPerformance200Response GetMutualFundsWorldPerformance(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

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
	symbol := "1535462D" // string | Symbol ticker of mutual fund (optional)
	figi := "BBG00HMMLCH1" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "LU1206782309" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "120678230" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsWorldPerformance(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsWorldPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsWorldPerformance`: GetMutualFundsWorldPerformance200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsWorldPerformance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsWorldPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of mutual fund | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetMutualFundsWorldPerformance200Response**](GetMutualFundsWorldPerformance200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsWorldPurchaseInfo

> GetMutualFundsWorldPurchaseInfo200Response GetMutualFundsWorldPurchaseInfo(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Purchase info



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
	symbol := "1535462D" // string | Symbol ticker of mutual fund (optional)
	figi := "BBG00HMMLCH1" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "LU1206782309" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "120678230" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsWorldPurchaseInfo(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsWorldPurchaseInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsWorldPurchaseInfo`: GetMutualFundsWorldPurchaseInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsWorldPurchaseInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsWorldPurchaseInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of mutual fund | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetMutualFundsWorldPurchaseInfo200Response**](GetMutualFundsWorldPurchaseInfo200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsWorldRatings

> GetMutualFundsWorldRatings200Response GetMutualFundsWorldRatings(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Ratings



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
	symbol := "1535462D" // string | Symbol ticker of mutual fund (optional)
	figi := "BBG00HMMLCH1" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "LU1206782309" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "120678230" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsWorldRatings(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsWorldRatings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsWorldRatings`: GetMutualFundsWorldRatings200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsWorldRatings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsWorldRatingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of mutual fund | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetMutualFundsWorldRatings200Response**](GetMutualFundsWorldRatings200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsWorldRisk

> GetMutualFundsWorldRisk200Response GetMutualFundsWorldRisk(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

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
	symbol := "1535462D" // string | Symbol ticker of mutual fund (optional)
	figi := "BBG00HMMLCH1" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "LU1206782309" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "120678230" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsWorldRisk(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsWorldRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsWorldRisk`: GetMutualFundsWorldRisk200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsWorldRisk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsWorldRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of mutual fund | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetMutualFundsWorldRisk200Response**](GetMutualFundsWorldRisk200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsWorldSummary

> GetMutualFundsWorldSummary200Response GetMutualFundsWorldSummary(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

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
	symbol := "1535462D" // string | Symbol ticker of mutual fund (optional)
	figi := "BBG00HMMLCH1" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "LU1206782309" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "120678230" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsWorldSummary(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsWorldSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsWorldSummary`: GetMutualFundsWorldSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsWorldSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsWorldSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of mutual fund | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetMutualFundsWorldSummary200Response**](GetMutualFundsWorldSummary200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFundsWorldSustainability

> GetMutualFundsWorldSustainability200Response GetMutualFundsWorldSustainability(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()

Sustainability



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
	symbol := "1535462D" // string | Symbol ticker of mutual fund (optional)
	figi := "BBG00HMMLCH1" // string | Filter by financial instrument global identifier (FIGI). This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "LU1206782309" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "120678230" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Accepts value in range [0,11] (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MutualFundsAPI.GetMutualFundsWorldSustainability(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MutualFundsAPI.GetMutualFundsWorldSustainability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFundsWorldSustainability`: GetMutualFundsWorldSustainability200Response
	fmt.Fprintf(os.Stdout, "Response from `MutualFundsAPI.GetMutualFundsWorldSustainability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFundsWorldSustainabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol ticker of mutual fund | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **dp** | **int64** | Number of decimal places for floating values. Accepts value in range [0,11] | [default to 5]

### Return type

[**GetMutualFundsWorldSustainability200Response**](GetMutualFundsWorldSustainability200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

