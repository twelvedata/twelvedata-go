# GetTimeSeriesLog10200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesLog10200ResponseMeta**](GetTimeSeriesLog10200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesLog10200ResponseValuesInner**](GetTimeSeriesLog10200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesLog10200Response

`func NewGetTimeSeriesLog10200Response(meta GetTimeSeriesLog10200ResponseMeta, values []GetTimeSeriesLog10200ResponseValuesInner, status string, ) *GetTimeSeriesLog10200Response`

NewGetTimeSeriesLog10200Response instantiates a new GetTimeSeriesLog10200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLog10200ResponseWithDefaults

`func NewGetTimeSeriesLog10200ResponseWithDefaults() *GetTimeSeriesLog10200Response`

NewGetTimeSeriesLog10200ResponseWithDefaults instantiates a new GetTimeSeriesLog10200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesLog10200Response) GetMeta() GetTimeSeriesLog10200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesLog10200Response) GetMetaOk() (*GetTimeSeriesLog10200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesLog10200Response) SetMeta(v GetTimeSeriesLog10200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesLog10200Response) GetValues() []GetTimeSeriesLog10200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesLog10200Response) GetValuesOk() (*[]GetTimeSeriesLog10200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesLog10200Response) SetValues(v []GetTimeSeriesLog10200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesLog10200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesLog10200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesLog10200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


