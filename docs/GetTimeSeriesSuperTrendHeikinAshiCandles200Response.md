# GetTimeSeriesSuperTrendHeikinAshiCandles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta**](GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner**](GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesSuperTrendHeikinAshiCandles200Response

`func NewGetTimeSeriesSuperTrendHeikinAshiCandles200Response(meta GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta, values []GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner, status string, ) *GetTimeSeriesSuperTrendHeikinAshiCandles200Response`

NewGetTimeSeriesSuperTrendHeikinAshiCandles200Response instantiates a new GetTimeSeriesSuperTrendHeikinAshiCandles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseWithDefaults

`func NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseWithDefaults() *GetTimeSeriesSuperTrendHeikinAshiCandles200Response`

NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseWithDefaults instantiates a new GetTimeSeriesSuperTrendHeikinAshiCandles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetMeta() GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetMetaOk() (*GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) SetMeta(v GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetValues() []GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetValuesOk() (*[]GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) SetValues(v []GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


