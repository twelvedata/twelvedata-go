# GetTimeSeriesHeikinashiCandles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesHeikinashiCandles200ResponseMeta**](GetTimeSeriesHeikinashiCandles200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesHeikinashiCandles200ResponseValuesInner**](GetTimeSeriesHeikinashiCandles200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesHeikinashiCandles200Response

`func NewGetTimeSeriesHeikinashiCandles200Response(meta GetTimeSeriesHeikinashiCandles200ResponseMeta, values []GetTimeSeriesHeikinashiCandles200ResponseValuesInner, status string, ) *GetTimeSeriesHeikinashiCandles200Response`

NewGetTimeSeriesHeikinashiCandles200Response instantiates a new GetTimeSeriesHeikinashiCandles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHeikinashiCandles200ResponseWithDefaults

`func NewGetTimeSeriesHeikinashiCandles200ResponseWithDefaults() *GetTimeSeriesHeikinashiCandles200Response`

NewGetTimeSeriesHeikinashiCandles200ResponseWithDefaults instantiates a new GetTimeSeriesHeikinashiCandles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesHeikinashiCandles200Response) GetMeta() GetTimeSeriesHeikinashiCandles200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesHeikinashiCandles200Response) GetMetaOk() (*GetTimeSeriesHeikinashiCandles200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesHeikinashiCandles200Response) SetMeta(v GetTimeSeriesHeikinashiCandles200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesHeikinashiCandles200Response) GetValues() []GetTimeSeriesHeikinashiCandles200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesHeikinashiCandles200Response) GetValuesOk() (*[]GetTimeSeriesHeikinashiCandles200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesHeikinashiCandles200Response) SetValues(v []GetTimeSeriesHeikinashiCandles200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesHeikinashiCandles200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesHeikinashiCandles200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesHeikinashiCandles200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


