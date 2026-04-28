# GetTimeSeriesHtSine200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesHtSine200ResponseMeta**](GetTimeSeriesHtSine200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesHtSine200ResponseValuesInner**](GetTimeSeriesHtSine200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesHtSine200Response

`func NewGetTimeSeriesHtSine200Response(meta GetTimeSeriesHtSine200ResponseMeta, values []GetTimeSeriesHtSine200ResponseValuesInner, status string, ) *GetTimeSeriesHtSine200Response`

NewGetTimeSeriesHtSine200Response instantiates a new GetTimeSeriesHtSine200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtSine200ResponseWithDefaults

`func NewGetTimeSeriesHtSine200ResponseWithDefaults() *GetTimeSeriesHtSine200Response`

NewGetTimeSeriesHtSine200ResponseWithDefaults instantiates a new GetTimeSeriesHtSine200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesHtSine200Response) GetMeta() GetTimeSeriesHtSine200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesHtSine200Response) GetMetaOk() (*GetTimeSeriesHtSine200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesHtSine200Response) SetMeta(v GetTimeSeriesHtSine200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesHtSine200Response) GetValues() []GetTimeSeriesHtSine200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesHtSine200Response) GetValuesOk() (*[]GetTimeSeriesHtSine200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesHtSine200Response) SetValues(v []GetTimeSeriesHtSine200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesHtSine200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesHtSine200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesHtSine200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


