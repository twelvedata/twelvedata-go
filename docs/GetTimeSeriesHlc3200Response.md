# GetTimeSeriesHlc3200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesHlc3200ResponseMeta**](GetTimeSeriesHlc3200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesHlc3200ResponseValuesInner**](GetTimeSeriesHlc3200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesHlc3200Response

`func NewGetTimeSeriesHlc3200Response(meta GetTimeSeriesHlc3200ResponseMeta, values []GetTimeSeriesHlc3200ResponseValuesInner, status string, ) *GetTimeSeriesHlc3200Response`

NewGetTimeSeriesHlc3200Response instantiates a new GetTimeSeriesHlc3200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHlc3200ResponseWithDefaults

`func NewGetTimeSeriesHlc3200ResponseWithDefaults() *GetTimeSeriesHlc3200Response`

NewGetTimeSeriesHlc3200ResponseWithDefaults instantiates a new GetTimeSeriesHlc3200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesHlc3200Response) GetMeta() GetTimeSeriesHlc3200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesHlc3200Response) GetMetaOk() (*GetTimeSeriesHlc3200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesHlc3200Response) SetMeta(v GetTimeSeriesHlc3200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesHlc3200Response) GetValues() []GetTimeSeriesHlc3200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesHlc3200Response) GetValuesOk() (*[]GetTimeSeriesHlc3200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesHlc3200Response) SetValues(v []GetTimeSeriesHlc3200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesHlc3200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesHlc3200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesHlc3200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


