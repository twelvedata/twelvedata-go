# \AdvancedAPI

All URIs are relative to *https://api.twelvedata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Advanced**](AdvancedAPI.md#Advanced) | **Post** /batch | Batches
[**GetApiUsage**](AdvancedAPI.md#GetApiUsage) | **Get** /api_usage | API usage



## Advanced

> Advanced200Response Advanced(ctx).Key(key).Execute()

Batches



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
	key := map[string]AdvancedRequestValue{"key": *openapiclient.NewAdvancedRequestValue()} // map[string]AdvancedRequestValue | Map of requests (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedAPI.Advanced(context.Background()).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.Advanced``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Advanced`: Advanced200Response
	fmt.Fprintf(os.Stdout, "Response from `AdvancedAPI.Advanced`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdvancedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | [**map[string]AdvancedRequestValue**](AdvancedRequestValue.md) | Map of requests | 

### Return type

[**Advanced200Response**](Advanced200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiUsage

> GetApiUsage200Response GetApiUsage(ctx).Format(format).Delimiter(delimiter).Timezone(timezone).Execute()

API usage



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
	format := openapiclient.FormatEnum("JSON") // FormatEnum | Output format (optional) (default to "JSON")
	delimiter := "delimiter_example" // string | Specify the delimiter used when downloading the CSV file (optional) (default to ";")
	timezone := "timezone_example" // string | Timezone at which output datetime will be displayed. Supports: <ul> <li>1. <code>UTC</code> for datetime at universal UTC standard</li> <li>2. Timezone name according to the IANA Time Zone Database. E.g. <code>America/New_York</code>, <code>Asia/Singapore</code>. Full list of timezones can be found <a href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\" target=\"blank\">here</a>.</li> </ul> <i>Take note that the IANA Timezone name is case-sensitive</i> (optional) (default to "UTC")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedAPI.GetApiUsage(context.Background()).Format(format).Delimiter(delimiter).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAPI.GetApiUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiUsage`: GetApiUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `AdvancedAPI.GetApiUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**FormatEnum**](FormatEnum.md) | Output format | [default to &quot;JSON&quot;]
 **delimiter** | **string** | Specify the delimiter used when downloading the CSV file | [default to &quot;;&quot;]
 **timezone** | **string** | Timezone at which output datetime will be displayed. Supports: &lt;ul&gt; &lt;li&gt;1. &lt;code&gt;UTC&lt;/code&gt; for datetime at universal UTC standard&lt;/li&gt; &lt;li&gt;2. Timezone name according to the IANA Time Zone Database. E.g. &lt;code&gt;America/New_York&lt;/code&gt;, &lt;code&gt;Asia/Singapore&lt;/code&gt;. Full list of timezones can be found &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\&quot; target&#x3D;\&quot;blank\&quot;&gt;here&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; &lt;i&gt;Take note that the IANA Timezone name is case-sensitive&lt;/i&gt; | [default to &quot;UTC&quot;]

### Return type

[**GetApiUsage200Response**](GetApiUsage200Response.md)

### Authorization

[authorizationHeader](../README.md#authorizationHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

