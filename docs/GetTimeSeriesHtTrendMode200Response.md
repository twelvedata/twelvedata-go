# GetTimeSeriesHtTrendMode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesHtTrendMode200ResponseMeta**](GetTimeSeriesHtTrendMode200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesHtTrendMode200ResponseValuesInner**](GetTimeSeriesHtTrendMode200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesHtTrendMode200Response

`func NewGetTimeSeriesHtTrendMode200Response(meta GetTimeSeriesHtTrendMode200ResponseMeta, values []GetTimeSeriesHtTrendMode200ResponseValuesInner, status string, ) *GetTimeSeriesHtTrendMode200Response`

NewGetTimeSeriesHtTrendMode200Response instantiates a new GetTimeSeriesHtTrendMode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtTrendMode200ResponseWithDefaults

`func NewGetTimeSeriesHtTrendMode200ResponseWithDefaults() *GetTimeSeriesHtTrendMode200Response`

NewGetTimeSeriesHtTrendMode200ResponseWithDefaults instantiates a new GetTimeSeriesHtTrendMode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesHtTrendMode200Response) GetMeta() GetTimeSeriesHtTrendMode200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesHtTrendMode200Response) GetMetaOk() (*GetTimeSeriesHtTrendMode200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesHtTrendMode200Response) SetMeta(v GetTimeSeriesHtTrendMode200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesHtTrendMode200Response) GetValues() []GetTimeSeriesHtTrendMode200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesHtTrendMode200Response) GetValuesOk() (*[]GetTimeSeriesHtTrendMode200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesHtTrendMode200Response) SetValues(v []GetTimeSeriesHtTrendMode200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesHtTrendMode200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesHtTrendMode200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesHtTrendMode200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


