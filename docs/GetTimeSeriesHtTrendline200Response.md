# GetTimeSeriesHtTrendline200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesHtTrendline200ResponseMeta**](GetTimeSeriesHtTrendline200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesHtTrendline200ResponseValuesInner**](GetTimeSeriesHtTrendline200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesHtTrendline200Response

`func NewGetTimeSeriesHtTrendline200Response(meta GetTimeSeriesHtTrendline200ResponseMeta, values []GetTimeSeriesHtTrendline200ResponseValuesInner, status string, ) *GetTimeSeriesHtTrendline200Response`

NewGetTimeSeriesHtTrendline200Response instantiates a new GetTimeSeriesHtTrendline200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtTrendline200ResponseWithDefaults

`func NewGetTimeSeriesHtTrendline200ResponseWithDefaults() *GetTimeSeriesHtTrendline200Response`

NewGetTimeSeriesHtTrendline200ResponseWithDefaults instantiates a new GetTimeSeriesHtTrendline200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesHtTrendline200Response) GetMeta() GetTimeSeriesHtTrendline200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesHtTrendline200Response) GetMetaOk() (*GetTimeSeriesHtTrendline200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesHtTrendline200Response) SetMeta(v GetTimeSeriesHtTrendline200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesHtTrendline200Response) GetValues() []GetTimeSeriesHtTrendline200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesHtTrendline200Response) GetValuesOk() (*[]GetTimeSeriesHtTrendline200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesHtTrendline200Response) SetValues(v []GetTimeSeriesHtTrendline200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesHtTrendline200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesHtTrendline200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesHtTrendline200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


