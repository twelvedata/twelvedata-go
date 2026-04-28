# GetTimeSeriesApo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesApo200ResponseMeta**](GetTimeSeriesApo200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesApo200ResponseValuesInner**](GetTimeSeriesApo200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesApo200Response

`func NewGetTimeSeriesApo200Response(meta GetTimeSeriesApo200ResponseMeta, values []GetTimeSeriesApo200ResponseValuesInner, status string, ) *GetTimeSeriesApo200Response`

NewGetTimeSeriesApo200Response instantiates a new GetTimeSeriesApo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesApo200ResponseWithDefaults

`func NewGetTimeSeriesApo200ResponseWithDefaults() *GetTimeSeriesApo200Response`

NewGetTimeSeriesApo200ResponseWithDefaults instantiates a new GetTimeSeriesApo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesApo200Response) GetMeta() GetTimeSeriesApo200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesApo200Response) GetMetaOk() (*GetTimeSeriesApo200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesApo200Response) SetMeta(v GetTimeSeriesApo200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesApo200Response) GetValues() []GetTimeSeriesApo200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesApo200Response) GetValuesOk() (*[]GetTimeSeriesApo200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesApo200Response) SetValues(v []GetTimeSeriesApo200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesApo200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesApo200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesApo200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


