# \AnalysisAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalystRatingsLight**](AnalysisAPI.md#GetAnalystRatingsLight) | **Get** /analyst_ratings/light | Analyst ratings snapshot
[**GetAnalystRatingsUsEquities**](AnalysisAPI.md#GetAnalystRatingsUsEquities) | **Get** /analyst_ratings/us_equities | Analyst ratings US equities
[**GetEarningsEstimate**](AnalysisAPI.md#GetEarningsEstimate) | **Get** /earnings_estimate | Earnings estimate
[**GetEpsRevisions**](AnalysisAPI.md#GetEpsRevisions) | **Get** /eps_revisions | EPS revisions
[**GetEpsTrend**](AnalysisAPI.md#GetEpsTrend) | **Get** /eps_trend | EPS trend
[**GetGrowthEstimates**](AnalysisAPI.md#GetGrowthEstimates) | **Get** /growth_estimates | Growth estimates
[**GetPriceTarget**](AnalysisAPI.md#GetPriceTarget) | **Get** /price_target | Price target
[**GetRecommendations**](AnalysisAPI.md#GetRecommendations) | **Get** /recommendations | Recommendations
[**GetRevenueEstimate**](AnalysisAPI.md#GetRevenueEstimate) | **Get** /revenue_estimate | Revenue estimate



## GetAnalystRatingsLight

> GetAnalystRatingsLight200Response GetAnalystRatingsLight(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).RatingChange(ratingChange).Outputsize(outputsize).Country(country).Execute()

Analyst ratings snapshot



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
	ratingChange := openapiclient.RatingChangeEnum("Maintains") // RatingChangeEnum | Filter by rating change action (optional)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 30)
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.GetAnalystRatingsLight(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).RatingChange(ratingChange).Outputsize(outputsize).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.GetAnalystRatingsLight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalystRatingsLight`: GetAnalystRatingsLight200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.GetAnalystRatingsLight`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalystRatingsLightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Filter by exchange name | 
 **ratingChange** | [**RatingChangeEnum**](RatingChangeEnum.md) | Filter by rating change action | 
 **outputsize** | **int64** | Number of records in response | [default to 30]
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 

### Return type

[**GetAnalystRatingsLight200Response**](GetAnalystRatingsLight200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalystRatingsUsEquities

> GetAnalystRatingsUsEquities200Response GetAnalystRatingsUsEquities(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).RatingChange(ratingChange).Outputsize(outputsize).Execute()

Analyst ratings US equities



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
	ratingChange := openapiclient.RatingChangeEnum("Maintains") // RatingChangeEnum | Filter by rating change action (optional)
	outputsize := int64(789) // int64 | Number of records in response (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.GetAnalystRatingsUsEquities(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Exchange(exchange).RatingChange(ratingChange).Outputsize(outputsize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.GetAnalystRatingsUsEquities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalystRatingsUsEquities`: GetAnalystRatingsUsEquities200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.GetAnalystRatingsUsEquities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalystRatingsUsEquitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **exchange** | **string** | Filter by exchange name | 
 **ratingChange** | [**RatingChangeEnum**](RatingChangeEnum.md) | Filter by rating change action | 
 **outputsize** | **int64** | Number of records in response | [default to 30]

### Return type

[**GetAnalystRatingsUsEquities200Response**](GetAnalystRatingsUsEquities200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEarningsEstimate

> GetEarningsEstimate200Response GetEarningsEstimate(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()

Earnings estimate



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
	figi := "BBG000B9Y5X2" // string | The FIGI of an instrument for which data is requested. This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | The country where the instrument is traded, e.g., `United States` or `US` (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.GetEarningsEstimate(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.GetEarningsEstimate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEarningsEstimate`: GetEarningsEstimate200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.GetEarningsEstimate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEarningsEstimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | The FIGI of an instrument for which data is requested. This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | The country where the instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **exchange** | **string** | Exchange where instrument is traded | 

### Return type

[**GetEarningsEstimate200Response**](GetEarningsEstimate200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEpsRevisions

> GetEpsRevisions200Response GetEpsRevisions(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()

EPS revisions



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
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	exchange := "NASDAQ" // string | Filter by exchange name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.GetEpsRevisions(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.GetEpsRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEpsRevisions`: GetEpsRevisions200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.GetEpsRevisions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEpsRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **exchange** | **string** | Filter by exchange name | 

### Return type

[**GetEpsRevisions200Response**](GetEpsRevisions200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEpsTrend

> GetEpsTrend200Response GetEpsTrend(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()

EPS trend



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
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	exchange := "NASDAQ" // string | Filter by exchange name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.GetEpsTrend(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.GetEpsTrend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEpsTrend`: GetEpsTrend200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.GetEpsTrend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEpsTrendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **exchange** | **string** | Filter by exchange name | 

### Return type

[**GetEpsTrend200Response**](GetEpsTrend200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGrowthEstimates

> GetGrowthEstimates200Response GetGrowthEstimates(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()

Growth estimates



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
	figi := "BBG000B9Y5X2" // string | The FIGI of an instrument for which data is requested. This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | The country where the instrument is traded, e.g., `United States` or `US` (optional)
	exchange := "NASDAQ" // string | Exchange where instrument is traded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.GetGrowthEstimates(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.GetGrowthEstimates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGrowthEstimates`: GetGrowthEstimates200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.GetGrowthEstimates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGrowthEstimatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | The FIGI of an instrument for which data is requested. This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | The country where the instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **exchange** | **string** | Exchange where instrument is traded | 

### Return type

[**GetGrowthEstimates200Response**](GetGrowthEstimates200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceTarget

> GetPriceTarget200Response GetPriceTarget(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()

Price target



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
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	exchange := "NASDAQ" // string | Filter by exchange name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.GetPriceTarget(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.GetPriceTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceTarget`: GetPriceTarget200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.GetPriceTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **exchange** | **string** | Filter by exchange name | 

### Return type

[**GetPriceTarget200Response**](GetPriceTarget200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendations

> GetRecommendations200Response GetRecommendations(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()

Recommendations



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
	figi := "BBG000B9Y5X2" // string | The FIGI of an instrument for which data is requested. This parameter is available on the <a href=\"https://twelvedata.com/pricing\">Ultra</a> plan (individual) and the <a href=\"https://twelvedata.com/pricing-business\">Enterprise</a> plan (business) and above. (optional)
	isin := "US0378331005" // string | Filter by international securities identification number (ISIN). ISIN access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	cusip := "594918104" // string | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the <a href=\"https://twelvedata.com/account/add-ons\">Data add-ons</a> section (optional)
	country := "United States" // string | The country where the instrument is traded, e.g., `United States` or `US` (optional)
	exchange := "NASDAQ" // string | The exchange name where the instrument is traded, e.g., `Nasdaq`, `NSE`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.GetRecommendations(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.GetRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecommendations`: GetRecommendations200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.GetRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | The FIGI of an instrument for which data is requested. This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | The country where the instrument is traded, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **exchange** | **string** | The exchange name where the instrument is traded, e.g., &#x60;Nasdaq&#x60;, &#x60;NSE&#x60;. | 

### Return type

[**GetRecommendations200Response**](GetRecommendations200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevenueEstimate

> GetRevenueEstimate200Response GetRevenueEstimate(ctx).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Dp(dp).Execute()

Revenue estimate



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
	country := "United States" // string | Filter by country name or alpha code, e.g., `United States` or `US` (optional)
	exchange := "NASDAQ" // string | Filter by exchange name (optional)
	dp := int64(789) // int64 | Number of decimal places for floating values. Should be in range [0,11] inclusive (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.GetRevenueEstimate(context.Background()).Symbol(symbol).Figi(figi).Isin(isin).Cusip(cusip).Country(country).Exchange(exchange).Dp(dp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.GetRevenueEstimate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevenueEstimate`: GetRevenueEstimate200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.GetRevenueEstimate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRevenueEstimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Filter by symbol | 
 **figi** | **string** | Filter by financial instrument global identifier (FIGI). This parameter is available on the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing\&quot;&gt;Ultra&lt;/a&gt; plan (individual) and the &lt;a href&#x3D;\&quot;https://twelvedata.com/pricing-business\&quot;&gt;Enterprise&lt;/a&gt; plan (business) and above. | 
 **isin** | **string** | Filter by international securities identification number (ISIN). ISIN access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **cusip** | **string** | The CUSIP of an instrument for which data is requested. CUSIP access is activating in the &lt;a href&#x3D;\&quot;https://twelvedata.com/account/add-ons\&quot;&gt;Data add-ons&lt;/a&gt; section | 
 **country** | **string** | Filter by country name or alpha code, e.g., &#x60;United States&#x60; or &#x60;US&#x60; | 
 **exchange** | **string** | Filter by exchange name | 
 **dp** | **int64** | Number of decimal places for floating values. Should be in range [0,11] inclusive | [default to 5]

### Return type

[**GetRevenueEstimate200Response**](GetRevenueEstimate200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

